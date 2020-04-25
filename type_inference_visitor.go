package main

import (
	"fmt"
	"github.com/tzmfreedom/alt-apex/parser"
	"regexp"
	"strings"
)

type TypeInferenceVisitor struct {
}

func NewTypeInferenceVisitor() *TypeInferenceVisitor {
	return &TypeInferenceVisitor{}
}

func (v *TypeInferenceVisitor) VisitFile(n *parser.File) (interface{}, error) {
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

func (v *TypeInferenceVisitor) VisitHeader(n *parser.Header) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitClass(n *parser.Class) (interface{}, error) {
	for _, d := range n.Declarations {
		_, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitProperty(n *parser.Property) (interface{}, error) {
	if n.Expression != nil {
		typeRef, err := n.Expression.Accept(v)
		if n.TypeRef == nil {
			if typeRef == nil {
				debug(n)
			}
			ref := typeRef.(*parser.TypeRef)
			listOfRegexp := regexp.MustCompile("listOf\\.(.+)")
			mapOfRegexp := regexp.MustCompile("mutableMap\\.(.+)\\.(.+)")
			methodName := strings.Join(ref.Name, ".")
			if methodName == "Database.insert" {
				n.TypeRef = &parser.TypeRef{
					Name: []string{"SObject"},
				}
			} else if listOfRegexp.MatchString(methodName) {
				values := listOfRegexp.FindAllStringSubmatch(methodName, -1)
				n.TypeRef = &parser.TypeRef{
					Name: []string{
						fmt.Sprintf("List<%s>", values[0][1]),
					},
				}
			} else if mapOfRegexp.MatchString(methodName) {
				values := mapOfRegexp.FindAllStringSubmatch(methodName, -1)
				n.TypeRef = &parser.TypeRef{
					Name: []string{
						fmt.Sprintf("Map<%s, %s>", values[0][1], values[0][2]),
					},
				}
			} else {
				n.TypeRef = ref
			}
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitConstructor(n *parser.Constructor) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *TypeInferenceVisitor) VisitParameter(n *parser.Parameter) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitMethod(n *parser.Method) (interface{}, error) {
	return n.Statements.Accept(v)
}

func (v *TypeInferenceVisitor) VisitModifier(n *parser.Modifier) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitIf(n *parser.If) (interface{}, error) {
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

func (v *TypeInferenceVisitor) VisitFor(n *parser.For) (interface{}, error) {
	_, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitSwitch(n *parser.Switch) (interface{}, error) {
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

func (v *TypeInferenceVisitor) VisitWhen(n *parser.When) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *TypeInferenceVisitor) VisitVariableDeclaration(n *parser.VariableDeclaration) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitBlock(n *parser.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitTypeRef(n *parser.TypeRef) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitBoolean(n *parser.Boolean) (interface{}, error) {
	return &parser.TypeRef{
		Name: []string{"Boolean"},
	}, nil
}

func (v *TypeInferenceVisitor) VisitInteger(n *parser.Integer) (interface{}, error) {
	return &parser.TypeRef{
		Name: []string{"Integer"},
	}, nil
}

func (v *TypeInferenceVisitor) VisitString(n *parser.String) (interface{}, error) {
	return &parser.TypeRef{
		Name: []string{"String"},
	}, nil
}

func (v *TypeInferenceVisitor) VisitIdentifier(n *parser.Identifier) (interface{}, error) {
	return &parser.TypeRef{
		Name: []string{n.Value},
	}, nil
}

func (v *TypeInferenceVisitor) VisitMemberAccess(n *parser.MemberAccess) (interface{}, error) {
	l, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	r, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	identifiers := append(l.(*parser.TypeRef).Name, r.(*parser.TypeRef).Name...)
	return &parser.TypeRef{
		Name: identifiers,
	}, nil
}

func (v *TypeInferenceVisitor) VisitBinaryOperator(n *parser.BinaryOperator) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitName(n *parser.Name) (interface{}, error) {
	return nil, nil
}

func (v *TypeInferenceVisitor) VisitMethodInvocation(n *parser.MethodInvocation) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	if exp == nil {
		return nil, nil
	}
	if n.Type != nil {
		name := append(
			exp.(*parser.TypeRef).Name,
			n.Type[0].Name...
		)
		if len(n.Type) > 1 {
			name = append(name, n.Type[1].Name...)
		}
		return &parser.TypeRef{
			Name: name,
		}, nil
	}
	return exp, nil
}

func (v *TypeInferenceVisitor) VisitWhile(n *parser.While) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *TypeInferenceVisitor) VisitLambda(n *parser.Lambda) (interface{}, error) {
	for _, stmt := range n.Statements {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return &parser.TypeRef{
		Name: []string{n.GetClassName()},
	}, nil
}

func (v *TypeInferenceVisitor) VisitReturn(n *parser.Return) (interface{}, error) {
	n.Expression.Accept(v)
	return nil, nil
}
