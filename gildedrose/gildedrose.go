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
	if s.Quality != 80 {
		s.Quality = 80
	}
}

type agedBrie struct{ *Item }

func (a *agedBrie) updateQuality() {
	a.SellIn -= 1
	if a.SellIn < 0 {
		a.Quality += 2
	} else {
		a.Quality += 1
	}
	if a.Quality > 50 {
		a.Quality = 50
	}
}

type backstagePasses struct{ *Item }

func (b *backstagePasses) updateQuality() {
	b.SellIn -= 1
	if b.SellIn < 0 {
		b.Quality = 0
		return
	}

	if b.SellIn < 5 {
		b.Quality += 3
	} else if b.SellIn < 10 {
		b.Quality += 2
	} else {
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
	if c.Quality > 50 {
		c.Quality = 50
	} else if c.Quality < 0 {
		c.Quality = 0
	}
}

type normalItem struct{ *Item }

func (n *normalItem) updateQuality() {
	n.SellIn -= 1
	if n.SellIn < 0 {
		n.Quality -= 2
	} else {
		n.Quality -= 1
	}
	if n.Quality > 50 {
		n.Quality = 50
	} else if n.Quality < 0 {
		n.Quality = 0
	}
}

func itemFactory(i *Item) (qu qualityUpdater) {
	switch n := i.Name; n {
	case "Aged Brie":
		qu = &agedBrie{i}
	case "Sulfuras, Hand of Ragnaros":
		qu = &sulfuras{i}
	case "Backstage passes to a TAFKAL80ETC concert":
		qu = &backstagePasses{i}
	case "Conjured":
		qu = &conjured{i}
	default:
		qu = &normalItem{i}
	}
	return
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		itemFactory(items[i]).updateQuality()
	}
}
