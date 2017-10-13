package a5

import  "fmt"


// types
type Int struct {
  x int    
}

type Pred struct {
  // TODO: fix to have operators
  v bool
}

type Var struct {
  val int
}

type Op interface {
  eval() int
}

// unary operator
type OpUnary struct {
  first Var
  Op
}

type OpPlus struct {
  Op
  first Var
  second Var
}

func (v *OpUnary) eval() int {
  return v.first.val
}

func (p *OpPlus) eval() int {
  return p.first.val + p.second.val
}

// API

func less(x Op, y Op) Pred {
  return Pred{x.eval() < y.eval()}
}

func value(v int) Op {
  return &OpUnary{first: Var{v}}
}

func plus(x Var, y Var) Op {
  return &OpPlus{first: x, second: y}
}

// TODO: implement
func solve() bool {
  fmt.Sprintf("solve")
  return true
}

// TODO: implement
func eval(code string) bool {
    return false
}

// TODO: implement
func exists(quantifier []Var, predicate Pred) bool {
    return false
}
