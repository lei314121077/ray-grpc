#!/usr/bin/env bash
protoc --go_out=plugins=grpc:. grpconepb.proto