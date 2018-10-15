#!/bin/bash

# setup
set -ex
scriptdir="$(dirname "$0")"
cd $scriptdir

# main
GOOS=linux go build -o synthpub
docker build -t synthpub .

# cleanup
rm synthpub
set +ex
