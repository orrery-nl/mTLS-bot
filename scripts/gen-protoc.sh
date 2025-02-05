#!/bin/bash

# Generate the protobuf files
#
protoc --go_out=. \
    --go-grpc_out=. \
    api/root/*.proto
