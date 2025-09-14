package btree

//import (
//	"bytes"
//	"regexp"
//)

type TokenType string

type Token struct {
	tokenType TokenType
	value     string
}

//func matchRegex(input, matchString string) bool {
//	if regexp.Find([]byte(input), []byte(matchString)) != nil {
//		return false
//	}
//	return true
//}
//
//func (t Token) readChar(index int, buffer *bytes.Buffer) rune {
//	buffer.ReadRune(int)
//}
