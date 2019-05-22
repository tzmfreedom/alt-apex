package main

import (
	"fmt"
	"github.com/tzmfreedom/alt-apex/parser"
	"strings"
)

type StringVisitor struct {}

func (v *StringVisitor) VisitFile(n *parser.File) (interface{}, error) {
	n.Header.Accept(v)
	stringArray := []string{}
	for _, class := range n.Classes {
		part, err := class.Accept(v)
		if err != nil {
			return nil, err
		}
		stringArray = append(stringArray, part.(string))
	}
	return strings.Join(stringArray, ","), nil
}

func (v *StringVisitor) VisitHeader(n *parser.Header) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitClass(n *parser.Class) (interface{}, error) {
	parameterStrings := make([]string, len(n.PrimaryConstructor.Parameters))
	for i, p := range n.PrimaryConstructor.Parameters {
		str, err := p.Accept(v)
		parameterStrings[i] = str.(string)
		if err != nil {
			return nil, err
		}
	}
	properties := make([]string, len(n.PrimaryConstructor.Parameters))
	for i, p := range n.PrimaryConstructor.Parameters {
		properties[i] = fmt.Sprintf("this.%s = %s;", p.Identifier, p.Identifier)
	}

	primaryConstructor := fmt.Sprintf(`public %s(%s) {
%s
}`, n.Name, strings.Join(parameterStrings, ", "), strings.Join(properties, "\n"))

	return fmt.Sprintf(`class %s {
%s
}`, n.Name, primaryConstructor), nil
}

func (v *StringVisitor) VisitProperty(n *parser.Property) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitConstructor(n *parser.Constructor) (interface{}, error) {
	parameterStrings := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		str, err := p.Accept(v)
		parameterStrings[i] = str.(string)
		if err != nil {
			return nil, err
		}
	}
	blockStr, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`%s %s`, strings.Join(parameterStrings, ", "), blockStr), nil
}

func (v *StringVisitor) VisitParameter(n *parser.Parameter) (interface{}, error) {
	typeRef, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s %s", typeRef, n.Identifier), nil
}

func (v *StringVisitor) VisitMethod(n *parser.Method) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitModifier(n *parser.Modifier) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitIf(n *parser.If) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitFor(n *parser.For) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitSwitch(n *parser.Switch) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitWhen(n *parser.When) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitVariableDeclaration(n *parser.VariableDeclaration) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitBlock(n *parser.Block) (interface{}, error) {
	statements := make([]string, len(n.Statements))
	for i, stmt := range n.Statements {
		str, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
		statements[i] = str.(string)
	}
	return fmt.Sprintf(`{
%s
}`, strings.Join(statements, "\n")), nil
}

func (v *StringVisitor) VisitTypeRef(n *parser.TypeRef) (interface{}, error) {
	return strings.Join(n.Name, "."), nil
}

func (v *StringVisitor) VisitBoolean(n *parser.Boolean) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitInteger(n *parser.Integer) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitString(n *parser.String) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitIdentifier(n *parser.Identifier) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitMemberAccess(n *parser.MemberAccess) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitBinaryOperator(n *parser.BinaryOperator) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitName(n *parser.Name) (interface{}, error) {
	return nil, nil
}

