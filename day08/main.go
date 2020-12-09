package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Instruction struct {
	cmd string
	val int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	instructions := []Instruction{}
	for scanner.Scan() {
		ins := Instruction{}
		fmt.Sscanf(scanner.Text(), "%s %d", &ins.cmd, &ins.val)
		instructions = append(instructions, ins)
	}

	var res int
	var cur Instruction

	var i int
	seen := make(map[int]int)
	for {
		cur = instructions[i]
		seen[i]++
		if seen[i] == 2 {
			fmt.Printf("breaking with acc: %d\n", res)
			break
		}
		if cur.cmd == "nop" {
			i++
		} else if cur.cmd == "acc" {
			res += cur.val
			i++
		} else if cur.cmd == "jmp" {
			i += cur.val
		}

	}
}
