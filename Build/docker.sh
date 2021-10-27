#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"

case "$1" in 
    build) 
        docker build -t act-build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) .
    ;;
    act)
        docker run -it --rm -v $PWD/..:/data act-build /data/Build/build.sh
    ;;
    examples)
        docker run -it --rm -v $PWD/..:/data act-build /data/Examples/build.sh
    ;;
    all)
        docker run -it --rm -v $PWD/..:/data act-build /data/Build/build.sh
        docker run -it --rm -v $PWD/..:/data act-build /data/Examples/build.sh
    ;;
    cli)
        docker run -it --rm -v $PWD/..:/data --entrypoint bash act-build -l
    ;;
    *)
    echo "Use one of availbale commands:"
    echo "  ./docker.sh build    - build docker image" 
    echo "  ./docker.sh act      - build ACT binaries" 
    echo "  ./docker.sh examples - build and run projects in Examples folder" 
    echo "  ./docker.sh all      - build ACT binaries and then build and run projects in Examples folder" 
    echo "  ./docker.sh cli      - open bash session inside Docker with source code mapped to '/data' directory" 
    exit 1
    ;;
esac
