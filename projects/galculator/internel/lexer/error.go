package lexer

import (
	"fmt"
)

// Error is the type for all lexing errors.
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Lexing Error: %s", e.Message)
}
