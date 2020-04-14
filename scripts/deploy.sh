#!/usr/bin/env bash

set -e

# extract the information from the CLOUD9 environment
TMP_ACCOUNT_ID="$(aws sts get-caller-identity --query Account --output text)"
TMP_DEFAULT_REGION="$(curl -s http://169.254.169.254/latest/dynamic/instance-identity/document | jq -c -r .region)"

# set defaults
PROJECT_NAME="${C9_PROJECT:-$PROJECT_NAME}"
AWS_ACCOUNT_ID="${AWS_ACCOUNT_ID:-$TMP_ACCOUNT_ID}"
AWS_DEFAULT_REGION="${AWS_DEFAULT_REGION:-$TMP_DEFAULT_REGION}"

if [ -z $PROJECT_NAME ]; then
    echo "PROJECT_NAME environment variable is not set."
    exit 1
fi

if [ -z $AWS_ACCOUNT_ID ]; then
    echo "AWS_ACCOUNT_ID environment variable is not set."
    exit 1
fi

if [ -z $AWS_DEFAULT_REGION ]; then
    echo "AWS_DEFAULT_REGION environment variable is not set."
    exit 1
fi

if [ -z $KEY_PAIR ]; then
    echo "KEYPAIR environment variable is not set."
    exit 1
fi

DIR="$(pwd)"
ECR_IMAGE_PREFIX=${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${PROJECT_NAME}
TEMPLATES="$DIR/templates"

deploy_images() {
    echo "Deploying Server images to ECR..."
    for app in envoy server; do
        aws ecr describe-repositories --repository-name ${PROJECT_NAME}/${app} >/dev/null 2>&1 || aws ecr create-repository --repository-name ${PROJECT_NAME}/${app}
        docker build -t ${ECR_IMAGE_PREFIX}/${app} -f ${app}.Dockerfile ${DIR} --build-arg GO_PROXY=${GO_PROXY:-"https://proxy.golang.org"}
        $(aws ecr get-login --no-include-email)
        docker push ${ECR_IMAGE_PREFIX}/${app}
    done
}

deploy_infra() {
    echo "Deploying Cloud Formation stack: \"${PROJECT_NAME}-infra\" containing VPC and Cloud Map namespace..."
    aws cloudformation deploy \
        --no-fail-on-empty-changeset \
        --stack-name "${PROJECT_NAME}-infra"\
        --template-file "${TEMPLATES}/infra.yaml" \
        --capabilities CAPABILITY_IAM \
        --parameter-overrides "ProjectName=${PROJECT_NAME}" "KeyPair=${KEY_PAIR}" "ClusterName=${PROJECT_NAME}"
}

deploy_app() {
    echo "Deploying Cloud Formation stack: \"${PROJECT_NAME}-app\" containing ALB, ECS Tasks, and Cloud Map Services..."
    aws cloudformation deploy \
        --no-fail-on-empty-changeset \
        --stack-name "${PROJECT_NAME}-app" \
        --template-file "${TEMPLATES}/app.yaml" \
        --capabilities CAPABILITY_IAM \
        --parameter-overrides "ProjectName=${PROJECT_NAME}" "ServerImage=${ECR_IMAGE_PREFIX}/server" "EnvoyFrontImage=${ECR_IMAGE_PREFIX}/envoy" "BootstrapBrokers=$(get_bootstrap_brokers)"
}

# Save this for later when App Mesh is ready
deploy_mesh() {
    mesh_name="${PROJECT_NAME}"

    echo "Deploying Mesh: \"${mesh_name}\"..."

    aws configure add-model \
      --service-name appmesh-preview \
      --service-model https://raw.githubusercontent.com/aws/aws-app-mesh-roadmap/master/appmesh-preview/service-model.json

    aws appmesh-preview create-mesh --mesh-name $mesh_name --cli-input-json file://${DIR}/mesh/mesh.json
    aws appmesh-preview create-virtual-node --mesh-name $mesh_name --cli-input-json file://${DIR}/mesh/serverNode.json
    aws appmesh-preview create-virtual-router --mesh-name $mesh_name --cli-input-json file://${DIR}/mesh/serverRouter.json
    aws appmesh-preview create-virtual-service --mesh-name $mesh_name --cli-input-json file://${DIR}/mesh/serverService.json
    aws appmesh-preview create-route --mesh-name $mesh_name --cli-input-json file://${DIR}/mesh/serverRoute.json

    # aws cloudformation deploy \
    #     --no-fail-on-empty-changeset \
    #     --stack-name "${PROJECT_NAME}-mesh" \
    #     --template-file "${DIR}/mesh.yaml" \
    #     --capabilities CAPABILITY_IAM \
    #     --parameter-overrides "ProjectName=${PROJECT_NAME}"
}

# Save this for later when App Mesh is ready
delete_mesh() {
    mesh_name="${PROJECT_NAME}"

    echo "Deleting Mesh: \"${mesh_name}\"..."

    aws appmesh-preview delete-route --mesh-name $mesh_name --virtual-router-name virtual-router --route-name server
    aws appmesh-preview delete-virtual-service --mesh-name $mesh_name --virtual-service-name server.content.local
    aws appmesh-preview delete-virtual-router --mesh-name $mesh_name --virtual-router-name virtual-router
    aws appmesh-preview delete-virtual-node --mesh-name $mesh_name --virtual-node-name server
    aws appmesh-preview delete-mesh --mesh-name $mesh_name
}

print_bastion() {
    ip=$(aws cloudformation describe-stacks \
        --stack-name="${PROJECT_NAME}-infra" \
        --query="Stacks[0].Outputs[?OutputKey=='BastionIp'].OutputValue" \
        --output=text)
    echo "${ip}"
}

print_endpoint() {
    prefix=$(aws cloudformation describe-stacks \
        --stack-name="${PROJECT_NAME}-app" \
        --query="Stacks[0].Outputs[?OutputKey=='PublicEndpoint'].OutputValue" \
        --output=text)
    echo "${prefix}"
}

print_bootstrap_brokers() {
  echo "Bootstrap Brokers:"
  echo "$(get_bootstrap_brokers)"
}

get_bootstrap_brokers() {
  local arn="$@"
  local output

  if [ "$arn" == "" ]; then
    arn=$(aws cloudformation describe-stacks \
      --stack-name="${PROJECT_NAME}-infra" \
      --query="Stacks[0].Outputs[?OutputKey=='MSKClusterArn'].OutputValue" \
      --output=text)
  fi

  output=$(aws kafka get-bootstrap-brokers \
      --cluster-arn="${arn}" \
      --output=text)

  echo $output
}

deploy_stacks() {
    deploy_images
    deploy_infra
    deploy_app
}

delete_cfn_stack() {
    stack_name=$1
    echo "Deleting Cloud Formation stack: \"${stack_name}\"..."
    aws cloudformation delete-stack --stack-name $stack_name
    echo 'Waiting for the stack to be deleted, this may take a few minutes...'
    aws cloudformation wait stack-delete-complete --stack-name $stack_name
    echo 'Done'
}

delete_images() {
    for app in server; do
        echo "deleting repository \"${app}\"..."
        aws ecr delete-repository \
           --repository-name $PROJECT_NAME/$app \
           --force
    done
}

delete_stacks() {
    delete_cfn_stack "${PROJECT_NAME}-app"
    delete_cfn_stack "${PROJECT_NAME}-infra"
    # delete_cfn_stack "${PROJECT_NAME}-mesh"
    delete_images
}

action=${1:-"deploy"}
if [ "$action" == "delete" ]; then
    delete_stacks
    exit 0
fi

if [ "$action" == "print_endpoint" ]; then
    print_endpoint
    exit 0
fi

deploy_stacks
