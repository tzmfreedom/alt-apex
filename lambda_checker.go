package main

import (
	"github.com/tzmfreedom/alt-apex/parser"
)

type LambdaChecker struct {
	Lambdas []*parser.Lambda
}

func NewLambdaChecker() *LambdaChecker {
	return &LambdaChecker{}
}

func (v *LambdaChecker) VisitFile(n *parser.File) (interface{}, error) {
	_, err := n.Header.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, class := range n.Classes {
		_, err := class.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *LambdaChecker) VisitHeader(n *parser.Header) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitClass(n *parser.Class) (interface{}, error) {
	for _, d := range n.Declarations {
		_, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *LambdaChecker) VisitProperty(n *parser.Property) (interface{}, error) {
	if n.Expression != nil {
		_, err := n.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *LambdaChecker) VisitConstructor(n *parser.Constructor) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *LambdaChecker) VisitParameter(n *parser.Parameter) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitMethod(n *parser.Method) (interface{}, error) {
	return n.Statements.Accept(v)
}

func (v *LambdaChecker) VisitModifier(n *parser.Modifier) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitIf(n *parser.If) (interface{}, error) {
	_, err := n.IfStatement.Accept(v)
	if err != nil {
		return nil, err
	}
	if n.ElseStatement != nil {
		_, err := n.ElseStatement.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *LambdaChecker) VisitFor(n *parser.For) (interface{}, error) {
	_, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *LambdaChecker) VisitSwitch(n *parser.Switch) (interface{}, error) {
	for _, w := range n.Whens {
		_, err := w.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	if n.Else != nil {
		_, err := n.Else.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *LambdaChecker) VisitWhen(n *parser.When) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *LambdaChecker) VisitVariableDeclaration(n *parser.VariableDeclaration) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitBlock(n *parser.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *LambdaChecker) VisitTypeRef(n *parser.TypeRef) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitBoolean(n *parser.Boolean) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitInteger(n *parser.Integer) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitString(n *parser.String) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitIdentifier(n *parser.Identifier) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitMemberAccess(n *parser.MemberAccess) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitBinaryOperator(n *parser.BinaryOperator) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitName(n *parser.Name) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitMethodInvocation(n *parser.MethodInvocation) (interface{}, error) {
	return nil, nil
}

func (v *LambdaChecker) VisitWhile(n *parser.While) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *LambdaChecker) VisitLambda(n *parser.Lambda) (interface{}, error) {
	v.Lambdas = append(v.Lambdas, n)
	return nil, nil
}
