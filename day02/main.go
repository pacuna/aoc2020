package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// part I
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var valid int
	for scanner.Scan() {
		var min, max int
		var c int
		var w string
		_, err = fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &min, &max, &c, &w)
		if err != nil {
			log.Fatalf("error parsing %s: %v\n", scanner.Text(), err)
		}

		count := strings.Count(w, string(c))
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Println(valid)
	f.Close()

	// part II
	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(f)
	valid = 0
	for scanner.Scan() {
		var idx1, idx2 int
		var c int
		var w string
		_, err = fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &idx1, &idx2, &c, &w)
		if err != nil {
			log.Fatalf("error parsing %s: %v\n", scanner.Text(), err)
		}
		if (int(w[idx1-1]) == c || int(w[idx2-1]) == c) && w[idx1-1] != w[idx2-1] {
			valid++
		}

	}
	fmt.Println(valid)
}
