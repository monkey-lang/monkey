package repl

import (
	"bufio"
	"fmt"
	"github.com/monkey-lang/monkey/scanner"
	"github.com/monkey-lang/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(stdin io.Reader, stdout io.Writer) {
	input := bufio.NewScanner(stdin)
	for {
		fmt.Printf(PROMPT)
		scanned := input.Scan()
		if !scanned {
			//something went wrong
			return
		}
		line := input.Text()
		l := scanner.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
