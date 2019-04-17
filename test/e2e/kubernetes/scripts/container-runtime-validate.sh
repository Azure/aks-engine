#!/bin/bash

# should have the expected java MaxHeapSize configuration
set -x
docker run -m 100m adoptopenjdk/openjdk11:jdk-11.0.2.9-alpine java -XX:+UnlockDiagnosticVMOptions -XX:+PrintCommandLineFlags -version | grep 'MaxHeapSize=52428800' || exit 1