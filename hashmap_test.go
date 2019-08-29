package hashmap

import (
	"testing"
	"fmt"
)

func TestInit(t *testing.T) {
	hm := Init(10)
	fmt.Println(hm)
}

func TestHashMap_Get(t *testing.T) {
	hm := Init(10)
	hm.Set("ccc", "c")
	fmt.Println(hm.Get("ccc"))
	hm.Set("ccc", "d")
	fmt.Println(hm.Get("ccc"))
	hm.Print()
}

func TestHashMap_Set(t *testing.T) {
	hm := Init(10)
	hm.Set("ccc", "c")
	hm.Set("c", "s")
	hm.Set("cccc", "ss")
	hm.Set("fuck", "you")
	fmt.Println(hm)
	hm.Print()
	fmt.Println(hm.Get("fuck"))
	fmt.Println(hm.Get("ccc"))
	fmt.Println(hm.Get("c"))
	fmt.Println(hm.Get("cccc"))
	fmt.Println(hm.Get("ss"))
}

func TestHashMap_CalculateHashIndex(t *testing.T) {
	hm := Init(10)
	var c rune = 'a'
	for {
		fmt.Println(hm.CalculateHashIndex(string(c)))
		c++
		if c == 'z'{
			break
		}
	}
}
