#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"

echo "Build and test example"
./RTTI/build.sh
./Injection/build.sh

echo "Examples: build and test are done and successful"