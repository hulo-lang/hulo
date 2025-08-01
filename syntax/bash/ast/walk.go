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

	case *CommentGroup:
		for _, c := range n.List {
			Walk(v, c)
		}

	case *Ident:
		// nothing to do

	case *Word:
		// nothing to do

	// Statements
	case *FuncDecl:
		Walk(v, n.Name)
		Walk(v, n.Body)

	case *ReturnStmt:
		if n.X != nil {
			Walk(v, n.X)
		}

	case *WhileStmt:
		Walk(v, n.Cond)
		Walk(v, n.Body)

	case *UntilStmt:
		Walk(v, n.Cond)
		Walk(v, n.Body)

	case *ForInStmt:
		Walk(v, n.Var)
		Walk(v, n.List)
		Walk(v, n.Body)

	case *ForStmt:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		if n.Cond != nil {
			Walk(v, n.Cond)
		}
		if n.Post != nil {
			Walk(v, n.Post)
		}
		Walk(v, n.Body)

	case *IfStmt:
		Walk(v, n.Cond)
		Walk(v, n.Body)
		if n.Else != nil {
			Walk(v, n.Else)
		}

	case *CaseStmt:
		Walk(v, n.X)
		for _, pattern := range n.Patterns {
			for _, cond := range pattern.Conds {
				Walk(v, cond)
			}
			Walk(v, pattern.Body)
		}
		if n.Else != nil {
			Walk(v, n.Else)
		}

	case *SelectStmt:
		Walk(v, n.Var)
		Walk(v, n.List)
		Walk(v, n.Body)

	case *ExprStmt:
		Walk(v, n.X)

	case *AssignStmt:
		Walk(v, n.Lhs)
		Walk(v, n.Rhs)

	case *BlockStmt:
		for _, s := range n.List {
			Walk(v, s)
		}

	// Expressions
	case *BinaryExpr:
		Walk(v, n.X)
		Walk(v, n.Y)

	case *CmdExpr:
		Walk(v, n.Name)
		for _, arg := range n.Recv {
			Walk(v, arg)
		}

	case *CmdListExpr:
		for _, cmd := range n.Cmds {
			Walk(v, cmd)
		}

	case *CmdGroup:
		for _, expr := range n.List {
			Walk(v, expr)
		}

	case *TestExpr:
		Walk(v, n.X)

	case *ExtendedTestExpr:
		Walk(v, n.X)

	case *ArithEvalExpr:
		Walk(v, n.X)

	case *CmdSubst:
		Walk(v, n.X)

	case *ProcSubst:
		Walk(v, n.X)

	case *ArithExpr:
		Walk(v, n.X)

	case *VarExpExpr:
		Walk(v, n.X)

	case *ParamExpExpr:
		Walk(v, n.Var)
		switch exp := n.ParamExp.(type) {
		case *DefaultValExp:
			Walk(v, exp.Val)
		case *DefaultValAssignExp:
			Walk(v, exp.Val)
		case *NonNullCheckExp:
			Walk(v, exp.Val)
		case *NonNullExp:
			Walk(v, exp.Val)
		case *DelPrefix:
			Walk(v, exp.Val)
		case *DelSuffix:
			Walk(v, exp.Val)
		case *ReplaceExp:
			// Old and New are strings, not expressions
		case *ReplacePrefixExp:
			// Old and New are strings, not expressions
		case *ReplaceSuffixExp:
			// Old and New are strings, not expressions
		}

	case *IndexExpr:
		Walk(v, n.X)
		Walk(v, n.Y)

	case *ArrExpr:
		for _, var_ := range n.Vars {
			Walk(v, var_)
		}

	case *UnaryExpr:
		Walk(v, n.X)

	case *PipelineExpr:
		for _, cmd := range n.Cmds {
			Walk(v, cmd)
		}

	case *Redirect:
		if n.N != nil {
			Walk(v, n.N)
		}
		Walk(v, n.Word)

	// File structure
	case *File:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		for _, s := range n.Stmts {
			Walk(v, s)
		}
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
