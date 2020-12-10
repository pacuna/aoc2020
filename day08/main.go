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

// return true if loops and last value of accumulator
func loops(instructions []Instruction) (bool, int) {
	var res int
	var cur Instruction

	var i int
	seen := make(map[int]int)
	for {
		if i == len(instructions) {
			return false, res
		}
		cur = instructions[i]
		seen[i]++
		if seen[i] == 2 {
			return true, res
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
	return false, res
}

func swap(instructions []Instruction) int {
	var res int
	for idx, ins := range instructions {
		if ins.cmd == "nop" {
			instructions[idx].cmd = "jmp"
			if l, acc := loops(instructions); !l {
				return acc
			} else {
				instructions[idx].cmd = "nop"
			}
		} else if ins.cmd == "jmp" {
			instructions[idx].cmd = "nop"
			if l, acc := loops(instructions); !l {
				return acc
			} else {
				instructions[idx].cmd = "jmp"
			}
		}
	}
	return res
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

	// part I:
	fmt.Println(loops(instructions))
	// part II:
	fmt.Println(swap(instructions))
}
