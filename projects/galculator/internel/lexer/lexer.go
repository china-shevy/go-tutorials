package lexer

import (
	"text/scanner"
)

type ToekenReceiver struct {
	tokens []Token
}

func Lex(input string) ([]Token, *Error) {
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
