package main

import (
	"bufio"
	"github.com/go-redis/redis"
	"io"
	"os"
	"strings"
)

func Process(source io.Reader) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	input := bufio.NewScanner(source)
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, "\t")
		key := parts[0]
		value := make([]string, 0)
		for i := 1; i < len(parts); i += 2 {
			value = append(value, parts[i]+"-"+parts[i+1])
		}
		client.RPush(key, value)
	}
}

func main() {
	file, _ := os.Open(os.Args[1])
	Process(file)
}
