package pratt

import (
	"fmt"
	"math"
	"strconv"
)

type Expr interface {
	String() string
	Eval() (int, error)
}
type Literal struct {
	Value string
}

func (l Literal) String() string {
	return l.Value
}

func (l Literal) Eval() (int, error) {
	numInt, err := strconv.Atoi(l.Value)
	if err != nil {
		return 0, fmt.Errorf("invalid Literal  %q: %w", l.Value, err)
	}
	return numInt, nil
}

type Binary struct {
	Left  Expr
	Right Expr
	Op    string
}

func (b Binary) String() string {
	return "(" + b.Op + " " + b.Left.String() + " " + b.Right.String() + ")"
}

func (b Binary) Eval() (int, error) {
	// usual case(considering we get a regular node)
	// we will do postorder travsersal here. first we travser the left subtree
	Left, err := b.Left.Eval()
	if err != nil {
		return 0, fmt.Errorf("invalid Binary left %q: %w", b.Left.String(), err)
	}
	// then we do the right subtree
	Right, err := b.Right.Eval()
	if err != nil {
		return 0, fmt.Errorf("invalid Binary right %q: %w", b.Right.String(), err)
	}
	// base case where we actually do the operations based on operators
	switch b.Op {
	case "+":
		return Left + Right, nil
	case "-":
		return Left - Right, nil
	case "*":
		return Left * Right, nil
	case "/":
		if Right == 0 {
			return 0, fmt.Errorf("division by zero: %s", b.String())
		}
		return Left / Right, nil
	case "^":
		return int(math.Pow(float64(Left), float64(Right))), nil
	}
	return 0, fmt.Errorf("unknown operator %q", b.Op)
}

type Unary struct {
	Op    string
	Right Expr
}

func (u Unary) String() string {
	return "(" + u.Op + u.Right.String() + ")"
}
func (u Unary) Eval() (int, error) {
	// yeah, let's say if we want to recursively chain something like "------------5"
	Right, err := u.Right.Eval()
	if err != nil {
		return 0, fmt.Errorf("invalid Unary right %q: %w", u.Right.String(), err)
	}
	switch u.Op {
	case "+":
		return Right, nil
	case "-":
		return (0 - Right), nil
	}
	return 0, nil
}

type Postfix struct {
	Left Expr
	Op   string
}

func (p Postfix) String() string {
	return "(" + p.Op + " " + p.Left.String() + ")"
}

func (p Postfix) Eval() (int, error) {
	// yup, recursively parsing again!
	Left, err := p.Left.Eval()
	if err != nil {
		return 0, fmt.Errorf("invalid Unary right %q: %w", p.Left.String(), err)
	}
	switch p.Op {
	case "!":
		if Left < 0 {
			return 0, fmt.Errorf("factorial of negative number: %d", Left)
		}
		numLeft := factorial(Left)
		return numLeft, nil
	}
	return 0, nil
}

func factorial(n int) int {
	var result int = 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}

type Parser struct {
	tokens []string
	pos    int
}

func (p *Parser) peek() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	return p.tokens[p.pos]
}
func (p *Parser) advance() string {
	cur_tok := p.pos
	p.pos++
	return p.tokens[cur_tok]
}

func bindingPower(op string) (left, right int) {
	switch op {
	case "+", "-":
		return 1, 2
	case "*", "/":
		return 3, 4
	case "^":
		return 6, 5
	case "!":
		return 8, 0
	default:
		return 0, 0
	}
}
func (p *Parser) parseExpr(floor int) Expr {
	tok := p.advance()
	var left Expr
	if tok == "(" {
		left = p.parseExpr(0)
		p.advance() // skipping the succeeding ")"
	} else if tok == "-" {
		// for now we will consider only "-"
		// left = Unary{Op: tok, Right: Literal{Value: p.advance()}}
		// we choose 7 since we want it to be highrt than the leftBP of ops which we do not want to consume
		left = Unary{Op: tok, Right: p.parseExpr(7)}
	} else if tok == "+" {
		// experimental shi with unary +. just saw an unary + in ruslan's blog lol
		left = Unary{Op: tok, Right: p.parseExpr(7)}
	} else {
		left = Literal{Value: tok}
	}
	// nud := Literal{Value: p.advance()}
	// var left Expr = nud
	for {
		l, r := bindingPower(p.peek())
		if floor >= l {
			break
		}
		op := p.advance()
		if op == "!" {
			left = Postfix{Left: left, Op: op}
		} else {
			right := p.parseExpr(r)
			left = Binary{Left: left, Right: right, Op: op}
		}
	}
	return left
}
func tokenize(data string) []string {
	var tokens []string
	i := 0
	for i < len(data) {
		switch {
		case data[i] == ' ':
			i++
		case data[i] >= '0' && data[i] <= '9':
			// we will consume consecutive digits into a token
			tok := ""
			for ; i < len(data) && data[i] >= '0' && data[i] <= '9'; i++ {
				tok += string(data[i])
			}
			tokens = append(tokens, tok)
		default:
			// if  there are single character tokens like '+','-', '*', '/'
			tokens = append(tokens, string(data[i]))
			i++
		}
	}
	return tokens
}
func Parse(data string) Expr {
	// list := strings.Split(data, " ")
	// if len(list) == 1 {
	// 	return Literal{Value: list[0]}
	// }
	// return Binary{Left: Literal{Value: list[0]}, Right: Literal{Value: list[2]}, Op: list[1]}
	// naive way by getting tokens by splitting on spaces
	// tokens := strings.Split(data, " ")
	// p := &Parser{tokens: tokens, pos: 0}
	// return p.parseExpr(0)

	// we will now introduce the tokenizer which does not necessarily split on spaces
	tokens := tokenize(data)
	p := &Parser{tokens: tokens, pos: 0}
	return p.parseExpr(0)
}
