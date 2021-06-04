package parsing

import "strconv"

func SafeParseFloat(strToParse string, defVal float32) float32 {
	if val, err := strconv.ParseFloat(strToParse, 32); err != nil {
		return float32(val)
	}
	return defVal
}

func SafeParseFloatDefaultZero(strToParse string) float32 {
	if val, err := strconv.ParseFloat(strToParse, 32); err != nil {
		return float32(val)
	}
	return 0.0
}
