# Content Streaming on the Web with Amazon MSK

The repository contains the sample code for the article.

## Introduction

The goal of the article is to examplify how you can build scalable web and mobile applications with [Amazon MSK](https://aws.amazon.com/msk/). The applications use [pub/sub](https://aws.amazon.com/pub-sub-messaging/) messaging to distribute content on the web or on mobile. New content is pushed to subscribed clients, rather then pulled by them. This is a more efficient model of content distribution.

Also the messages already exist in the materalized representation for the client. You either store the final [JSON](https://www.json.org/), [Thrift](https://thrift.apache.org) or in this example [Protocol Buffers](https://developers.google.com/protocol-buffers) messages in the underlying message queue. There is no need to query a database for a request.

The example for such a [pub/sub](https://aws.amazon.com/pub-sub-messaging/) messaging service you find in this repository is a microblogging service. It consists of a [React](https://reactjs.org/) app that allows you to publish and read articles. A service which implements the [gRPC](https://grpc.io/) to publish and subscribe to the messages (articles) that are created. And the [AWS CloudFormation](https://aws.amazon.com/cloudformation/) templates to run the service. The client uses [gRPC for Web Clients](https://github.com/grpc/grpc-web) to communicate with the pub/sub service.

[Protobuf](https://developers.google.com/protocol-buffers) are a language-neutral, platform-neutral extensible mechanism for seralizing structured data. It is like XML for for your data. You define your service data and interface, and use source code generation to produce the needed coded for the server and client. No need to write data object mappers or type-safe your API.

## Components

- [Protobuf](https://developers.google.com/protocol-buffers) protocol for the service
- [React](https://reactjs.org/) app with [gRPC for Web Clients](https://github.com/grpc/grpc-web)
- [gRPC](https://grpc.io) service to subscribe and publish articles
- [AWS CloudFormation](https://aws.amazon.com/cloudformation/) templates to run the service

---

## Instructions

> :warning: Running the application will create and consume AWS resources. This will cost you money. Make sure you shut down/remove all resources once you are finished to avoid ongoing charges to your AWS account.

First, you need your own AWS account. Follow these [steps](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/) if you do not have an account.

Next, you have to setup the environment for the `deploy.sh` script. You use the script to deploy and destroy the application.

```bash
export PROJECT_NAME=<YOUR_PROJECT_NAME>
export AWS_ACCOUNT_ID=<YOUR_ACCOUNT_ID>
export AWS_DEFAULT_REGION=<YOUR_AWS_REGION>
# (see https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html)
export ENVOY_IMAGE=<ENOVY_IMAGE>
```

Running the script will deploy the application.

> This can take a while, as you will create a high available Kafka with Amazon MSK.

```bash
./deploy.sh
```

When the process is finished, you get the URL for the application with the following command.

```bash
./deploy.sh print_endpoint
```

The client is run by the following command. It is not deployed with the application, and listens on [localhost:3000](http://localhost:3000) by default.

```bash
make start
```

## License

[MIT](/LICENSE)
