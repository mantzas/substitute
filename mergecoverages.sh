#!/bin/bash
set -ev

echo "mode: atomic" > substitute.coverprofile

for pkg in `find . -name "*.coverprofile" -type f | grep -v 'substitute.coverprofile'`; do
    echo $pkg
    tail -n +2 $pkg >> substitute.coverprofile
done