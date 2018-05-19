#!/usr/bin/env bash
protoc server.proto --go_out=plugins=grpc:.