#!/usr/bin/env bash
# Patrick Wieschollek <mail@patwie.com>

pp=`realpath libtf/lib`

echo "export LD_LIBRARY_PATH=${pp}:\${LD_LIBRARY_PATH}"
echo "export LIBRARY_PATH=${pp}:\${LIBRARY_PATH}"

export LD_LIBRARY_PATH=${pp}:${LD_LIBRARY_PATH}
export LIBRARY_PATH=${pp}:${LIBRARY_PATH}

go get
go build tensorflow-serving-lite.go
