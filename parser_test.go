package pratt

import (
	"testing"
)

func TestParseLiteral(t *testing.T) {
	got := Parse("42")

	want := "42"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "42", got.String(), want)
	}
}

func TestParseBinary(t *testing.T) {
	got := Parse("1 + 2")
	want := "(+ 1 2)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "1 + 2", got.String(), want)
	}
}

func TestParseLeftAssoc(t *testing.T) {
	got := Parse("1 + 2 + 3")
	want := "(+ (+ 1 2) 3)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "1 + 2 + 3", got.String(), want)
	}
}
func TestParsePrecedence(t *testing.T) {
	got := Parse("1 + 2 * 3")
	want := "(+ 1 (* 2 3))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "1 + 2 * 3", got.String(), want)
	}
}
func TestParseGrouping(t *testing.T) {
	got := Parse("( 1 + 2 ) * 3")
	want := "(* (+ 1 2) 3)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "( 1 + 2 ) * 3", got.String(), want)
	}
}
func TestParseRightAssoc(t *testing.T) {
	got := Parse("2 ^ 3 ^ 2")
	want := "(^ 2 (^ 3 2))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "2 ^ 3 ^ 2", got.String(), want)
	}
}
func TestParseUnary(t *testing.T) {
	got := Parse("- 5 + 3")
	want := "(+ (-5) 3)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "- 5 + 3", got.String(), want)
	}
}

func TestParseUnary2(t *testing.T) {
	got := Parse("3 + - 5 * 12")
	want := "(+ 3 (* (-5) 12))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "3 + - 5 * 12", got.String(), want)
	}
}

func TestParseUnaryGrouping(t *testing.T) {
	got := Parse("- ( 1 + 2 ) + 3")
	want := "(+ (-(+ 1 2)) 3)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "- ( 1 + 2 ) + 3", got.String(), want)
	}
}

func TestParseUnaryGrouping2(t *testing.T) {
	got := Parse("- - 5")
	want := "(-(-5))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "- - 5", got.String(), want)
	}
}

func TestParsePostfix(t *testing.T) {
	got := Parse("5 !")
	want := "(! 5)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "5 !", got.String(), want)
	}
}

func TestParseEverything(t *testing.T) {
	got := Parse("( 1 + 2 ) * 3 ^ 2 ! + - 4")
	want := "(+ (* (+ 1 2) (^ 3 (! 2))) (-4))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "( 1 + 2 ) * 3 ^ 2 ! + - 4", got.String(), want)
	}
}

// tokenizer test 1
func TestParseNoSpaces(t *testing.T) {
	got := Parse("1+2")
	want := "(+ 1 2)"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "1+2", got.String(), want)
	}
}

func TestParseNoSpacesMultiDigit(t *testing.T) {
	got := Parse("12+34*2")
	want := "(+ 12 (* 34 2))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "12+34*2", got.String(), want)
	}
}

func TestParseEverythingNoSpaces(t *testing.T) {
	got := Parse("(1+2)*3^2!+-4")
	want := "(+ (* (+ 1 2) (^ 3 (! 2))) (-4))"
	if got.String() != want {
		t.Errorf("Parse(%q) = %q, want %q", "(1+2)*3^2!+-4", got.String(), want)
	}
}

func TestEvalLiteral(t *testing.T) {
	got, err := Parse("42").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "42", err)
	}
	want := 42
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "42", got, want)
	}
}
func TestEvalBinary(t *testing.T) {
	got, err := Parse("1 + 2").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "1 + 2", err)
	}
	want := 3
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "1 + 2", got, want)
	}
}

func TestEvalPower(t *testing.T) {
	got, err := Parse("2 ^ 3").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "2 ^ 3", err)
	}
	want := 8
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "2 ^ 3", got, want)
	}
}

func TestEvalDivision(t *testing.T) {
	got, err := Parse("6 / 2").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "6 / 2", err)
	}
	want := 3
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "6 / 2", got, want)
	}
}

func TestEvalDivisionByZero(t *testing.T) {
	_, err := Parse("1 / 0").Eval()
	if err == nil {
		t.Fatalf("Parse(%q).Eval() = nil error, want a division-by-zero error", "1 / 0")
	}
}

func TestEvalUnary(t *testing.T) {
	got, err := Parse("- 5").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "- 5", err)
	}
	want := -5
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "- 5", got, want)
	}
}
func TestEvalUnaryPlus(t *testing.T) {
	got, err := Parse("+ 5").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "+ 5", err)
	}
	want := 5
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "+ 5", got, want)
	}
}
func TestEvalPostfix(t *testing.T) {
	got, err := Parse("5 !").Eval()
	if err != nil {
		t.Fatalf("Parse(%q).Eval() returned error: %v", "5 !", err)
	}
	want := 120
	if got != want {
		t.Errorf("Parse(%q).Eval() = %d, want %d", "5 !", got, want)
	}
}

func TestEvalEverything(t *testing.T) {
	got, err := Parse("( 1 + 2 ) * 3 ^ 2 ! + - 4").Eval()
	if err != nil {
		t.Fatalf("Eval() returned error: %v", err)
	}
	// (1+2)=3, 2!=2, 3^2=9, 3*9=27, -4, 27 + (-4) = 23
	want := 23
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
