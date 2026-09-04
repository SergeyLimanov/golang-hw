package hw02_unpack_string

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var b strings.Builder
	runes := []rune(s)

	var prev rune
	hasPrev := false

	for i := 0; i < len(runes); {
		r := runes[i]

		if r == '\' {
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}
			next := runes[i+1]
			if !(unicode.IsDigit(next) || next == '\') {
				return "", ErrInvalidString
			}

			if hasPrev {
				b.WriteRune(prev)
			}

			prev = next
			hasPrev = true

			i += 2
			continue
		}

		if unicode.IsDigit(r) {
			if !hasPrev {
				return "", ErrInvalidString
			}

			n, _ := strconv.Atoi(string(r))

			if n > 0 {
				b.WriteString(strings.Repeat(string(prev), n))
			}

			hasPrev = false
			i++
			continue
		}

		if hasPrev {
			b.WriteRune(prev)
		}
		prev = r
		hasPrev = true
		i++
	}

	if hasPrev {
		b.WriteRune(prev)
	}

	return b.String(), nil
}