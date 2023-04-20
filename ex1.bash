#!/bin/bash
rm -f output1.bin output2.bin

./bin-create -hex FF -size 1048576 -file output1.bin
./bin-fill -file output1.bin -start 256 -end 512 -char 11

./bin-create -hex FF -size 1024 -file output2.bin
./bin-cp -source output1.bin -dest output2.bin -source-offset 256 -dest-offset 16 -length 64
