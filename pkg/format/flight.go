package format

import "regexp"

func FlightCodeToCode(code string) string {
	re, err := regexp.Compile(`[^a-zA-Z]`)
	if err != nil {
		return ""
	}
	return re.ReplaceAllString(code, "")
}
