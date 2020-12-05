package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
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

func validate(name, value string) bool {
	if name == "byr" {
		if len(value) != 4 {
			return false
		}

		s, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return s >= 1920 && s <= 2002
	} else if name == "iyr" {
		if len(value) != 4 {
			return false
		}

		s, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return s >= 2010 && s <= 2020
	} else if name == "eyr" {
		if len(value) != 4 {
			return false
		}

		s, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return s >= 2020 && s <= 2030
	} else if name == "hgt" {
		var n int
		var m string
		fmt.Sscanf(value, "%d%s", &n, &m)
		if m == "cm" {
			return n >= 150 && n <= 193
		} else if m == "in" {
			return n >= 59 && n <= 76
		} else {
			return false
		}
	} else if name == "ecl" {
		return value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth"
	} else if name == "pid" {
		if len(value) != 9 {
			return false
		}
	} else if name == "cid" {
		return true
	} else if name == "hcl" {
		var m string
		fmt.Sscanf(value, "#%s", &m)
		if len(m) != 6 {
			return false
		}
		src := []byte(m)
		dst := make([]byte, hex.DecodedLen(len(src)))
		_, err := hex.Decode(dst, src)
		return err == nil
	}
	return true
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
		data := strings.Split(field, ":")
		if _, ok := validFields[data[0]]; ok && validate(data[0], data[1]) {
			delete(validFields, data[0])
		}
	}
	return len(validFields) == 0
}
