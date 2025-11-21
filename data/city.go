package data

var mapCities = map[string]string{
	"DPS": "Denpasar",
	"CGK": "Cengkareng",
}

func GetCityName(code string) string {
	if name, ok := mapCities[code]; ok {
		return name
	}
	return ""
}
