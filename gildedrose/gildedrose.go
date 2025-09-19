package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type qualityUpdater interface {
	updateQuality()
}

type sulfuras struct{ *Item }

func (s *sulfuras) updateQuality() {
	s.SellIn -= 1
	s.Quality = 80
}

type agedBrie struct{ *Item }

func (a *agedBrie) updateQuality() {
	a.SellIn -= 1
	if a.SellIn < 0 {
		a.Quality += 2
	} else {
		a.Quality += 1
	}
	a.Quality = min(a.Quality, 50)
}

type backstagePasses struct{ *Item }

func (b *backstagePasses) updateQuality() {
	b.SellIn -= 1
	if b.SellIn < 0 {
		b.Quality = 0
		return
	}

	switch {
	case b.SellIn < 5:
		b.Quality += 3
	case b.SellIn < 10:
		b.Quality += 2
	default:
		b.Quality += 1
	}
}

type conjured struct{ *Item }

func (c *conjured) updateQuality() {
	c.SellIn -= 1
	if c.SellIn < 0 {
		c.Quality -= 4
	} else {
		c.Quality -= 2
	}
	c.Quality = min(max(c.Quality, 0), 50)
}

type normalItem struct{ *Item }

func (n *normalItem) updateQuality() {
	n.SellIn -= 1
	if n.SellIn < 0 {
		n.Quality -= 2
	} else {
		n.Quality -= 1
	}
	n.Quality = min(max(n.Quality, 0), 50)
}

func itemFactory(i *Item) qualityUpdater {
	switch n := i.Name; n {
	case "Aged Brie":
		return &agedBrie{i}
	case "Sulfuras, Hand of Ragnaros":
		return &sulfuras{i}
	case "Backstage passes to a TAFKAL80ETC concert":
		return &backstagePasses{i}
	case "Conjured":
		return &conjured{i}
	default:
		return &normalItem{i}
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		itemFactory(items[i]).updateQuality()
	}
}
