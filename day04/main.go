package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var valids int
	var line string
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			if isValid(line) {
				valids++
			}
			line = ""
		} else {
			line += scanner.Text()
			line += " "
		}
	}
	if isValid(line) {
		valids++
	}
	fmt.Println(valids)

}

func isValid(data string) bool {
	validFields := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}
	for _, field := range strings.Fields(data) {
		fieldName := strings.Split(field, ":")[0]
		if _, ok := validFields[fieldName]; ok {
			delete(validFields, fieldName)
		}
	}
	return len(validFields) == 0
}
