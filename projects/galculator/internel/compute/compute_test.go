package compute

import (
	"galculator/internel/lexer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	require.Equal(t, "-3", Compute("1 + 1 - 3 * 9 - 1 / 3"))
	require.Equal(t, "-8", Compute("1 + 1 - 10"))
	require.Equal(t, "2", Compute("1+1"))
	require.Equal(t, "2", Compute("(1+1)"))
}

func TestParseParenthesisExpression(t *testing.T) {
	tokens1, _ := lexer.Lex("1+1")
	tokens2, _ := lexer.Lex("1+1)")
	require.Equal(t,
		tokens1,
		ParseParenthesisExpression(tokens2).Tokens,
	)

	tokens1, _ = lexer.Lex("1+(1+1)")
	tokens2, _ = lexer.Lex("1+(1+1))")
	require.Equal(t,
		tokens1,
		ParseParenthesisExpression(tokens2).Tokens,
	)
}
