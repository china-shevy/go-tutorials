package lexer

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

type LeftParentheses struct{}

func (p LeftParentheses) Literal() string {
	return "("
}

func (p LeftParentheses) Type() string {
	return "Left Parentheses"
}

type RightParentheses struct{}

func (p RightParentheses) Literal() string {
	return ")"
}

func (p RightParentheses) Type() string {
	return "Right Parentheses"
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
