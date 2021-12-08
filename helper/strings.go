package helper

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func SplitAndCheck(s string, arg string, expectedSize int) ([]string, error) {
	parts := strings.Split(s, arg)
	if len(parts) != expectedSize {
		return nil, errors.New(fmt.Sprintf("Expected %s to have %d parts but it has %d", s, expectedSize, len(parts)))
	}
	return parts, nil
}
