package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	timestampThreshold = flag.Int("timestampThreshold", 0, "timestamp threshold")
)

type Item struct {
	userId    string
	itemId    string
	timeStamp int
}

func PrintAssociation(item1 Item, item2 Item, mp map[string]bool) {
	record := strings.Join([]string{item1.itemId, item2.itemId}, "\t")
	if !mp[record] {
		fmt.Println(record)
	} else {
		mp[record] = true
	}
}

func Reduce() {
	input := bufio.NewScanner(os.Stdin)

	prevUserId := ""
	clickedItems := make([]Item, 0)
	mp := make(map[string]bool)

	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, "\t")
		userId := parts[0]
		timeStamp, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		curItem := Item{
			userId:    userId,
			timeStamp: timeStamp,
			itemId:    parts[2],
		}

		// deal with another user
		if userId != prevUserId {
			clickedItems = make([]Item, 0)
			mp = make(map[string]bool)
		}

		// FIFO queue to discard old click events
		for len(clickedItems) > 0 && *timestampThreshold > 0 && curItem.timeStamp-clickedItems[0].timeStamp > *timestampThreshold {
			clickedItems = clickedItems[1:]
		}
		for _, clickedItem := range clickedItems {
			PrintAssociation(curItem, clickedItem, mp)
			PrintAssociation(clickedItem, curItem, mp)
		}

		prevUserId = userId
		clickedItems = append(clickedItems, curItem)
	}
}

func main() {
	Reduce()
}
