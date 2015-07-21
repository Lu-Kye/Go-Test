#!/usr/bin/env bash

p=$(cd `dirname $0`; pwd)
src=$p/src
test=$src/test

cd $test
go install
