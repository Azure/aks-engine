#!/bin/bash

# should have the expected java MaxHeapSize configuration
MEMORY_LIMIT=100000000
set -x
docker run -m $MEMORY_LIMIT adoptopenjdk/openjdk11:jdk-11.0.2.9-alpine java -XX:+UnlockDiagnosticVMOptions -XX:+PrintCommandLineFlags -version > /tmp/java-runtime.out || exit 1
[ -s `cat /tmp/java-runtime.out | awk '{print $2}' | sed 's,.*=,,' | awk -v limit="$MEMORY_LIMIT" '0+$1 >= limit {print}'` ]
if [ $? -ne 0  ]; then
  echo "java should not have a MaxHeapSize greater than its container runtime allows"
  echo "showing the MaxHeapSize settings for a java runtime inside a container with a memory limit of ${MEMORY_LIMIT} bytes"
  cat /tmp/java-runtime.out
  exit 1
fi