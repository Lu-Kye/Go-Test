#!/usr/bin/env bash

p=$(cd `dirname $0`; pwd)
bin=$p/bin

cd $p
./build.sh

cd $bin
./test	
