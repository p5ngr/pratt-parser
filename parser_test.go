package pratt

import "testing"

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
