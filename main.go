package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	var dat []byte

	lenArgs := len(os.Args)

	allFlags := strings.Split(flags, " ")

	isCorreectLink := true

	for _, flag := range allFlags {
		if strings.Contains(os.Args[lenArgs-1], flag) {
			isCorreectLink = false
			break
		}
	}

	fmt.Println("Arguments:", os.Args)
	fmt.Println(isCorreectLink)

	var fileRef string = ""
	var otherArgs []string

	if isCorreectLink {
		fileRef = os.Args[lenArgs-1]
		if lenArgs >= 2 {
			otherArgs = os.Args[1 : lenArgs-1]
		}
	} else {
		fileRef = ""
		if lenArgs >= 1 {
			otherArgs = os.Args[1:lenArgs]
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
