package counter

import (
	"sort"
)

type Counter[T comparable] struct {
	items map[T]uint
}

type Count[T comparable] struct {
	Item  T
	Count uint
}

type MostCommonCounts[T comparable] []Count[T]

func (c MostCommonCounts[T]) Len() int {
	return len(c)
}

func (c MostCommonCounts[T]) Less(i, j int) bool {
	return c[i].Count > c[j].Count
}

func (c MostCommonCounts[T]) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type LeastCommonCounts[T comparable] []Count[T]

func (c LeastCommonCounts[T]) Len() int {
	return len(c)
}

func (c LeastCommonCounts[T]) Less(i, j int) bool {
	return c[i].Count < c[j].Count
}

func (c LeastCommonCounts[T]) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func New[T comparable]() *Counter[T] {
	return &Counter[T]{items: map[T]uint{}}
}

func (c *Counter[T]) Add(item T) {
	c.items[item] += 1
}

func (c *Counter[T]) AddMany(items ...T) {
	for _, item := range items {
		c.items[item] += 1
	}
}

func (c *Counter[T]) Inc(item T, n uint) {
	c.items[item] += n
}

func (c *Counter[T]) Dec(item T, n uint) {
	if v, ok := c.items[item]; ok && v > 0 {
		c.items[item] -= n
	}
}

func (c *Counter[T]) Items() []T {
	keys := make([]T, len(c.items))
	i := 0
	for k := range c.items {
		keys[i] = k
		i++
	}
	return keys
}

func (c *Counter[T]) counts() []Count[T] {
	counts := make([]Count[T], len(c.items))
	i := 0
	for k, v := range c.items {
		counts[i] = Count[T]{Item: k, Count: v}
		i++
	}
	return counts
}

func (c *Counter[T]) MostCommon() []Count[T] {
	counts := c.counts()
	sort.Sort(MostCommonCounts[T](counts))
	return counts
}

func (c *Counter[T]) LeastCommon() []Count[T] {
	counts := c.counts()
	sort.Sort(LeastCommonCounts[T](counts))
	return counts
}

func (c *Counter[T]) Length() int {
	return len(c.items)
}

func (c *Counter[T]) Total() uint {
	var total uint
	for _, v := range c.items {
		total += v
	}
	return total
}

func (c *Counter[T]) Clear() {
	c.items = map[T]uint{}
}
