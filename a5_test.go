package a5

import (
  "testing"
  "fmt"
)

func Test_solver_intefarce_candidate1(t *testing.T) {
  // does not work
  fmt.Sprintf("%q", solve())
  code := `
    P:=  x^2 + y^2 < 4
    Q:=  x < 1
    ∃x∃y(P ∧ Q)
  `
  if eval(code) != true {
    t.Errorf("evaluation is wrong")
  }
}

func TestSolver(t *testing.T) {
  var x Var
  var y Var
  // ∃x∃y(x + y < 4)
  P := exists([]Var{x, y}, less(plus(x, y), value(4)))
  t.Logf("%q", P)
}