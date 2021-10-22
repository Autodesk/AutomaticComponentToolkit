#!/bin/bash
set -euxo pipefail

cd "$(dirname "$0")"
echo "Download JNA"
[ -f jna-5.5.0.jar ] || wget https://repo1.maven.org/maven2/net/java/dev/jna/jna/5.5.0/jna-5.5.0.jar

echo "Compile Java Bindings"
javac -classpath *.jar rtti/*

echo "Create JAR"
jar cvf rtti-1.0.0.jar rtti
