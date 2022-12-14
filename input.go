package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

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

func mustUnmarshalJSON(data []byte, v any) any {
	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatalf("json unmarshal: %s", err.Error())
	}

	return v
}

func mustMarshalJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("json marshal: %s", err.Error())
	}

	return string(b)
}
