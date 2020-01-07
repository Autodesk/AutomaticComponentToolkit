#!/bin/bash

JnaJar="jna-5.5.0.jar"
Classpath=".:${JnaJar}:../../Bindings/Java/"
if [[ "$OSTYPE" == "linux-gnu" ]]; then
	Classpath=".:${JnaJar}:../../Bindings/Java/"
elif [[ "$OSTYPE" == "darwin"* ]]; then
	Classpath=".:${JnaJar}:../../Bindings/Java/"
elif [[ "$OSTYPE" == "cygwin" ]]; then
	Classpath=".;${JnaJar};../../Bindings/Java/"
elif [[ "$OSTYPE" == "msys" ]]; then
	Classpath=".;${JnaJar};../../Bindings/Java/"
elif [[ "$OSTYPE" == "win32" ]]; then
	Classpath=".;${JnaJar};../../Bindings/Java/"
else
	echo "Unknown system: "$OSTYPE
	exit 1
fi

echo "Download JNA"
wget http://repo1.maven.org/maven2/net/java/dev/jna/jna/5.5.0/jna-5.5.0.jar

echo "Compile Java bindings"
javac -classpath "${JnaJar}" ../../Bindings/Java/libprimes/*.java
echo "Compile Java example"
javac -classpath $Classpath LibPrimes_Example.java

echo "Execute example"
java -classpath $Classpath LibPrimes_Example
