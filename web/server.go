package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	MaxLen = 50
)

type Response struct {
	ItemId string `json:"item_id"`
	Degree int    `json:"degree"`
}

var client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS-SERVER") + ":" + os.Getenv("REDIS-PORT"),
	Password: "",
	DB:       0,
})

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["ItemId"]
	minDegree, _ := strconv.Atoi(vars["minDegree"])

	list, _ := client.LRange(itemId, 0, -1).Result()
	result := make([]Response, 0)
	for i, item := range list {
		parts := strings.Split(item, "-")
		curDegree, _ := strconv.Atoi(parts[2])
		// if already more than 50 items or curDegree smaller than minDegree, then break
		// note that associated items stored in Redis are in descending of degree
		if i >= MaxLen || curDegree < minDegree {
			break
		}
		result = append(result, Response{ItemId: parts[0], Degree: curDegree})
	}

	_ = json.NewEncoder(w).Encode(result)
}

func handleRequests(addr string) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/vav", handler).Methods(http.MethodGet).Queries("item_id", "{ItemId}", "min_degree", "{minDegree}")
	log.Fatal(http.ListenAndServe(addr, myRouter))
}

func main() {
	handleRequests(":" + os.Args[1])
}
