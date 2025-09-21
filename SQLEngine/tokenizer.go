package SQLEngine

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kaushtubhkanishk/streamlens/internal/lexer"
	"github.com/kaushtubhkanishk/streamlens/internal/token"
)

const PROMPT = ">>"

func Tokenize(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.TokenType != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
