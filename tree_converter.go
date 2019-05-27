package main

import (
	"github.com/tzmfreedom/alt-apex/parser"
)

type TreeConverter struct {
	NameSpace string
}

func NewTreeConverter() *TreeConverter {
	return &TreeConverter{}
}

func (v *TreeConverter) VisitFile(n *parser.File) (interface{}, error) {
	header, err := n.Header.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Header = header.(*parser.Header)
	for i, class := range n.Classes {
		converted, err := class.Accept(v)
		if err != nil {
			return nil, err
		}
		n.Classes[i] = converted.(*parser.Class)
	}
	return n, nil
}

func (v *TreeConverter) VisitHeader(n *parser.Header) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitClass(n *parser.Class) (interface{}, error) {
	for i, d := range n.Declarations {
		decl, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
		n.Declarations[i] = decl.(parser.Node)
	}
	return n, nil
}

func (v *TreeConverter) VisitProperty(n *parser.Property) (interface{}, error) {
	if n.Expression != nil {
		if ifStmt, ok := n.Expression.(*parser.If); ok {
			ifStatements := ifStmt.IfStatement.Statements
			lastIndex := len(ifStatements) - 1
			ifStatements[lastIndex] = &parser.BinaryOperator{
				Left: &parser.Identifier{
					Value: n.Name,
				},
				Right:    ifStatements[lastIndex].(parser.Node),
				Operator: "=",
			}
			if ifStmt.ElseStatement != nil {
				elseStatements := ifStmt.ElseStatement.Statements
				lastIndex := len(elseStatements) - 1
				elseStatements[lastIndex] = &parser.BinaryOperator{
					Left: &parser.Identifier{
						Value: n.Name,
					},
					Right:    elseStatements[lastIndex].(parser.Node),
					Operator: "=",
				}
			}
			n.Expression = nil
			return []parser.Node{
				n,
				ifStmt,
			}, nil
		} else if switchStmt, ok := n.Expression.(*parser.Switch); ok {
			for _, w := range switchStmt.Whens {
				statements := w.Block.Statements
				lastIndex := len(statements) - 1
				statements[lastIndex] = &parser.BinaryOperator{
					Left: &parser.Identifier{
						Value: n.Name,
					},
					Right:    statements[lastIndex].(parser.Node),
					Operator: "=",
				}
			}
			if switchStmt.Else != nil {
				elseStatements := switchStmt.Else.Statements
				lastIndex := len(elseStatements)-1
				elseStatements[lastIndex] = &parser.BinaryOperator{
					Left: &parser.Identifier{
						Value: n.Name,
					},
					Right: elseStatements[lastIndex].(parser.Node),
					Operator: "=",
				}
			}
			n.Expression = nil
			return []parser.Node{
				n,
				switchStmt,
			}, nil
		} else {
			expression, err := n.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			n.Expression = expression.(parser.Node)
		}
	}
	return n, nil
}

func (v *TreeConverter) VisitConstructor(n *parser.Constructor) (interface{}, error) {
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Block = block.(*parser.Block)
	return n, nil
}

func (v *TreeConverter) VisitParameter(n *parser.Parameter) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitMethod(n *parser.Method) (interface{}, error) {
	block, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Statements = block.(*parser.Block)
	return n, nil
}

func (v *TreeConverter) VisitModifier(n *parser.Modifier) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitIf(n *parser.If) (interface{}, error) {
	ifStmt, err := n.IfStatement.Accept(v)
	if err != nil {
		return nil, err
	}
	n.IfStatement = ifStmt.(*parser.Block)
	if n.ElseStatement != nil {
		elseStmt, err := n.ElseStatement.Accept(v)
		if err != nil {
			return nil, err
		}
		n.ElseStatement = elseStmt.(*parser.Block)
	}
	return n, nil
}

func (v *TreeConverter) VisitFor(n *parser.For) (interface{}, error) {
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Block = block.(*parser.Block)
	return n, nil
}

func (v *TreeConverter) VisitSwitch(n *parser.Switch) (interface{}, error) {
	for i, w := range n.Whens {
		when, err := w.Accept(v)
		if err != nil {
			return nil, err
		}
		n.Whens[i] = when.(*parser.When)
	}
	if n.Else != nil {
		e, err := n.Else.Accept(v)
		if err != nil {
			return nil, err
		}
		n.Else = e.(*parser.Block)
	}
	return n, nil
}

func (v *TreeConverter) VisitWhen(n *parser.When) (interface{}, error) {
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Block = block.(*parser.Block)
	return n, nil
}

func (v *TreeConverter) VisitVariableDeclaration(n *parser.VariableDeclaration) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitBlock(n *parser.Block) (interface{}, error) {
	statements := []parser.Node{}
	for _, stmt := range n.Statements {
		newStmt, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
		switch ret := newStmt.(type) {
		case parser.Node:
			statements = append(statements, ret)
		case []parser.Node:
			for _, r := range ret {
				statements = append(statements, r)
			}
		}
	}
	n.Statements = statements
	return n, nil
}

func (v *TreeConverter) VisitTypeRef(n *parser.TypeRef) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitBoolean(n *parser.Boolean) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitInteger(n *parser.Integer) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitString(n *parser.String) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitIdentifier(n *parser.Identifier) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitMemberAccess(n *parser.MemberAccess) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitBinaryOperator(n *parser.BinaryOperator) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitName(n *parser.Name) (interface{}, error) {
	return n, nil
}

func (v *TreeConverter) VisitMethodInvocation(n *parser.MethodInvocation) (interface{}, error) {
	// TODO: impl
	if ident, ok := n.Expression.(*parser.Identifier); ok && ident.Value == "sum" {
		n.Expression = &parser.Name{
			Value: []string{
				ident.Value,
				"run",
			},
		}
	}
	for _, p := range n.Parameters {
		p.Accept(v)
	}
	return n, nil
}

func (v *TreeConverter) VisitWhile(n *parser.While) (interface{}, error) {
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Block = block.(*parser.Block)
	return n, nil
}

func (v *TreeConverter) VisitLambda(n *parser.Lambda) (interface{}, error) {
	lastIndex := len(n.Statements)-1
	n.Statements[lastIndex] = &parser.Return{
		Expression: n.Statements[lastIndex],
	}
	return n, nil
}

func (v *TreeConverter) VisitReturn(n *parser.Return) (interface{}, error) {
	n.Expression.Accept(v)
	return n, nil
}
