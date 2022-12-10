package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("os open: %s", err.Error())
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("io read all: %s", err.Error())
	}

	return fileContent
}

func inputLines(path string) []string {
	return strings.Split(string(readFile(path)), "\n")
}

func mustInt(i string) int {
	num, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("atoi: %s", err.Error())
	}

	return num
}
