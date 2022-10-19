package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	itemId string
}

type Association struct {
	associatedItem Item
	degree         int
}

type Associations []Association

func (a Associations) Len() int {
	return len(a)
}

func (a Associations) Less(i, j int) bool {
	return a[i].degree > a[j].degree
}

func (a Associations) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Deal(self Item, mp *map[Item]int) string {
	associations := make(Associations, 0)
	for k, v := range *mp {
		associations = append(associations, Association{associatedItem: k, degree: v})
	}
	// all associated items are sorted in descending order of degree
	sort.Sort(associations)

	// print current item and its associated item in the format of
	// itemId shopId associatedItem1 associatedShopId1 degree1 associatedItemId2 associatedShopId2 degree2 ...
	str := self.itemId
	for _, a := range associations {
		item := a.associatedItem
		str += "\t" + item.itemId + "\t" + strconv.Itoa(a.degree)
	}
	*mp = make(map[Item]int)
	return str
}

func Reduce() {
	input := bufio.NewScanner(os.Stdin)
	mp := make(map[Item]int)
	prevItemId := ""
	var self = Item{}
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, "\t")
		curItemId := parts[0]

		// deal with another item
		if curItemId != prevItemId && prevItemId != "" {
			fmt.Println(Deal(self, &mp))
		}

		associated := Item{itemId: parts[1]}
		self = Item{itemId: curItemId}
		prevItemId = curItemId
		mp[associated] += 1
	}
	// deal with remaining data
	fmt.Println(Deal(self, &mp))
}

func main() {
	Reduce()
}
