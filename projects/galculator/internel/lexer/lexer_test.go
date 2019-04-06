package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	tokens, _ := Lex("1+1")
	require.Equal(t,
		[]Token{
			Number{Value: "1"},
			Operator{Value: "+"},
			Number{Value: "1"},
		},
		tokens)

	tokens, _ = Lex("1 + 1")
	require.Equal(t,
		[]Token{
			Number{Value: "1"},
			Operator{Value: "+"},
			Number{Value: "1"},
		},
		tokens)

	tokens, _ = Lex("1 + 1 0000 0000") // support white space between numbers for better big number literal displaying
	require.Equal(t,
		[]Token{
			Number{Value: "1"},
			Operator{Value: "+"},
			Number{Value: "100000000"},
		},
		tokens)

	tokens, _ = Lex("1 + 1 - 20") // support multiple digit numbers
	require.Equal(t,
		[]Token{
			Number{Value: "1"},
			Operator{Value: "+"},
			Number{Value: "1"},
			Operator{Value: "-"},
			Number{Value: "20"},
		},
		tokens)
}

func TestLexerError(t *testing.T) {
	tokens, err := Lex("a")
	require.Empty(t, tokens)
	require.EqualError(t, err, "character 'a' is not expected")

	_, err = Lex("1 +a")
	require.EqualError(t, err, "character 'a' is not expected after an operator")
	_, err = Lex("1 + function")
	require.EqualError(t, err, "character 'f' is not expected after an operator")
	_, err = Lex("1 + +")
	require.EqualError(t, err, "character '+' is not expected after an operator")

	_, err = Lex("1 function")
	require.EqualError(t, err, "character 'f' is not expected after a number")
}
