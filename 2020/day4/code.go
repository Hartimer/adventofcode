package day4

import (
	"adventofcode/helper"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var eyeCollors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var hairColorRegex = regexp.MustCompile("#[a-z0-9]{6}")

func validEyeColor(color string) bool {
	for _, eyecolor := range eyeCollors {
		if eyecolor == color {
			return true
		}
	}
	return false
}

type passport struct {
	byr string // Birth Year
	iyr string // Issue Year
	eyr string // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

func (p *passport) set(key, value string) {
	switch key {
	case "byr":
		p.byr = value
	case "iyr":
		p.iyr = value
	case "eyr":
		p.eyr = value
	case "hgt":
		p.hgt = value
	case "hcl":
		p.hcl = value
	case "ecl":
		p.ecl = value
	case "pid":
		p.pid = value
	case "cid":
		p.cid = value
	default:
		panic(fmt.Sprintf("Unknown key %s", key))
	}
}

func (p *passport) safeSet(key, value string) {
	switch key {
	case "byr":
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil || intValue < 1920 || intValue > 2002 {
			return
		}
		p.byr = value
	case "iyr":
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil || intValue < 2010 || intValue > 2020 {
			return
		}
		p.iyr = value
	case "eyr":
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil || intValue < 2020 || intValue > 2030 {
			return
		}
		p.eyr = value
	case "hgt":
		newValue := strings.TrimSuffix(value, "in")
		if len(newValue) == len(value) {
			newValue = strings.TrimSuffix(value, "cm")
		} else if intValue, err := strconv.ParseInt(value, 10, 32); err != nil || intValue < 59 || intValue > 76 {
			return
		}
		if len(newValue) == len(value) {
			return
		} else if intValue, err := strconv.ParseInt(value, 10, 32); err != nil || intValue < 150 || intValue > 193 {
			return
		}
		p.hgt = value
	case "hcl":
		if hairColorRegex.MatchString(value) {
			p.hcl = value
		}
	case "ecl":
		if validEyeColor(value) {
			p.ecl = value
		}
	case "pid":
		if _, err := strconv.ParseInt(value, 10, 32); err == nil && len(value) == 9 {
			p.pid = value
		}
	case "cid":
		p.cid = value
	default:
		panic(fmt.Sprintf("Unknown key %s", key))
	}
}

func (p *passport) isValid() bool {
	return len(p.byr) > 0 &&
		len(p.iyr) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0
}

func Solve1(inputFilePath string) (int, error) {
	var workingPassword *passport
	validPassports := 0

	for fileLine := range helper.FileLineReader(inputFilePath) {
		if workingPassword == nil {
			workingPassword = &passport{}
		}

		inputLine := strings.TrimSpace(fileLine)
		log.Printf("Processing %q", inputLine)
		if len(inputLine) == 0 {
			if workingPassword.isValid() {
				validPassports++
			}
			workingPassword = nil
		} else {
			for _, keyValuePair := range strings.Split(inputLine, " ") {
				keyValueParts := strings.Split(keyValuePair, ":")
				if len(keyValueParts) != 2 {
					return 0, errors.New(fmt.Sprintf("Expected key value pair %q to have 2 parts, but has %d", keyValuePair, len(keyValueParts)))
				}
				workingPassword.set(keyValueParts[0], keyValueParts[1])
			}
		}
	}

	if workingPassword != nil && workingPassword.isValid() {
		validPassports++
	}

	return validPassports, nil
}

func Solve2(inputFilePath string) (int, error) {
	var workingPassword *passport
	validPassports := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if workingPassword == nil {
			workingPassword = &passport{}
		}

		inputLine := strings.TrimSpace(fileLine)
		log.Printf("Processing %q", inputLine)
		if len(inputLine) == 0 {
			if workingPassword.isValid() {
				validPassports++
			}
			workingPassword = nil
		} else {
			for _, keyValuePair := range strings.Split(inputLine, " ") {
				keyValueParts := strings.Split(keyValuePair, ":")
				if len(keyValueParts) != 2 {
					return 0, errors.New(fmt.Sprintf("Expected key value pair %q to have 2 parts, but has %d", keyValuePair, len(keyValueParts)))
				}
				workingPassword.safeSet(keyValueParts[0], keyValueParts[1])
			}
		}
	}

	if workingPassword != nil && workingPassword.isValid() {
		validPassports++
	}

	return validPassports, nil
}
