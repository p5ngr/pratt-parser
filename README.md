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

### Known limitation

The tokenizer splits on whitespace only, so every token needs a
surrounding space (`"( 1 + 2 ) * 3"`, not `"(1+2)*3"`). A real
tokenizer (multi-digit numbers, no-space input) is next.
