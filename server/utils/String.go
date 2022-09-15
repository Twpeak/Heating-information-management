package utils

import "strings"

//防止sql注入

func SqlInj(str string) bool {
	injStr := "'|and|exec|insert|select|delete|update|count|*|%|chr|mid|master|truncate|char|declare|; |or|-|+|, "
	injStra := strings.Split(injStr, "|")
	for  i := 0; i < len(injStra); i++ {
		if strings.Index(str, injStra[i]) >= 0 {
			return true
		}
	}
	return false
}
