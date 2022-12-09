#!/bin/bash

nohup ./test.sh 1>dev>null 2>&1 &

exit $?