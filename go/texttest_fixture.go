package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	iterItems := gildedrose.NewWrappedItems(items)
	copiedItems := make([]*gildedrose.Item, 0, len(items))
	for _, item := range items {
		copiedItems = append(copiedItems, &gildedrose.Item{
			Name:    item.Name,
			Quality: item.Quality,
			SellIn:  item.SellIn,
		})
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("Name, SellIn, Quality")
		for i := 0; i < len(iterItems); i++ {
			fmt.Printf("new:    %v\n", iterItems[i].GetItem())
			fmt.Printf("copied: %v\n", copiedItems[i])
		}
		fmt.Println("")
		gildedrose.UpdateQuality(iterItems)
		gildedrose.UpdateQuality(copiedItems)
	}
}
