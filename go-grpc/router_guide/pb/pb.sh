#!/usr/bin/env bash
protoc chat.proto --go_out=plugins=grpc:.