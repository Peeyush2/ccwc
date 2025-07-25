package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	flags = "-c -l -w -m"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var dat = []byte{}

	allArgs := os.Args

	execPath, err := os.Executable()
	check(err)

	var execFile = filepath.Join(filepath.Dir(execPath), "main.exe")

	var otherArgs []string
	fileRef := ""

	allFlags := strings.Split(flags, " ")

	for _, val := range allArgs {
		isFlag := false
		for _, flag := range allFlags {
			if strings.ToLower(val) == flag {
				otherArgs = append(otherArgs, flag)
				isFlag = true
				continue
			}
		}
		if !isFlag && val != execFile {
			fileRef = val
		}
	}

	if len(fileRef) == 0 {
		reader := bufio.NewReader(os.Stdin)
		data, err := io.ReadAll(reader)
		check(err)
		dat = data
	} else {
		data, err := os.ReadFile(fileRef)
		check(err)
		dat = data
	}

	if len(otherArgs) == 0 {
		otherArgs = []string{"-c", "-l", "-w"}
	}

	var result string = ""

	for x := range otherArgs {
		result += " "
		result += handleFlags(otherArgs[x], dat)
	}

	result += " " + fileRef

	fmt.Println(result)

}

func handleFlags(flagValue string, dat []byte) string {
	switch flagValue {
	case "-c":
		return strconv.Itoa(len(dat))
	case "-l":
		lines := strings.Split(string(dat), "\n")
		return strconv.Itoa((len(lines) - 1))
	case "-w":
		words := strings.Fields(string(dat))
		return strconv.Itoa(len(words))
	case "-m":
		return strconv.Itoa(utf8.RuneCount(dat))
	default:
		return ""
	}
}
