package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	by, _ := ioutil.ReadFile("2020/day4/input.txt")
	tmp := strings.Split(string(by), "\n\n")
	validPassports := 0
	for i := range tmp {
		passport := make(map[string]string)
		fields := strings.Split(strings.Join(strings.Split(tmp[i], "\n"), " "), " ")
		for _, field := range fields {
			field := strings.Split(field, ":")
			passport[field[0]] = field[1]
		}
		mandatoryFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		isValid := true
		for _, m := range mandatoryFields {
			if _, ok := passport[m]; !ok {
				isValid = false
				break
			}
		}
		if isValid {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func part2() {
	by, _ := ioutil.ReadFile("2020/day4/input.txt")
	tmp := strings.Split(string(by), "\n\n")
	validPassports := 0
	for i := range tmp {
		passport := make(map[string]string)
		fields := strings.Split(strings.Join(strings.Split(tmp[i], "\n"), " "), " ")
		for _, field := range fields {
			field := strings.Split(field, ":")
			passport[field[0]] = field[1]
		}
		mandatoryFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		isValid := true
		for _, m := range mandatoryFields {
			if v, ok := passport[m]; ok {
				if m == "byr" {
					byr, _ := strconv.Atoi(v)
					if byr < 1920 || byr > 2002 {
						isValid = false
						break
					}
				} else if m == "iyr" {
					iyr, _ := strconv.Atoi(v)
					if iyr < 2010 || iyr > 2020 {
						isValid = false
						break
					}
				} else if m == "eyr" {
					eyr, _ := strconv.Atoi(v)
					if eyr < 2020 || eyr > 2030 {
						isValid = false
						break
					}
				} else if m == "hgt" {
					rgx := regexp.MustCompile(`(\d+)(in|cm)`)
					m := rgx.FindStringSubmatch(v)
					if !rgx.MatchString(v) {
						isValid = false
						break
					}
					h, _ := strconv.Atoi(m[1])
					if m[2] == "cm" && (h < 150 || h > 193) {
						isValid = false
						break
					} else if m[2] == "in" && (h < 59 || h > 76) {
						isValid = false
						break
					}
				} else if m == "hcl" {
					if ok, _ := regexp.MatchString(`#[0-9a-f]{6}`, v); !ok {
						isValid = false
						break
					}
				} else if m == "ecl" {
					if v != "amb" && v != "blu" && v != "brn" && v != "gry" && v != "grn" && v != "hzl" && v != "oth" {
						isValid = false
						break
					}
				} else if m == "pid" {
					if ok, _ := regexp.MatchString(`^\d{9}$`, v); !ok {
						isValid = false
						break
					}
				}
			} else {
				isValid = false
				break
			}
		}
		if isValid {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func main() {
	part1()
	part2()
}
