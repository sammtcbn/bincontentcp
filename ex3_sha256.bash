#!/bin/bash
rm -f ex3_1M_rand.bin ex3_1M_rand.sha256.bin

./bin-create -randhex -size 1048576 -file ex3_1M_rand.bin -bs 1024
echo

echo ./bin-sha256 -file ex3_1M_rand.bin
./bin-sha256 -file ex3_1M_rand.bin
echo

./bin-sha256 -file ex3_1M_rand.bin -outbin ex3_1M_rand.sha256.bin
echo ./bin-sha256 -file ex3_1M_rand.bin -outbin ex3_1M_rand.sha256.bin
echo

echo hexdump -C -v ex3_1M_rand.sha256.bin
hexdump -C -v ex3_1M_rand.sha256.bin || exit 1
echo

echo sha256sum ex3_1M_rand.bin
sha256sum ex3_1M_rand.bin
echo

echo 'sha256sum ex3_1M_rand.bin | cut -d " " -f 1 | xxd -r -ps >  ex3_1M_rand.sha256.bin2'
sha256sum ex3_1M_rand.bin | cut -d " " -f 1 | xxd -r -ps >  ex3_1M_rand.sha256.bin2 || exit 1
echo

echo hexdump -C -v ex3_1M_rand.sha256.bin2
hexdump -C -v ex3_1M_rand.sha256.bin2 || exit 1
