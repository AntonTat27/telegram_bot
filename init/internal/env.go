package internal

import (
	"log"
	"os"
	"strconv"
)

func GetEnvStr(key string) string {
	res := os.Getenv(key)
	if res == "" {
		log.Fatalf("the variable '%s' is empty", key)
	}

	return res
}

func GetEnvInt(key string) int {
	temp := GetEnvStr(key)

	res, err := strconv.Atoi(temp)

	if err != nil {
		log.Fatal(err)
	}
	return res
}
