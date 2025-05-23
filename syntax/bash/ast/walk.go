// Copyright 2025 The Hulo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package ast

// A Visitor's Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(node Node) (w Visitor)
}

func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *Comment:
		// nothing to do

	case *Ident:
		// nothing to do

	case *ReturnStmt:
		Walk(v, n)

	case *BreakStmt:
		Walk(v, n)

	case *ContinueStmt:
		Walk(v, n)

	case *WhileStmt:
		Walk(v, n)

	case *UntilStmt:
		Walk(v, n)

	case *ForInStmt:
		Walk(v, n)

	case *ForStmt:
		Walk(v, n)

	case *IfStmt:
		Walk(v, n)

	case *CaseStmt:
		Walk(v, n)

	case *SelectStmt:
		Walk(v, n)

	case *ExprStmt:
		Walk(v, n)

	case *AssignStmt:
		Walk(v, n)
	}
}

type inspector func(Node) bool

func (f inspector) Visit(node Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}

func Inspect(node Node, f func(Node) bool) {
	Walk(inspector(f), node)
}
