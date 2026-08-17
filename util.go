package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	CLR_W = ""
	CLR_R = "\x1b[31;1m"
	CLR_G = "\x1b[32;1m"
	CLR_B = "\x1b[34;1m"
	CLR_Y = "\x1b[33;1m"
)

const (
	DATE_FORMAT               = "2006-01-02 15:04:05"
	DATE_FORMAT_WITH_TIMEZONE = "2006-01-02 15:04:05 -0700"
)

var exitCode int

// Print log
func Log(info any) {
	fmt.Printf("%s\n", info)
}

// Print warning log
func Warn(info any) {
	if runtime.GOOS == "windows" {
		fmt.Printf("WARNING: %s\n", info)
	} else {
		fmt.Printf("%s%s\n%s", CLR_Y, info, "\x1b[0m")
	}
}

// Print error log
func Error(info any) {
	if runtime.GOOS == "windows" {
		fmt.Printf("ERR: %s\n", info)
	} else {
		fmt.Printf("%s%s\n%s", CLR_R, info, "\x1b[0m")
	}
	exitCode = 1
}

// Print error log and exit
func Fatal(info any) {
	Error(info)
	os.Exit(1)
}

// Parse date by std date string
func ParseDate(dateStr string) time.Time {
	date, err := time.Parse(fmt.Sprint(DATE_FORMAT_WITH_TIMEZONE), dateStr)
	if err != nil {
		date, err = time.ParseInLocation(fmt.Sprint(DATE_FORMAT), dateStr, time.Now().Location())
		if err != nil {
			Fatal(err.Error())
		}
	}
	return date
}

// Check file if exist
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Check file if is directory
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

func walkSymlinks(root string, fn filepath.WalkFunc) error {
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		return err
	}
	visitedDirs := make([]os.FileInfo, 0)

	var walkFn filepath.WalkFunc
	walkFn = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fn(path, info, err)
		}

		if info.IsDir() {
			for _, visited := range visitedDirs {
				if os.SameFile(info, visited) {
					return filepath.SkipDir
				}
			}
			visitedDirs = append(visitedDirs, info)
		}

		if err := fn(path, info, nil); err != nil {
			return err
		}

		if info.Mode()&os.ModeSymlink == 0 {
			return nil
		}

		linkedInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if !linkedInfo.IsDir() {
			return nil
		}
		for _, visited := range visitedDirs {
			if os.SameFile(linkedInfo, visited) {
				return nil
			}
		}

		linkedPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}
		return filepath.Walk(linkedPath, walkFn)
	}

	return filepath.Walk(root, walkFn)
}

// Copy folder and file
// Refer to https://www.socketloop.com/tutorials/golang-copy-directory-including-sub-directories-files
func CopyFile(source string, dest string) {
	sourcefile, err := os.Open(source)
	if err != nil {
		Fatal(err.Error())
	}
	destfile, err := os.Create(dest)
	if err != nil {
		Fatal(err.Error())
	}
	defer func() {
		if err := destfile.Close(); err != nil {
			Fatal(err.Error())
		}
	}()
	defer wg.Done()
	_, err = io.Copy(destfile, sourcefile)
	if err != nil {
		Fatal(err.Error())
	}
	sourceinfo, err := os.Stat(source)
	if err != nil {
		Fatal(err.Error())
	}
	err = os.Chmod(dest, sourceinfo.Mode())
	if err != nil {
		Fatal(err.Error())
	}
	if err := sourcefile.Close(); err != nil {
		Fatal(err.Error())
	}
}

func CopyDir(source string, dest string) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		Fatal(err.Error())
	}
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		Fatal(err.Error())
	}
	directory, err := os.Open(source)
	if err != nil {
		Fatal(err.Error())
	}
	defer func() {
		if err := directory.Close(); err != nil {
			Fatal(err.Error())
		}
	}()
	defer wg.Done()
	objects, err := directory.Readdir(-1)
	if err != nil {
		Fatal(err.Error())
	}
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()
		if obj.IsDir() {
			wg.Add(1)
			CopyDir(sourcefilepointer, destinationfilepointer)
		} else {
			wg.Add(1)
			go CopyFile(sourcefilepointer, destinationfilepointer)
		}
	}
}
