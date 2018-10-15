#!/bin/bash

# setup
set -ex
scriptdir="$(dirname "$0")"
cd $scriptdir

# main
GOOS=linux go build -o synthsub
docker build -t synthsub .

# cleanup
rm synthsub
set +ex
