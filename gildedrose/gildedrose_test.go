package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

type test struct {
	item     *gildedrose.Item
	expected gildedrose.Item
}

func TestNormalItem(t *testing.T) {
	tests := []test{
		{
			item:     &gildedrose.Item{"Elixir of the Mongoose", 2, 10},
			expected: gildedrose.Item{"Elixir of the Mongoose", 1, 9},
		},
		{
			item:     &gildedrose.Item{"Elixir of the Mongoose", 0, 10},
			expected: gildedrose.Item{"Elixir of the Mongoose", -1, 8},
		},
	}

	for _, ts := range tests {
		in := []*gildedrose.Item{ts.item}
		gildedrose.UpdateQuality(in)
		for _, out := range in {
			if out.Quality != ts.expected.Quality {
				t.Errorf("Quality: Expected %d but got %d ", ts.expected.Quality, out.Quality)
			}

		}
	}
}

func TestSulfuras(t *testing.T) {
	tests := []test{
		{
			item:     &gildedrose.Item{"Sulfuras, Hand of Ragnaros", 2, 80},
			expected: gildedrose.Item{"Sulfuras, Hand of Ragnaros", 1, 80},
		},
		{
			item:     &gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 80},
			expected: gildedrose.Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		},
	}

	for _, ts := range tests {
		in := []*gildedrose.Item{ts.item}
		gildedrose.UpdateQuality(in)
		for _, out := range in {
			if out.Quality != ts.expected.Quality {
				t.Errorf("Quality: Expected %d but got %d ", ts.expected.Quality, out.Quality)
			}

		}
	}
}

func TestBackstagePass(t *testing.T) {
	tests := []test{
		{
			item:     &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 10, 0},
			expected: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 9, 2},
		},
		{
			item:     &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 5, 0},
			expected: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 4, 3},
		},
		{
			item:     &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 0, 10},
			expected: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		},
		{
			item:     &gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 0, -1},
			expected: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		},
	}

	for _, ts := range tests {
		in := []*gildedrose.Item{ts.item}
		gildedrose.UpdateQuality(in)
		for _, out := range in {
			if out.Quality != ts.expected.Quality {
				t.Errorf("Quality: Expected %d but got %d ", ts.expected.Quality, out.Quality)
			}

		}
	}
}

func TestBrie(t *testing.T) {
	tests := []test{
		{
			item:     &gildedrose.Item{"Aged Brie", 2, 0},
			expected: gildedrose.Item{"Aged Brie", 1, 1},
		},
		{
			item:     &gildedrose.Item{"Aged Brie", 0, 0},
			expected: gildedrose.Item{"Aged Brie", -1, 2},
		},
		{
			item:     &gildedrose.Item{"Aged Brie", -1, 2},
			expected: gildedrose.Item{"Aged Brie", -2, 4},
		},
		{
			item:     &gildedrose.Item{"Aged Brie", 1, 50},
			expected: gildedrose.Item{"Aged Brie", 0, 50},
		},
		{
			item:     &gildedrose.Item{"Aged Brie", 1, -1},
			expected: gildedrose.Item{"Aged Brie", 0, 0},
		},
	}

	for _, ts := range tests {
		in := []*gildedrose.Item{ts.item}
		gildedrose.UpdateQuality(in)
		for _, out := range in {
			if out.Quality != ts.expected.Quality {
				t.Errorf("Quality: Expected %d but got %d ", ts.expected.Quality, out.Quality)
			}

		}
	}
}

func TestConjured(t *testing.T) {
	tests := []test{
		{
			item:     &gildedrose.Item{"Conjured", 2, 2},
			expected: gildedrose.Item{"Conjured", 1, 0},
		},
		{
			item:     &gildedrose.Item{"Conjured", 2, 0},
			expected: gildedrose.Item{"Conjured", 1, 0},
		},
		{
			item:     &gildedrose.Item{"Conjured", 0, 4},
			expected: gildedrose.Item{"Conjured", -1, 0},
		},
	}

	for _, ts := range tests {
		in := []*gildedrose.Item{ts.item}
		gildedrose.UpdateQuality(in)
		for _, out := range in {
			if out.Quality != ts.expected.Quality {
				t.Errorf("Quality: Expected %d but got %d ", ts.expected.Quality, out.Quality)
			}

		}
	}
}
