package day16

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

var conversionTable = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func parseInputs(inputFilePath string) (string, error) {
	fileLine := <-helper.FileLineReader(inputFilePath)
	result := ""
	for _, c := range fileLine {
		result += conversionTable[c]
	}
	return result, nil
}

type packet string

func handleRawPacket(p packet, versionCounter map[int]int) (int, packet, error) {
	rawVersion := p[:3]
	version, err := strconv.ParseInt(string(rawVersion), 2, 64)
	if err != nil {
		return 0, "", errors.Wrap(err, "")
	}
	if currentCount, exists := versionCounter[int(version)]; exists {
		versionCounter[int(version)] = currentCount + 1
	} else {
		versionCounter[int(version)] = 1
	}
	rawTypeID := p[3:6]
	typeID, err := strconv.ParseInt(string(rawTypeID), 2, 64)
	if err != nil {
		return 0, "", errors.Wrap(err, "")
	}
	result, remainingPacket, err := handlePacket(int(version), int(typeID), p[6:], versionCounter)
	return result, remainingPacket, errors.Wrap(err, "")
}

func handleMultiArg(remainingPacket packet, versionCounter map[int]int) ([]int, packet, error) {
	var result []int
	switch remainingPacket[0] {
	case '0':
		rawLength := remainingPacket[1:16]
		argLength, err := strconv.ParseInt(string(rawLength), 2, 64)
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		remainingPacket = remainingPacket[16:]
		for argLength > 0 {
			partial, newRemainingPacket, err := handleRawPacket(remainingPacket, versionCounter)
			if err != nil {
				return nil, "", errors.Wrap(err, "")
			}
			argLength -= int64(len(remainingPacket) - len(newRemainingPacket))
			remainingPacket = newRemainingPacket
			result = append(result, partial)
		}
		return result, remainingPacket, nil
	case '1':
		rawArgCount := remainingPacket[1:12]
		argCount, err := strconv.ParseInt(string(rawArgCount), 2, 64)
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		remainingPacket = remainingPacket[12:]
		for i := 0; i < int(argCount); i++ {
			var partial int
			partial, remainingPacket, err = handleRawPacket(remainingPacket, versionCounter)
			if err != nil {
				return nil, "", errors.Wrap(err, "")
			}
			result = append(result, partial)
		}
		return result, remainingPacket, nil
	default:
		return nil, "", errors.New(fmt.Sprintf("Unsupported length type ID %v", remainingPacket[0]))
	}
}

func handlePacket(version, typeID int, remainingPacket packet, versionCounter map[int]int) (int, packet, error) {
	switch typeID {
	case 0:
		// Sum
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		sum := 0
		for _, n := range args {
			sum += n
		}
		return sum, newRemainingPacket, nil
	case 1:
		// Product
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		product := 1
		for _, n := range args {
			product *= n
		}
		return product, newRemainingPacket, nil
	case 2:
		// Min
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		min := math.MaxInt
		for _, n := range args {
			if n < min {
				min = n
			}
		}
		return min, newRemainingPacket, nil
	case 3:
		// Max
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		max := math.MinInt
		for _, n := range args {
			if n > max {
				max = n
			}
		}
		return max, newRemainingPacket, nil
	case 4:
		// Literal
		var rawLiteral packet = ""
		for groupBit := remainingPacket[0]; groupBit != '0'; groupBit = remainingPacket[0] {
			rawLiteral += remainingPacket[1:5]
			remainingPacket = remainingPacket[5:]
		}
		rawLiteral += remainingPacket[1:5]
		literal, err := strconv.ParseInt(string(rawLiteral), 2, 64)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		return int(literal), remainingPacket[5:], nil
	case 5:
		// Greater than
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		if len(args) != 2 {
			return 0, "", errors.New(fmt.Sprintf("Expected gt args to have length 2, but got %v (%d)", args, len(args)))
		}
		result := 0
		if args[0] > args[1] {
			result = 1
		}
		return result, newRemainingPacket, nil
	case 6:
		// Less than
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		if len(args) != 2 {
			return 0, "", errors.New(fmt.Sprintf("Expected lt args to have length 2, but got %v (%d)", args, len(args)))
		}
		result := 0
		if args[0] < args[1] {
			result = 1
		}
		return result, newRemainingPacket, nil
	case 7:
		// Equal
		args, newRemainingPacket, err := handleMultiArg(remainingPacket, versionCounter)
		if err != nil {
			return 0, "", errors.Wrap(err, "")
		}
		if len(args) != 2 {
			return 0, "", errors.New(fmt.Sprintf("Expected eq args to have length 2, but got %v (%d)", args, len(args)))
		}
		result := 0
		if args[0] == args[1] {
			result = 1
		}
		return result, newRemainingPacket, nil
	}
	return 0, "", errors.New(fmt.Sprintf("Unsupported type ID %d", typeID))
}

func Solve1(inputFilePath string) (int, error) {
	inputs, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	versionCounter := map[int]int{}
	_, _, err = handleRawPacket(packet(inputs), versionCounter)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	sum := 0
	for version, count := range versionCounter {
		sum += (version * count)
	}
	return sum, nil
}

func Solve2(inputFilePath string) (int, error) {
	inputs, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	versionCounter := map[int]int{}
	result, _, err := handleRawPacket(packet(inputs), versionCounter)
	return result, errors.Wrap(err, "")
}
