# Weak
The `weak` package offers a strategy for reducing GC pressure in applications dealing with many short-lived objects by introducing the notion of so-called "weak" pointers and slices.

Weak pointers and slices allow you to hold a non-owning reference to an object or slice, which does not affect the object's or slice's lifetime. The lifetime will not be extended on the basis of the weak pointer or slice, meaning if all "strong" references are deleted, the object or slice will be deleted even if there are weak references remaining.

While this does place added responsibility on your shoulders to ensure your objects are not prematurely swept up by the garbage collector, this is balanced by a substantial GC pressure reduction on tasks like lexical analysis.

## Usage
### Slice 
```go
type Token struct {
    Type  int
    Value weak.Slice[rune]
}


func TokenizeTo(dst []Token, in []rune) {
    off := 0
    for {
        tok, delta := nextToken(in[off:])
        if n == 0 {
            return dst
        }
        off += delta

        // Token contents can be accessed via tok.Indirect().
        dst = append(dst, tok)
    }
}

// token is a method of convenience utilised by the nextToken.
func token(t TokenType, in []rune, delta int) (Token, int) {
    return Token{
        Type:  t,
        Value: weak.NewSlice(in),
    }, delta
}
```

### Pointer
```go
obj := Object{...}
ptr := weak.NewPointer(&obj)
```
