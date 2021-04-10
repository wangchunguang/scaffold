package func_string

import "strings"

func StrReplace(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
