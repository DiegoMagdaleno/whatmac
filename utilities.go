package whatmac

import "strconv"

func toFloat(value string) float64 {
	valueFloat, _ := strconv.ParseFloat(value, 64)
	return valueFloat
}
