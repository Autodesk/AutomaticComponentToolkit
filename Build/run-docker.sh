#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"

docker run -it --rm -v $PWD/..:/data act-build /data/Build/build.sh
docker run -it --rm -v $PWD/..:/data act-build /data/Examples/build.sh
