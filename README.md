## YAPP!
Yet another pratt parser!

Built test-first, in Go. Supports binary operators with precedence
(`+` `-` `*` `/` `^`), left- and right-associativity, grouping with
parentheses, prefix `-`, and postfix `!`.

### Running tests

```
go test ./...
```

Add `-v` for per-test output.

### Usage

```go
expr := pratt.Parse("1 + 2 * 3")
fmt.Println(expr.String()) // (+ 1 (* 2 3))
```

### Changelog

- Real tokenizer — handles multi-digit numbers and operators with no
  surrounding whitespace required (`"1+2*3"` now parses the same as
  `"1 + 2 * 3"`)
- Postfix operator (`!`)
- Prefix operator (`-`), including recursing into grouped/nested
  operands
- Grouping with parentheses
- Right-associativity (`^`)
- Left-associativity
- Binary operators with precedence (`+` `-` `*` `/`)
- Literal values
