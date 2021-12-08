package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func Solve(prefix string, target string) int {
	i := 0
	for {
		hash := md5.Sum([]byte(prefix + fmt.Sprint(i)))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), target) {
			return i
		}
		i++
	}
}
