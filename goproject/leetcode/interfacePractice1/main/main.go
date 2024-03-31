package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (heroSlice HeroSlice) Len() int {
	return len(heroSlice)
}

func (heroSlice HeroSlice) Less(i, j int) bool {
	return heroSlice[i].Age < heroSlice[j].Age
}

func (heroSlice HeroSlice) Swap(i, j int) {
	tmp := heroSlice[i]
	heroSlice[i] = heroSlice[j]
	heroSlice[j] = tmp
}

func main() {
	var heroSlice HeroSlice = HeroSlice(make([]Hero, 0))
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	fmt.Println(heroSlice) // [{hero55 65} {hero98 91} {hero25 1} {hero71 18} {hero24 65} {hero16 24} {hero71 16} {hero61 47} {hero63 54} {hero27 30}]
	sort.Sort(heroSlice)   // 虽然起了别名但是还是引用类型
	fmt.Println(heroSlice) // [{hero25 1} {hero71 16} {hero71 18} {hero16 24} {hero27 30} {hero61 47} {hero63 54} {hero55 65} {hero24 65} {hero98 91}]
	// 证明HeroSlice是引用类型
	heroSlice.Swap(0, 1)
	fmt.Println(heroSlice) // [{hero25 16} {hero71 1} {hero71 18} {hero16 24} {hero27 30} {hero61 47} {hero63 54} {hero55 65} {hero24 65} {hero98 91}]
}
