package main

import (
    "flag"
    "fmt"
    "io"
    "os"
)

const version = "1.0"

func main() {
    var (
        sourcePath   string
        destPath     string
        sourceOffset int64
        destOffset   int64
        length       int64
        showVersion  bool
    )

    flag.StringVar(&sourcePath,  "source",        "",   "source file path")
    flag.StringVar(&destPath,    "dest",          "",   "destination file path")
    flag.Int64Var(&sourceOffset, "source-offset", 0,    "source offset")
    flag.Int64Var(&destOffset,   "dest-offset",   0,    "destination offset")
    flag.Int64Var(&length,       "length",        0,    "length of data to copy")
    flag.BoolVar(&showVersion,   "version",       false,"display version number")
    flag.Parse()

    if showVersion {
        fmt.Println("bin-cp", version)
        return
    }

    if sourcePath == "" || destPath == "" || length == 0 {
        fmt.Println("Please specify source file path, destination file path, and length")
        return
    }

    sourceFile, err := os.Open(sourcePath)
    if err != nil {
        fmt.Println("Error opening source file:", err)
        return
    }
    defer sourceFile.Close()

    destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error opening destination file:", err)
        return
    }
    defer destFile.Close()

    _, err = sourceFile.Seek(sourceOffset, io.SeekStart)
    if err != nil {
        fmt.Println("Error seeking source file:", err)
        return
    }

    _, err = destFile.Seek(destOffset, io.SeekStart)
    if err != nil {
        fmt.Println("Error seeking destination file:", err)
        return
    }

    buf := make([]byte, length)
    _, err = io.ReadFull(sourceFile, buf)
    if err != nil {
        fmt.Println("Error reading source file:", err)
        return
    }

    _, err = destFile.Write(buf)
    if err != nil {
        fmt.Println("Error writing destination file:", err)
        return
    }

    fmt.Println("Copy operation completed successfully")
}
