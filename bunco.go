package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	// 1 Billion
	trials int = 1000000000
)

func main(){
	//	0	1	2	3	4
	// $0, $1, $2, $5, $21
	results := [5]int{}
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i<trials; i++{
		rolls := [3]int{gen.Intn(6)+1, gen.Intn(6)+1, gen.Intn(6)+1}
		ones := 0
		// Counting number of 1s
		for _,pl := range rolls{
			if pl == 1{
				ones++
			}
		}
		if ones == 1{
			results[1]++
		}else if ones == 2{
			results[2]++
		}else if ones == 3{
			results[4]++
		}else if (rolls[0] == rolls[1]) && (rolls[0] == rolls[2]){
			results[3]++
		}else{
			results[0]++
		}
	}
	fmt.Printf("$0:\tTotal: %d\tPercent: %f\n", rolls[0], rolls[0]/trials)
	fmt.Printf("$1:\tTotal: %d\tPercent: %f\n", rolls[1], rolls[1]/trials)
	fmt.Printf("$2:\tTotal: %d\tPercent: %f\n", rolls[2], rolls[2]/trials)
	fmt.Printf("$5:\tTotal: %d\tPercent: %f\n", rolls[3], rolls[3]/trials)
	fmt.Printf("$21:\tTotal: %d\tPercent: %f\n", rolls[4], rolls[4]/trials)

}