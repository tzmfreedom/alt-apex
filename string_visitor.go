package main

import (
	"fmt"
	"github.com/tzmfreedom/alt-apex/parser"
	"strconv"
	"strings"
)

var typeMapper = map[string]string{
	"int": "Integer",
}

type StringVisitor struct {
	NameSpace string
	Lambdas   []*parser.Lambda
}

func (v *StringVisitor) VisitFile(n *parser.File) (interface{}, error) {
	namespace, err := n.Header.Accept(v)
	if err != nil {
		return nil, err
	}
	v.NameSpace = namespace.(string)
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
	namespaces := make([]string, len(n.Value))
	for i, val := range n.Value {
		namespaces[i] = strings.ToUpper(string(val[0])) + strings.ToLower(val[1:])
	}
	return strings.Join(namespaces, "_"), nil
}

func (v *StringVisitor) VisitClass(n *parser.Class) (interface{}, error) {
	modifierString, err := v.modifierString(n.Modifiers)
	if err != nil {
		return nil, err
	}
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

	declarations := make([]string, len(n.Declarations))
	for i, decl := range n.Declarations {
		str, err := decl.Accept(v)
		if err != nil {
			return nil, err
		}
		declarations[i] = str.(string)
	}

	lambdas := make([]string, len(v.Lambdas))
	for _, l := range v.Lambdas {
		lambda, err := v.createLambdaMethod(l.GetClassName(), l)
		if err != nil {
			return nil, err
		}
		lambdas = append(lambdas, lambda)
	}

	propertyDeclarations := make([]string, len(n.PrimaryConstructor.Parameters))
	for i, p := range n.PrimaryConstructor.Parameters {
		s, err := p.TypeRef.Accept(v)
		if err != nil {
			return nil, err
		}
		if p.Expression != nil {
			e, err := p.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			propertyDeclarations[i] = fmt.Sprintf("%s %s = %s;", s.(string), p.Identifier, e.(string))
		} else {
			propertyDeclarations[i] = fmt.Sprintf("%s %s;", s.(string), p.Identifier)
		}
	}

	return fmt.Sprintf(`%s class %s_%s {
%s
%s
%s
%s
}`,
		modifierString,
		v.NameSpace,
		n.Name,
		strings.Join(propertyDeclarations, "\n"),
		primaryConstructor,
		strings.Join(declarations, "\n"),
		strings.Join(lambdas, "\n"),
	), nil
}

func (v *StringVisitor) VisitProperty(n *parser.Property) (interface{}, error) {
	r, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	typeRef := r.(string)
	expression := ""
	if n.Expression != nil {
		e, err := n.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		expression = " = " + e.(string)
	}
	return fmt.Sprintf(`%s %s%s`, typeRef, n.Name, expression), nil
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
	return fmt.Sprintf(`public %s(%s) %s`, n.Name, strings.Join(parameterStrings, ", "), blockStr), nil
}

func (v *StringVisitor) VisitParameter(n *parser.Parameter) (interface{}, error) {
	typeRef, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s %s", typeRef, n.Identifier), nil
}

func (v *StringVisitor) VisitMethod(n *parser.Method) (interface{}, error) {
	modifierString, err := v.modifierString(n.AccessModifiers)
	if err != nil {
		return nil, err
	}
	parameterStrings := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		str, err := p.Accept(v)
		parameterStrings[i] = str.(string)
		if err != nil {
			return nil, err
		}
	}
	blockStr, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s %s(%s) %s",
		modifierString,
		strings.Join(n.Identifier, "."),
		strings.Join(parameterStrings, ", "),
		blockStr,
	), nil
}

func (v *StringVisitor) VisitModifier(n *parser.Modifier) (interface{}, error) {
	return n.Name, nil
}

func (v *StringVisitor) VisitIf(n *parser.If) (interface{}, error) {
	cond, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	ifStmt, err := n.IfStatement.Accept(v)
	if err != nil {
		return nil, err
	}
	elseStmt := ""
	if n.ElseStatement != nil {
		e, err := n.ElseStatement.Accept(v)
		if err != nil {
			return nil, err
		}
		elseStmt = " else " + e.(string)
	}
	return fmt.Sprintf(`if (%s) %s%s`, cond.(string), ifStmt.(string), elseStmt), nil
}

func (v *StringVisitor) VisitFor(n *parser.For) (interface{}, error) {
	expression, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`for (%s : %s) %s`, n.Identifier, expression, block), nil
}

func (v *StringVisitor) VisitSwitch(n *parser.Switch) (interface{}, error) {
	expression, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	whens := make([]string, len(n.Whens))
	for i, when := range n.Whens {
		w, err := when.Accept(v)
		if err != nil {
			return nil, err
		}
		whens[i] = w.(string)
	}
	if n.Else != nil {
		e, err := n.Else.Accept(v)
		if err != nil {
			return nil, err
		}
		whens = append(whens, fmt.Sprintf(`when else %s`, e.(string)))
	}
	return fmt.Sprintf(`switch on %s {
%s
}`, expression.(string), strings.Join(whens, "\n")), nil
}

func (v *StringVisitor) VisitWhen(n *parser.When) (interface{}, error) {
	conditions := make([]string, len(n.Conditions))
	for i, condition := range n.Conditions {
		c, err := condition.Accept(v)
		if err != nil {
			return nil, err
		}
		conditions[i] = c.(string)
	}
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`when %s %s`, strings.Join(conditions, ", "), block.(string)), nil
}

func (v *StringVisitor) VisitVariableDeclaration(n *parser.VariableDeclaration) (interface{}, error) {
	return nil, nil
}

func (v *StringVisitor) VisitBlock(n *parser.Block) (interface{}, error) {
	block, err := v.createStatementsString(n.Statements)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`{
%s
}`, block), nil
}

func (v *StringVisitor) VisitTypeRef(n *parser.TypeRef) (interface{}, error) {
	name := strings.Join(n.Name, ".")
	lowered := strings.ToLower(name)
	if mapped, ok := typeMapper[lowered]; ok {
		return mapped, nil
	}
	return name, nil
}

func (v *StringVisitor) VisitBoolean(n *parser.Boolean) (interface{}, error) {
	if n.Value {
		return "true", nil
	}
	return "false", nil
}

func (v *StringVisitor) VisitInteger(n *parser.Integer) (interface{}, error) {
	return strconv.Itoa(n.Value), nil
}

func (v *StringVisitor) VisitString(n *parser.String) (interface{}, error) {
	return fmt.Sprintf("'%s'", n.Value), nil
}

func (v *StringVisitor) VisitIdentifier(n *parser.Identifier) (interface{}, error) {
	return n.Value, nil
}

func (v *StringVisitor) VisitMemberAccess(n *parser.MemberAccess) (interface{}, error) {
	left, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s.%s", left.(string), right.(string)), nil
}

func (v *StringVisitor) VisitBinaryOperator(n *parser.BinaryOperator) (interface{}, error) {
	left, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`%s %s %s`, left.(string), n.Operator, right.(string)), nil
}

func (v *StringVisitor) VisitName(n *parser.Name) (interface{}, error) {
	return strings.Join(n.Value, "."), nil
}

func (v *StringVisitor) VisitMethodInvocation(n *parser.MethodInvocation) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	parameterStrings := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		parameterString, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		parameterStrings[i] = parameterString.(string)
	}
	if ident, ok := n.Expression.(*parser.Identifier); ok {
		val := string(ident.Value[0])
		if strings.ToUpper(val) == val {
			return fmt.Sprintf("new %s(%s)", exp.(string), strings.Join(parameterStrings, ", ")), nil
		}
	}
	values := strings.Split(exp.(string), ".")
	if len(values) > 0 {
		val := string(values[len(values)-1][0])
		if strings.ToUpper(val) == val {
			return fmt.Sprintf("new %s(%s)", exp.(string), strings.Join(parameterStrings, ", ")), nil
		}
	}
	if exp.(string) == "listOf" {
		if n.Type != nil {
			return fmt.Sprintf("new List<%s>()", strings.Join(n.Type[0].Name, ".")), nil
		}
	}
	if exp.(string) == "mutableMap" {
		if n.Type != nil {
			return fmt.Sprintf(
				"new Map<%s, %s>()",
				strings.Join(n.Type[0].Name, "."),
				strings.Join(n.Type[1].Name, "."),
			), nil
		}
	}
	return fmt.Sprintf("%s(%s)", exp.(string), strings.Join(parameterStrings, ", ")), nil
}

func (v *StringVisitor) VisitWhile(n *parser.While) (interface{}, error) {
	cond, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	block, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`while (%s) %s`, cond.(string), block.(string)), nil
}

func (v *StringVisitor) VisitLambda(n *parser.Lambda) (interface{}, error) {
	return fmt.Sprintf(`new %s()`, n.GetClassName()), nil
}

func (v *StringVisitor) VisitReturn(n *parser.Return) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(`return %s`, exp.(string)), nil
}

func (v *StringVisitor) modifierString(modifiers []*parser.Modifier) (string, error) {
	modifierStrings := make([]string, len(modifiers))
	for i, accessModifier := range modifiers {
		modifierString, err := accessModifier.Accept(v)
		if err != nil {
			return "", err
		}
		modifierStrings[i] = modifierString.(string)
	}
	return strings.Join(modifierStrings, " "), nil
}

func (v *StringVisitor) createLambdaMethod(name string, n *parser.Lambda) (string, error) {
	parameters := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		parameter, err := p.Accept(v)
		if err != nil {
			return "", err
		}
		parameters[i] = parameter.(string)
	}
	statements := make([]string, len(n.Statements))
	for i, s := range n.Statements {
		stmt, err := s.Accept(v)
		if err != nil {
			return "", err
		}
		statements[i] = stmt.(string)
	}
	block, err := v.createStatementsString(n.Statements)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`class %s {
    public run(%s) {
    %s
    }
}`,
		name,
		strings.Join(parameters, ", "),
		block,
	), nil
}

func (v *StringVisitor) createStatementsString(statements []parser.Node) (string, error) {
	statementStrings := make([]string, len(statements))
	for i, stmt := range statements {
		str, err := stmt.Accept(v)
		if err != nil {
			return "", err
		}
		statementStrings[i] = str.(string)
		switch stmt.(type) {
		case *parser.If, *parser.For, *parser.While, *parser.Switch:
		default:
			statementStrings[i] += ";"
		}
	}
	return strings.Join(statementStrings, "\n"), nil
}
