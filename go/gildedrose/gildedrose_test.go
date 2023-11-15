package gildedrose_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Fuzz_UpdateQuality(f *testing.F) {
	testCases := []string{
		gildedrose.AgedBrieName,
		gildedrose.BackStagePassesName,
		gildedrose.SulfurasName,
		"Elixir of the Mongoose",
		"Conjured Mana Cake",
		"+5 Dexterity Vest",
	}

	for _, tt := range testCases {
		f.Add(tt)
	}

	f.Fuzz(func(t *testing.T, source string) {
		itemQuality := rand.Intn(100)
		itemSellIn := rand.Intn(100)
		items := []gildedrose.ItemIterator{
			gildedrose.NewItem(source, itemQuality, itemSellIn),
		}

		itemsCopy := make([]*gildedrose.Item, 0, len(items))
		for _, item := range items {
			iterItem := item.GetItem()
			itemsCopy = append(itemsCopy, &iterItem)
		}

		gildedrose.UpdateQuality(items)
		gildedrose.UpdateQuality(itemsCopy)

		if !isEqual(items, itemsCopy) {
			t.Log("Mismatch Items: \nNew Items:\n")

			for _, item := range items {
				t.Logf("%+v\n", item.GetItem())
			}

			fmt.Printf("Old Items:\n")
			for _, item := range itemsCopy {
				t.Logf("%+v\n", *item)
			}

			t.Errorf("Item Mismatch")
		}
	})
}

func isEqual(a []gildedrose.ItemIterator, b []*gildedrose.Item) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		itemA := a[i].GetItem()
		itemB := *b[i]

		if itemA.Name != itemB.Name || itemA.Quality != itemB.Quality || itemA.SellIn != itemB.SellIn {
			return false
		}
	}

	return true
}
