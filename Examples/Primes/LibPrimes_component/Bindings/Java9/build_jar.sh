#!/bin/bash

echo "Download JNA"
wget http://repo1.maven.org/maven2/net/java/dev/jna/jna/5.5.0/jna-5.5.0.jar

echo "Compile Java Bindings"
javac -classpath *.jar libprimes/*

echo "Create JAR"
jar cvf libprimes-1.2.0.jar libprimes
