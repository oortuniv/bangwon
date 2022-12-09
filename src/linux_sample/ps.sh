#!/bin/bash

TARGET="test.sh"

PID = ps -ef | grep $TARGET | awk '{print $2}'

EXIT_CODE = 1
if [ "$PID" != "" ]; then
  EXIT_CODE = 0
fi

exit $EXIT_CODE