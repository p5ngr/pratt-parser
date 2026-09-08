package pratt

import "strings"

type Expr interface {
	String() string
}
type Literal struct {
	Value string
}

func (l Literal) String() string {
	return l.Value
}

type Binary struct {
	Left  Expr
	Right Expr
	Op    string
}

func (b Binary) String() string {
	return "(" + b.Op + " " + b.Left.String() + " " + b.Right.String() + ")"
}

type Unary struct {
	Op    string
	Right Expr
}

func (u Unary) String() string {
	return "(" + u.Op + u.Right.String() + ")"
}

type Postfix struct {
	Left Expr
	Op   string
}

func (p Postfix) String() string {
	return "(" + p.Op + " " + p.Left.String() + ")"
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
func Parse(data string) Expr {
	// list := strings.Split(data, " ")
	// if len(list) == 1 {
	// 	return Literal{Value: list[0]}
	// }
	// return Binary{Left: Literal{Value: list[0]}, Right: Literal{Value: list[2]}, Op: list[1]}
	tokens := strings.Split(data, " ")
	p := &Parser{tokens: tokens, pos: 0}
	return p.parseExpr(0)
}
