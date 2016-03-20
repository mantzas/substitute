#!/bin/bash
set -ev
echo "mode: count" > coverage.out

for pkg in $(go list ./... | grep -v '/vendor/')
do
    go test -v -covermode=count -coverprofile=coverage_tmp.out $pkg
    
    if [ -f coverage_tmp.out ];
    then
        tail -n +2 coverage_tmp.out >> coverage.out
    fi
done