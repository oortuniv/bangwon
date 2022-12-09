#!/bin/bash

TARGET="test.sh"

PID = ps -ef | grep $TARGET | awk '{print $2}'

kill -9 $PID

exit $?