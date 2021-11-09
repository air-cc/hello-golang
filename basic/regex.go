package main

import (
	"fmt"
	"regexp"
	"strings"
)

// TableFieldFilter .
type TableFieldFilter struct {
	FieldName  string
	FieldValue interface{}
	Operator   string
}

func main() {
	filterStr := "category=edu&style=daily&hold_date>=2021-09-01"

	pat := `([^\=\>\<\&]+)([\=\>\<]+)([^\=\>\<\&]+)`
	re, err := regexp.Compile(pat)
	if err != nil {
		return
	}

	filterItemsStr := strings.Split(filterStr, "&")

	filters := make([]*TableFieldFilter, len(filterItemsStr))
	for i, itemStr := range filterItemsStr {
		m := re.FindAllStringSubmatch(itemStr, -1)
		if m == nil {
			fmt.Println("no match for" + itemStr)
			return
		}

		fmt.Println(m[0][1], m[0][3], m[0][2])

		filters[i] = &TableFieldFilter{
			FieldName:  m[0][1],
			FieldValue: m[0][3],
			Operator:   m[0][2],
		}
	}
}
