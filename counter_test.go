package counter_test

import (
	"testing"

	"github.com/st3fan/counter"
)

func Test_Items(t *testing.T) {
	c := counter.New[string]()
	c.AddMany("1", "2", "2", "3", "3", "3", "4", "4", "4", "4")

	items := c.Items()
	if len(items) != 4 {
		t.Fail()
	}
}

func Test_MostCommon(t *testing.T) {
	c := counter.New[string]()
	c.AddMany("1", "2", "2", "3", "3", "3", "4", "4", "4", "4")

	most := c.MostCommon()
	if len(most) != 4 {
		t.Fail()
	}

	if most[0].Item != "4" || most[0].Count != 4 {
		t.Fail()
	}
}

func Test_LeastCommon(t *testing.T) {
	c := counter.New[string]()
	c.AddMany("1", "2", "2", "3", "3", "3", "4", "4", "4", "4")

	least := c.LeastCommon()
	if len(least) != 4 {
		t.Fail()
	}

	if least[0].Item != "1" || least[0].Count != 1 {
		t.Fail()
	}
}

func Test_Something(t *testing.T) {
	type thing struct{ value int }
	c := counter.New[thing]()
	c.Add(thing{10})
	c.Add(thing{10})
	c.Add(thing{10})
	c.Add(thing{20})
	c.Add(thing{25})
	c.Add(thing{25})

	if c.Length() != 3 {
		t.Fail()
	}

	mc := c.MostCommon()
	if len(mc) != 3 {
		t.Fail()
	}

	expected := thing{value: 10}
	if mc[0].Count != 3 || mc[0].Item != expected {
		t.Fail()
	}
}
