package btree

import (
	"fmt"
	"strings"
)

func SplitLine(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("the passed string is empty")
	}

	var token []string
	token = strings.Split(s, " ")

	fmt.Println("The tokens present in the string are:\n ")
	fmt.Println(token)

	return nil
}
