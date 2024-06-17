package funky

import (
	"fmt"
	"strings"
)

func ToString[T any](arr []T) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, fmt.Sprint(v))
	}
	return strings.Join(strArr, ",")
}
