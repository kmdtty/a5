package a5

import (
  "testing"
  "fmt"
)

func TestSolve(t *testing.T) {
  fmt.Sprintf("%q", solve())
  code := `
    P:  x^2 + y^2 < 4
    Q:  x < 1
    [[P and Q]]
  `
  eval(code)
}
