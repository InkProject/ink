package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

const (
    CLR_W = ""
    CLR_R = "\x1b[31;1m"
    CLR_G = "\x1b[32;1m"
    CLR_B = "\x1b[34;1m"
)

const (
    LOG = "LOG"
    ERR = "ERR"
)

const (
    STD_FORMAT = "2006-01-02 15:04:05"
)

// Print colorful log
func Log(method string, info interface{}) {
    // fmt.Printf("%s%s\n%s", color, info, "\x1b[0m")
    method = ""
    fmt.Printf("%s\n", info)
}

func Fatal(info interface{}) {
    Log(ERR, info)
    os.Exit(1)
}

func ParseDate(dateStr string) time.Time {
    format := fmt.Sprintf(STD_FORMAT)
    date, err := time.ParseInLocation(format, dateStr, time.Now().Location())
    if err != nil {
        Fatal(err.Error())
    }
    return date.Local()
}

func Exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return false
}

// Copy folder and file
// Refer to https://www.socketloop.com/tutorials/golang-copy-directory-including-sub-directories-files

func CopyFile(source string, dest string) (err error) {
    sourcefile, err := os.Open(source)
    if err != nil {
        return err
    }
    defer sourcefile.Close()
    destfile, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destfile.Close()
    _, err = io.Copy(destfile, sourcefile)
    if err == nil {
        sourceinfo, err := os.Stat(source)
        if err != nil {
            err = os.Chmod(dest, sourceinfo.Mode())
        }
    }
    return
}

func CopyDir(source string, dest string) (err error) {
    sourceinfo, err := os.Stat(source)
    if err != nil {
        return err
    }
    err = os.MkdirAll(dest, sourceinfo.Mode())
    if err != nil {
        return err
    }
    directory, _ := os.Open(source)
    objects, err := directory.Readdir(-1)
    for _, obj := range objects {
        sourcefilepointer := source + "/" + obj.Name()
        destinationfilepointer := dest + "/" + obj.Name()
        if obj.IsDir() {
            err = CopyDir(sourcefilepointer, destinationfilepointer)
            if err != nil {
                fmt.Println(err)
            }
        } else {
            err = CopyFile(sourcefilepointer, destinationfilepointer)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
    return
}
