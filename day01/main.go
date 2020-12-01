package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	s := []int{}
	d := make(map[int]int)
	for scanner.Scan() {
		var x int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &x)
		if err != nil {
			log.Printf("error parsing %s: %v\n", scanner.Text(), err)
			continue
		}
		if x < 0 || x > 2020 {
			continue
		} else {
			s = append(s, x)
		}
	}

	// part I
	for i, el := range s {
		d[el] = i
	}

	for k, _ := range d {
		if val, ok := d[2020-k]; ok {
			fmt.Println(k, s[val], k*s[val])
			break
		}
	}

	// part II
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(s); k++ {
				if s[i]+s[j]+s[k] == 2020 {
					fmt.Println(s[i], s[j], s[k], s[i]*s[j]*s[k])
					break
				}
			}
		}
	}

}
