#! /bin/bash

cd $(dirname $0)/../rpc

SRC_DIR=./proto
DST_DIR=./gen

protoc --proto_path=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/*.proto
