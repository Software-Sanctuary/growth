package parsing

import "strconv"

func SafeParseInt32(strToParse string, defVal int) int {
	if val, err := strconv.ParseInt(strToParse, 10, 64); err != nil {
		return int(val)
	}
	return defVal
}

func SafeParseInt32DefaultZero(strToParse string) int {
	if val, err := strconv.ParseInt(strToParse, 10, 64); err != nil {
		return int(val)
	}
	return 0
}

func SafeParseInt64(strToParse string, defVal int64) int64 {
	if val, err := strconv.ParseInt(strToParse, 10, 64); err != nil {
		return val
	}
	return defVal
}

func SafeParseInt64DefaultZero(strToParse string) int64 {
	if val, err := strconv.ParseInt(strToParse, 10, 64); err != nil {
		return val
	}
	return 0
}
