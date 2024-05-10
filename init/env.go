package main

import (
	"log"
	"os"
	"strconv"
)

func getenvStr(key string) string {
	res := os.Getenv(key)
	if res == "" {
		log.Fatal("the variable is empty")
	}

	return res
}

func getenvInt(key string) int {
	temp := getenvStr(key)

	res, err := strconv.Atoi(temp)

	if err != nil {
		log.Fatal(err)
	}
	return res
}
