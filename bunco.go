package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
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
	outf,_ := os.Create("data/graph.dat")
	defer outf.Close()
	outf.Write([]byte{})
	outf.Write([]byte(fmt.Sprintf("0\t%f\n", float64(results[0])/float64(trials))))
	outf.Write([]byte(fmt.Sprintf("1\t%f\n", float64(results[1])/float64(trials))))
	outf.Write([]byte(fmt.Sprintf("2\t%f\n", float64(results[2])/float64(trials))))
	outf.Write([]byte(fmt.Sprintf("5\t%f\n", float64(results[3])/float64(trials))))
	outf.Write([]byte(fmt.Sprintf("21\t%f\n", float64(results[4])/float64(trials))))
	fmt.Printf("$0:\tTotal: %d\tFrequency: %f\n", results[0], float64(results[0])/float64(trials))
	fmt.Printf("$1:\tTotal: %d\tFrequency: %f\n", results[1], float64(results[1])/float64(trials))
	fmt.Printf("$2:\tTotal: %d\tFrequency: %f\n", results[2], float64(results[2])/float64(trials))
	fmt.Printf("$5:\tTotal: %d\tFrequency: %f\n", results[3], float64(results[3])/float64(trials))
	fmt.Printf("$21:\tTotal: %d\tFrequency: %f\n", results[4], float64(results[4])/float64(trials))

}