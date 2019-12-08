package unpackstr

import (
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

// Unpack string
// a4bc2d5e => aaaabccddddde
func Unpack(line string) (string, error) {
	if len(line) <= 1 {
		return "", errors.New("String length must be greater than 1")
	}

	matched, err := regexp.MatchString(`^\d.|\d{2,}`, line)
	if err != nil {
		return "", err
	}

	if matched {
		return "", errors.New("Invalid string")
	}

	reg, _ := regexp.Compile(`[0-9]`)
	result := make([]string, 0)

	for _, r := range line {
		str := string(r)

		if matched := reg.MatchString(str); matched {
			digit, _ := strconv.Atoi(str)
			unpackedStr := strings.Repeat(result[len(result)-1], digit - 1)
			result = append(result, unpackedStr)
		} else {
			result = append(result, str)
		}
	}

	return strings.Join(result, ""), nil
}
