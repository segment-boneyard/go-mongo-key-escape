package mongo

import "regexp"

// Escapes the mongo `key`.

func Escape(key string) string {
	re := regexp.MustCompile(`\$`)
	key = re.ReplaceAllString(key, "\uFF04")
	re = regexp.MustCompile(`\.`)
	key = re.ReplaceAllString(key, "\uFF0E")
	return key
}

// Unescapes the mongo `key`.

func Unescape(key string) string {
	re := regexp.MustCompile("\uFF04")
	key = re.ReplaceAllString(key, "$")
	re = regexp.MustCompile("\uFF0E")
	key = re.ReplaceAllString(key, ".")
	return key
}
