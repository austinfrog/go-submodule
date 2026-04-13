# go-submodule

A repository containing a Go submodule.

## Structure

```
foo/         // Directory containing a simple Go module
  - foo.go
  - go.mod
  - go.sum
README.md    // This file
```

## Usage

```go
import "github.com/austinfrog/go-submodule/foo"

func main() {
    msg := foo.Hello()
    fmt.Println(msg)
}
```
