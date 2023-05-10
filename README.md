# sammtcbn's Binary Utility

## bin-create

Here's a tool for generating a binary file filled with a specified hexadecimal value.

Usage:
```
./bin-create -hex <hexadecimal value> -size <file size> -file <file path> -bs <block size>
```

You can use this command to generate a binary file filled with the specified hexadecimal value:

```
$ go run main.go -hex "FF" -size 4096 -file "output.bin"
```

This command will generate a 4KB binary file filled with the hexadecimal value "FF" at the path "output.bin".

For example, to generate a 1GB binary file filled with the hexadecimal value "FF" using a block size of 1MB and save it as "output.bin" in the current directory, run the following command:

```
$ ./bin-create -hex "FF" -size 1073741824 -file "output.bin" -bs 1048576
```

This will write 1MB blocks of "FF" to the file until the file size is reached.

## bin-cp

Here's a tool written in Go that can copy the contents of a source file to a destination file, with the ability to specify the start offset and length of the source file and the start offset of the destination file.

This tool accepts the following command line arguments:

* -source: specifies the path of the source file.
* -dest: specifies the path of the destination file.
* -source-offset: specifies the start offset of the source file.
* -dest-offset: specifies the start offset of the destination file.
* -length: specifies the length of the data to be copied.

To use the tool, you can run it like this:

```
go run bin-cp.go -source /path/to/source/file -dest /path/to/destination/file -source-offset 100 -dest-offset 200 -length 50
```

This will copy 50 bytes of data starting from the 100th byte of the source file to the destination file starting at the 200th byte. Note that if the destination file does not exist, the tool will create it automatically.

## bin-fill

This tool can be used with the following command-line flags:

* -file: Specify the file name.
* -start: Specify the start position, starting from 0.
* -end: Specify the end position, not including that position. If set to -1, it means to the end of the file.
* -char: Specify the hex character to replace, which must be 2 characters.

For example, if you want to replace all binary between position 10 and 20 in the file example.bin with the hex character 0A, you can use the following command:

```
$ go run bin-fill.go -file example.bin -start 10 -end 20 -char 0A
```

This will replace all binary between position 10 and 20 in the file example.bin with the hex character 0A.

# build executable

To build all tools, you can run this:
```
./build.bash
```
