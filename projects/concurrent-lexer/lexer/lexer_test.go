package lexer

import "testing"

func Test(t *testing.T) {
	a := A{}
	s := "1+2"
	a.run(s) // == 3
	a.run("a=3+(1-2)")
}
