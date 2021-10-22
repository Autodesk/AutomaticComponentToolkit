#!/bin/bash
set -euxo pipefail

cd "$(dirname "$0")"

JnaJar="jna-5.5.0.jar"
Classpath=".:${JnaJar}:../../Bindings/Java9/"
if [[ "$OSTYPE" == "linux-gnu" ]]; then
	Classpath=".:${JnaJar}:../../Bindings/Java9/"
elif [[ "$OSTYPE" == "darwin"* ]]; then
	Classpath=".:${JnaJar}:../../Bindings/Java9/"
elif [[ "$OSTYPE" == "cygwin" ]]; then
	Classpath=".;${JnaJar};../../Bindings/Java9/"
elif [[ "$OSTYPE" == "msys" ]]; then
	Classpath=".;${JnaJar};../../Bindings/Java9/"
elif [[ "$OSTYPE" == "win32" ]]; then
	Classpath=".;${JnaJar};../../Bindings/Java9/"
else
	echo "Unknown system: "$OSTYPE
	exit 1
fi

echo "Download JNA"
[ -f jna-5.5.0.jar ] || wget https://repo1.maven.org/maven2/net/java/dev/jna/jna/5.5.0/jna-5.5.0.jar

echo "Compile Java bindings"
javac -classpath "${JnaJar}" ../../Bindings/Java9/rtti/*.java
echo "Compile Java example"
javac -classpath $Classpath RTTI_Example.java

echo "Execute example"
java -classpath $Classpath RTTI_Example
