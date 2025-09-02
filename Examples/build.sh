#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"

echo "Build and test examples"
./RTTI/build.sh
./Injection/build.sh
./Calculator/build.sh
./UnitTest/build.sh
./Primes/build.sh
./OptionalClass/build.sh
./ThreadSafety/build.sh

echo "Examples: build and test are done and successful"