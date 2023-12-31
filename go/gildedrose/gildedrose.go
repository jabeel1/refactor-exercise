package gildedrose

import "fmt"

const (
	// Known items
	AgedBrieName        = "Aged Brie"
	BackStagePassesName = "Backstage passes to a TAFKAL80ETC concert"
	SulfurasName        = "Sulfuras, Hand of Ragnaros"

	maximumQuality = 50
	minimumQuality = 0
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) IncrementQuality() {
	if i.Quality < maximumQuality {
		i.Quality++
	}
}

func (i *Item) DecrementQuality() {
	if i.Quality > minimumQuality {
		i.Quality--
	}
}

type ItemIterator interface {
	Update()
	GetItem() Item
}

// AgedBrie is an item wrapper which increments over time and
// starts to double increment when the sell in is less than zero.
type AgedBrie struct {
	Item
}

func (a *AgedBrie) Update() {
	a.IncrementQuality()
	if a.SellIn <= 0 {
		a.IncrementQuality()
	}

	a.SellIn--
}

func (a *AgedBrie) GetItem() Item {
	return a.Item
}

// BackStagePasses represents a special item which depreciates over
// time and has quality boosters as the sell in decreases.
type BackStagePasses struct {
	Item
}

func (b *BackStagePasses) Update() {
	if b.SellIn <= 0 {
		b.Quality = 0
		b.SellIn--
		return
	}

	b.IncrementQuality()
	if b.SellIn < 11 {
		b.IncrementQuality()
	}

	if b.SellIn < 6 {
		b.IncrementQuality()
	}

	b.SellIn--
}

func (b *BackStagePasses) GetItem() Item {
	return b.Item
}

// Sulfuras represents an item which does not depreciate
// in quality or sell in over time.
type Sulfuras struct {
	Item
}

func (s *Sulfuras) Update() {}

func (s *Sulfuras) GetItem() Item {
	return s.Item
}

// Default represents an item which has no special logic.
// It is the base case.
type Default struct {
	Item
}

func (d *Default) Update() {
	d.DecrementQuality()
	if d.SellIn <= 0 {
		d.DecrementQuality()
	}

	d.SellIn--
}

func (d *Default) GetItem() Item {
	return d.Item
}

// NewItem creates a new ItemIterator with the provided fields.
func NewItem(name string, sellIn int, quality int) ItemIterator {
	underlyingItem := Item{
		Name:    name,
		Quality: quality,
		SellIn:  sellIn,
	}

	switch name {
	case AgedBrieName:
		return &AgedBrie{
			underlyingItem,
		}
	case BackStagePassesName:
		return &BackStagePasses{
			underlyingItem,
		}
	case SulfurasName:
		return &Sulfuras{
			underlyingItem,
		}
	}

	return &Default{
		underlyingItem,
	}
}

// NewWrappedItems converts the old list of items to a new slice of
// item iterators that can be used with the new UpdateQuality function.
func NewWrappedItems(items []*Item) []ItemIterator {
	iters := make([]ItemIterator, 0, len(items))
	for _, item := range items {
		iters = append(iters, NewItem(item.Name, item.SellIn, item.Quality))
	}

	return iters
}

// UpdateQuality is a bridging function to the new and old updateQuality functions.
// It takes any type in but is specifically looking for []*Item and []ItemIterator.
// Any other type provided will cause the function to panic.
func UpdateQuality(items any) {
	switch t := items.(type) {
	case []*Item:
		updateQualityV1(t)
	case []ItemIterator:
		updateQualityV2(t)
	default:
		panic(fmt.Sprintf("invalid type provided: %t", t))
	}
}

func updateQualityV2(items []ItemIterator) {
	for _, item := range items {
		item.Update()
	}
}

// updateQualityV1 is the legacy method of updating the quality of items
func updateQualityV1(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}

}
