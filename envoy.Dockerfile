FROM envoyproxy/envoy-dev:latest

RUN apt-get update && apt-get install -y curl \
  && rm -rf /var/lib/apt/lists/*

COPY ./envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log
