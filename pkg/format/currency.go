package format

import (
	"fmt"
	"ota/constant"
)

func CurrencyComma(n int64) string {
	s := fmt.Sprintf("%d", n)
	negative := false

	if n < 0 {
		negative = true
		s = s[1:]
	}

	var out []rune
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		out = append([]rune{rune(s[i])}, out...)
		count++

		if count == 3 && i != 0 {
			out = append([]rune{constant.DefaultSeparator}, out...)
			count = 0
		}
	}

	final := string(out)
	if negative {
		final = "-" + final
	}

	return final
}
