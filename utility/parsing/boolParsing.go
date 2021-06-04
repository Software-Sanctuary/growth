package parsing

import "strconv"

func SafeParseBool(strToParse string, defVal bool) bool {
	if val, err := strconv.ParseBool(strToParse); err != nil {
		return val
	}
	return defVal
}
