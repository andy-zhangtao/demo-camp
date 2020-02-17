#!/bin/bash

docker run -d -p 80:80 -v $PWD/default.conf:/etc/nginx/conf.d/default.conf --name nginx nginx

docker run -d -v $PWD/envoy.yaml:/etc/envoy/envoy.yaml -p 9901:9901 -p 10000:10000  --name envoy envoyproxy/envoy -l trace -c /etc/envoy/envoy.yaml

docker logs -f envoy