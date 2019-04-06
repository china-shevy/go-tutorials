package lexer

import (
	"text/scanner"
)

type ToekenReceiver struct {
	tokens []Token
}

func Lex(input string) ([]Token, error) {
	t := ToekenReceiver{}
	r := NewRuneEmitter(input)
	for next, err := StateBegin(r, &t); next != nil || err != nil; {
		if err != nil {
			return t.tokens, err
		}
		next, err = next(r, &t)
	}
	return t.tokens, nil
}

func (l *ToekenReceiver) Send(t Token) {
	l.tokens = append(l.tokens, t)
}

type RuneEmitter struct {
	ch chan rune
}

func NewRuneEmitter(input string) *RuneEmitter {
	r := RuneEmitter{ch: make(chan rune)}
	go func() {
		for _, c := range input {
			r.ch <- c
		}
		r.ch <- scanner.EOF
	}()
	return &r
}

func (r *RuneEmitter) Next() rune {
	next := <-r.ch
	return next
}

type Token interface {
	Literal() string
	Type() string
}

type Number struct {
	Value string
}

func (n Number) Literal() string {
	return n.Value
}

func (n Number) Type() string {
	return "Number"
}

type Operator struct {
	Value string
}

func (o Operator) Literal() string {
	return o.Value
}

func (o Operator) Type() string {
	return "Operator"
}

var (
	Add = Operator{
		Value: "+",
	}
	Sub = Operator{
		Value: "-",
	}
	Mul = Operator{
		Value: "*",
	}
	Div = Operator{
		Value: "/",
	}
)
