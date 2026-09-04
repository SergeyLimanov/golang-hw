
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

	var runes []rune = []rune(s)


	var prev rune
	var hasPrev bool = false


	var i int = 0
	for i < len(runes) {
		var r rune = runes[i]


		if r == '\\' {

			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}
			var next rune = runes[i+1]

			if !(unicode.IsDigit(next) || next == '\\') {
				return "", ErrInvalidString
			}


			if hasPrev {
				b.WriteRune(prev)
				hasPrev = false
			}

			prev = next
			hasPrev = true


			i += 2
			continue
		}

		//
		if unicode.IsDigit(r) {

			if !hasPrev {
				return "", ErrInvalidString
			}


			var n int
			n, _ = strconv.Atoi(string(r))


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
		// Текущий символ становится новым "ожидающим".
		prev = r
		hasPrev = true
		i++
	}


	if hasPrev {
		b.WriteRune(prev)
	}


	return b.String(), nil
}