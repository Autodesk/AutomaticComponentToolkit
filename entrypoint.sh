#!/bin/sh -l

echo "Action selected: $1"
time=$(date)
echo "::set-output name=time::$time"

case "$1" in 
    act)
       sh Build/build.sh
    ;;
    examples)
       sh Examples/build.sh
    ;;
    all)
       sh Build/build.sh
       sh Examples/build.sh
    ;;
    *)
    echo "Use one of available commands:"
    echo "  ./entrypoint.sh act      - build ACT binaries" 
    echo "  ./entrypoint.sh examples - build and run projects in Examples folder" 
    echo "  ./entrypoint.sh all      - build ACT binaries and then build and run projects in Examples folder" 
    exit 1
    ;;
esac


