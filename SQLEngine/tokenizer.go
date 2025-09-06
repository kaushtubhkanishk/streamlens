package SQLEngine

import (
	"strings"

	log "github.com/rs/zerolog"
)

type Token struct {
	value string
}

func Tokenize(text string, log log.Logger) []Token {
	var tokens []Token
	texts := strings.Split(text, " ")
	for _, text := range texts {
		tokens = append(tokens, Token{text})
	}
	log.Debug().Msgf("The tokens present in the string are:\n ")
	for _, token := range tokens {
		log.Debug().Msgf("%s ", token.value)
	}
	return tokens
}
