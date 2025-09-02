#!/bin/bash
set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

JnaJar="jna-5.5.0.jar"
Classpath=".:${JnaJar}:../../../Calculation_component/Bindings/Java9/:../../../Numbers_component/Bindings/Java9"
if [[ "$OSTYPE" == "linux-gnu" ]]; then
	Classpath=".:${JnaJar}:../../../Calculation_component/Bindings/Java9/:../../../Numbers_component/Bindings/Java9"
elif [[ "$OSTYPE" == "darwin"* ]]; then
	Classpath=".:${JnaJar}:../../../Calculation_component/Bindings/Java9/:../../../Numbers_component/Bindings/Java9"
elif [[ "$OSTYPE" == "cygwin" ]]; then
	Classpath=".;${JnaJar};../../../Calculation_component/Bindings/Java9/:../../../Numbers_component/Bindings/Java9"
elif [[ "$OSTYPE" == "msys" ]]; then
	Classpath=".;${JnaJar};../../../Calculation_component/Bindings/Java9/:../../../Numbers_component/Bindings/Java9"
elif [[ "$OSTYPE" == "win32" ]]; then
	Classpath=".;${JnaJar};../../../Calculation_component/Bindings/Java9/:../../../Numbers_component/Bindings/Java9"
else
	echo "Unknown system: "$OSTYPE
	exit 1
fi

echo "Download JNA"
[ -f jna-5.5.0.jar ] || curl -O https://repo1.maven.org/maven2/net/java/dev/jna/jna/5.5.0/jna-5.5.0.jar

echo "Compile Java bindings"
javac -encoding UTF8 -classpath "${JnaJar}" ../../../Calculation_component/Bindings/Java9/calculation/*.java ../../../Numbers_component/Bindings/Java9/numbers/*.java
echo "Compile Java example"
javac -encoding UTF8 -classpath $Classpath Calculation_Example.java

echo "Test C++ library"
java -ea -classpath $Classpath Calculation_Example $PWD/../../../Calculation_component/Implementations/Cpp/build/calculation$OSLIBEXT $PWD/../../../Numbers_component/Implementations/Cpp/build/numbers$OSLIBEXT
