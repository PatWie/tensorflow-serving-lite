#!/usr/bin/env bash

TF_TYPE="cpu" # Change to "gpu" for GPU support

if [ ! -d libtf ]; then
  mkdir -p libtf
  cd libtf
  curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.10.1.tar.gz" | tar -C . -xz
  cd ../
fi


# git clone https://github.com/tensorflow/serving.git _serving
# git clone https://github.com/tensorflow/tensorflow.git _tensorflow

# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
# go get -u google.golang.org/grpc
# go get -u github.com/golang/protobuf/ptypes/wrappers

# export PATH=$GOPATH/bin:${PATH}

# mkdir -p vendor
# PROTOC_OPTS='-I _tensorflow -I _serving --go_out=plugins=grpc:vendor'

# protoc $PROTOC_OPTS _serving/tensorflow_serving/apis/*.proto
# protoc $PROTOC_OPTS _serving/tensorflow_serving/config/*.proto
# protoc $PROTOC_OPTS _serving/tensorflow_serving/util/*.proto
# protoc $PROTOC_OPTS _serving/tensorflow_serving/sources/storage_path/*.proto
# protoc $PROTOC_OPTS _tensorflow/tensorflow/core/framework/*.proto
# protoc $PROTOC_OPTS _tensorflow/tensorflow/core/example/*.proto
# protoc $PROTOC_OPTS _tensorflow/tensorflow/core/lib/core/*.proto
# protoc $PROTOC_OPTS _tensorflow/tensorflow/core/protobuf/{saver,meta_graph,saved_model}.proto

