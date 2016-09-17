# conv

Convert any values to primiteves.

## Import

```go
import "github.com/najeira/conv"
```

## Usage

```go
var n int = 123
s := conv.String(n) // "123"
```

```go
var s string = "123"
n := conv.Int(s) // 123
```

Funcs returns second argument when convering is failed.

```go
var s string = "not integer"
n := conv.Int(s, 456) // 456
```
