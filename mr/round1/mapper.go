package main

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

type Log struct {
	UserId    int32  `json:"user_id"`
	ItemId    int32  `json:"item_id"`
	SessionId string `json:"session_id"`
	Operation string `json:"operation"`
	TimeStamp int64  `json:"time_stamp"`
}

func Map() {
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			break
		}

		t := Log{}
		err = jsoniter.UnmarshalFromString(line, &t)
		if err != nil {
			continue
		}

		// filter out non-view data
		if t.Operation != "view" {
			continue
		}

		fmt.Printf("%d\t%d\t%d\n", t.UserId, t.TimeStamp, t.ItemId)
	}
}

func main() {
	Map()
}
