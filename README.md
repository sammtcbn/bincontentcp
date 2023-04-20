Here's a tool written in Go that can copy the contents of a source file to a destination file, with the ability to specify the start offset and length of the source file and the start offset of the destination file.

This tool accepts the following command line arguments:

* -source: specifies the path of the source file.
* -dest: specifies the path of the destination file.
* -source-offset: specifies the start offset of the source file.
* -dest-offset: specifies the start offset of the destination file.
* -length: specifies the length of the data to be copied.

To use the tool, you can run it like this:

```
go run bincontentcp.go -source /path/to/source/file -dest /path/to/destination/file -source-offset 100 -dest-offset 200 -length 50
```

This will copy 50 bytes of data starting from the 100th byte of the source file to the destination file starting at the 200th byte. Note that if the destination file does not exist, the tool will create it automatically.
