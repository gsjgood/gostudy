package main

import "fmt"

//分金币 50个金币
var (
	gold = 50
	user = []string{
		"e", "c", "c", "c", "c",
	}
	dispathCoin = make(map[string]int, len(user))
)

func redundant() int {

	var yingshe = make(map[string]int, 8)
	yingshe["e"] = 1
	yingshe["E"] = 1
	yingshe["i"] = 2
	yingshe["I"] = 2
	yingshe["o"] = 3
	yingshe["O"] = 3
	yingshe["u"] = 4
	yingshe["U"] = 4
	for _, name := range user {
		//fmt.Println(name)
		gold_tmp := 0
		for _, byt := range name {

			//fmt.Println(string(byt))
			gold_tmp += yingshe[string(byt)]
		}
		if gold == 0 {
			break
		}
		dispathCoin[name] = gold_tmp
		gold -= gold_tmp
	}

	return gold
}

func main() {
	golds := redundant()
	fmt.Println(golds)
}
