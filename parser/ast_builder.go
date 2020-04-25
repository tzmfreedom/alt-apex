package parser // KotlinParser
import (
	"github.com/k0kubun/pp"
	"strconv"
	"strings"
)

// A complete Visitor for a parse tree produced by KotlinParser.
type AstBuilder struct {
	*BaseKotlinParserVisitor
}

func (v *AstBuilder) VisitKotlinFile(ctx *KotlinFileContext) interface{} {
	preamble := ctx.Preamble().Accept(v)
	klass := ctx.AllTopLevelObject()[0].Accept(v)
	return &File{
		Header:  preamble.(*Header),
		Classes: []*Class{klass.(*Class)},
	}
}

func (v *AstBuilder) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPreamble(ctx *PreambleContext) interface{} {
	return ctx.PackageHeader().Accept(v)
}

func (v *AstBuilder) VisitFileAnnotations(ctx *FileAnnotationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFileAnnotation(ctx *FileAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPackageHeader(ctx *PackageHeaderContext) interface{} {
	return &Header{
		Value: ctx.Identifier().Accept(v).([]string),
	}
}

func (v *AstBuilder) VisitImportList(ctx *ImportListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitImportHeader(ctx *ImportHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitImportAlias(ctx *ImportAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTopLevelObject(ctx *TopLevelObjectContext) interface{} {
	return ctx.ClassDeclaration().Accept(v)
}

func (v *AstBuilder) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	var modifiers []*Modifier
	if m := ctx.ModifierList(); m != nil {
		modifiers = m.Accept(v).([]*Modifier)
	} else {
		modifiers = []*Modifier{publicModifier}
	}
	name := ctx.SimpleIdentifier().Accept(v).(string)
	declarations := ctx.ClassBody().Accept(v).([]Node)
	var constructor *Constructor
	if c := ctx.PrimaryConstructor(); c != nil {
		constructor = c.Accept(v).(*Constructor)
	}
	for _, decl := range declarations {
		switch d := decl.(type) {
		case *Constructor:
			d.Name = name
		}
	}
	return &Class{
		Modifiers:          modifiers,
		Name:               name,
		PrimaryConstructor: constructor,
		Declarations:       declarations,
	}
}

func (v *AstBuilder) VisitPrimaryConstructor(ctx *PrimaryConstructorContext) interface{} {
	// ctx.ModifierList()
	return ctx.ClassParameters().Accept(v).(*Constructor)
}

func (v *AstBuilder) VisitClassParameters(ctx *ClassParametersContext) interface{} {
	parameters := make([]*Parameter, len(ctx.AllClassParameter()))
	for i, p := range ctx.AllClassParameter() {
		parameters[i] = p.Accept(v).(*Parameter)
	}
	return &Constructor{
		Parameters: parameters,
	}
}

func (v *AstBuilder) VisitClassParameter(ctx *ClassParameterContext) interface{} {
	name := ctx.SimpleIdentifier().Accept(v).(string)
	ktype := ctx.Ktype().Accept(v).(*TypeRef)
	var expression Node
	if e := ctx.Expression(); e != nil {
		expression = e.Accept(v).(Node)
	}
	var modifiers []*Modifier
	if m := ctx.ModifierList(); m != nil {
		modifiers = m.Accept(v).([]*Modifier)
	}
	return &Parameter{
		Modifiers:  modifiers,
		Identifier: name,
		TypeRef:    ktype,
		Expression: expression,
	}
}

func (v *AstBuilder) VisitDelegationSpecifiers(ctx *DelegationSpecifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitDelegationSpecifier(ctx *DelegationSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstructorInvocation(ctx *ConstructorInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitExplicitDelegation(ctx *ExplicitDelegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassBody(ctx *ClassBodyContext) interface{} {
	declarations := make([]Node, len(ctx.AllClassMemberDeclaration()))
	for i, decl := range ctx.AllClassMemberDeclaration() {
		declarations[i] = decl.Accept(v).(Node)
	}
	return declarations
}

func (v *AstBuilder) VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{} {
	if decl := ctx.PropertyDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.FunctionDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.SecondaryConstructor(); decl != nil {
		return decl.Accept(v)
	}
	return nil
}

func (v *AstBuilder) VisitAnonymousInitializer(ctx *AnonymousInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSecondaryConstructor(ctx *SecondaryConstructorContext) interface{} {
	//modifiers := ctx.ModifierList().Accept(v).([]*Modifier)
	parameters := ctx.FunctionValueParameters().Accept(v).([]*Parameter)
	block := ctx.Block().Accept(v).(*Block)
	return &Constructor{
		Parameters: parameters,
		Block:      block,
	}
}

func (v *AstBuilder) VisitConstructorDelegationCall(ctx *ConstructorDelegationCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumClassBody(ctx *EnumClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumEntries(ctx *EnumEntriesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumEntry(ctx *EnumEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	// TODO: impl
	var modifiers []*Modifier
	if m := ctx.ModifierList(); m != nil {
		modifiers = ctx.ModifierList().Accept(v).([]*Modifier)
	} else {
		modifiers = []*Modifier{publicModifier}
	}
	identifier := ctx.Identifier().Accept(v).([]string)
	parameters := ctx.FunctionValueParameters().Accept(v).([]*Parameter)
	block := ctx.FunctionBody().Accept(v).(*Block)
	return &Method{
		AccessModifiers: modifiers,
		Identifier:      identifier,
		Parameters:      parameters,
		Statements:      block,
	}
}

func (v *AstBuilder) VisitFunctionValueParameters(ctx *FunctionValueParametersContext) interface{} {
	parameters := make([]*Parameter, len(ctx.AllFunctionValueParameter()))
	for i, p := range ctx.AllFunctionValueParameter() {
		parameters[i] = p.Accept(v).(*Parameter)
	}
	return parameters
}

func (v *AstBuilder) VisitFunctionValueParameter(ctx *FunctionValueParameterContext) interface{} {
	var modifiers []*Modifier
	if modifierList := ctx.ModifierList(); modifierList != nil {
		modifiers = modifierList.Accept(v).([]*Modifier)
	}
	parameter := ctx.Parameter().Accept(v).(*Parameter)
	parameter.Modifiers = modifiers
	if exp := ctx.Expression(); exp != nil {
		parameter.Expression = exp.Accept(v).(Node)
	}
	return parameter
}

func (v *AstBuilder) VisitParameter(ctx *ParameterContext) interface{} {
	identifier := ctx.SimpleIdentifier().Accept(v).(string)
	typeRef := ctx.Ktype().Accept(v).(*TypeRef)
	return &Parameter{
		Identifier: identifier,
		TypeRef:    typeRef,
	}
}

func (v *AstBuilder) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	if b := ctx.Block(); b != nil {
		return b.Accept(v)
	}
	return &Block{
		Statements: []Node{
			ctx.Expression().Accept(v).(Node),
		},
	}
}

func (v *AstBuilder) VisitObjectDeclaration(ctx *ObjectDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCompanionObject(ctx *CompanionObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{} {
	// TODO: impl
	var modifiers []*Modifier
	if m := ctx.ModifierList(); m != nil {
		modifiers = ctx.ModifierList().Accept(v).([]*Modifier)
	} else {
		modifiers = []*Modifier{publicModifier}
	}
	//ktype := ctx.Ktype().Accept(v).(*TypeRef)
	var varDecl *VariableDeclaration
	if decl := ctx.VariableDeclaration(); decl != nil {
		varDecl = decl.Accept(v).(*VariableDeclaration)
	}
	expression := ctx.Expression().Accept(v).(Node)

	return &Property{
		Modifiers:  modifiers,
		TypeRef:    varDecl.TypeRef,
		Name:       varDecl.Identifier,
		Expression: expression,
	}
}

func (v *AstBuilder) VisitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	name := ctx.SimpleIdentifier().Accept(v).(string)
	var typeRef *TypeRef
	if t := ctx.Ktype(); t != nil {
		typeRef = t.Accept(v).(*TypeRef)
	}
	return &VariableDeclaration{
		Identifier: name,
		TypeRef:    typeRef,
	}
}

func (v *AstBuilder) VisitGetter(ctx *GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSetter(ctx *SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeAlias(ctx *TypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitKtype(ctx *KtypeContext) interface{} {
	if f := ctx.TypeReference(); f != nil {
		name := f.Accept(v).([]string)
		return &TypeRef{
			Name: name,
		}
	}
	return nil
}

func (v *AstBuilder) VisitTypeModifierList(ctx *TypeModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitParenthesizedType(ctx *ParenthesizedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitNullableType(ctx *NullableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	if ut := ctx.UserType(); ut != nil {
		return ut.Accept(v).([]string)
	}
	return nil
}

func (v *AstBuilder) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionTypeReceiver(ctx *FunctionTypeReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitUserType(ctx *UserTypeContext) interface{} {
	userTypes := make([]string, len(ctx.AllSimpleUserType()))
	for i, userType := range ctx.AllSimpleUserType() {
		userTypes[i] = userType.Accept(v).(string)
	}
	return userTypes
}

func (v *AstBuilder) VisitSimpleUserType(ctx *SimpleUserTypeContext) interface{} {
	return ctx.SimpleIdentifier().Accept(v).(string)
}

func (v *AstBuilder) VisitFunctionTypeParameters(ctx *FunctionTypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeConstraints(ctx *TypeConstraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeConstraint(ctx *TypeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitBlock(ctx *BlockContext) interface{} {
	return &Block{
		Statements: ctx.Statements().Accept(v).([]Node),
	}
}

func (v *AstBuilder) VisitStatements(ctx *StatementsContext) interface{} {
	statements := make([]Node, len(ctx.AllStatement()))
	for i, stmt := range ctx.AllStatement() {
		statements[i] = stmt.Accept(v).(Node)
	}
	return statements
}

func (v *AstBuilder) VisitStatement(ctx *StatementContext) interface{} {
	if decl := ctx.Declaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.BlockLevelExpression(); decl != nil {
		return decl.Accept(v)
	}
	panic("not pass")
}

func (v *AstBuilder) VisitBlockLevelExpression(ctx *BlockLevelExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *AstBuilder) VisitDeclaration(ctx *DeclarationContext) interface{} {
	if decl := ctx.ClassDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.FunctionDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.PropertyDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.TypeAlias(); decl != nil {
		return decl.Accept(v)
	}
	panic("not pass")
}

func (v *AstBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	// TODO: impl
	if len(ctx.AllDisjunction()) == 1 {
		return ctx.Disjunction(0).Accept(v)
	}
	exp := &BinaryOperator{
		Right: ctx.Disjunction(0).Accept(v).(Node),
	}
	disjunctions := ctx.AllDisjunction()
	for i, operator := range ctx.AllAssignmentOperator() {
		assignmentOperator := operator.Accept(v).(string)
		node := &BinaryOperator{}
		node.Left = exp.Right
		node.Operator = assignmentOperator
		node.Right = disjunctions[i+1].Accept(v).(Node)
		exp.Right = node
		exp = node
	}
	return exp
}

func (v *AstBuilder) VisitDisjunction(ctx *DisjunctionContext) interface{} {
	// TODO: impl
	return ctx.Conjunction(0).Accept(v)
}

func (v *AstBuilder) VisitConjunction(ctx *ConjunctionContext) interface{} {
	// TODO: impl
	return ctx.EqualityComparison(0).Accept(v)
}

func (v *AstBuilder) VisitEqualityComparison(ctx *EqualityComparisonContext) interface{} {
	// TODO: impl
	if len(ctx.AllComparison()) == 1 {
		return ctx.Comparison(0).Accept(v)
	}
	exp := &BinaryOperator{
		Right: ctx.Comparison(0).Accept(v).(Node),
	}
	comparisons := ctx.AllComparison()
	for i, operator := range ctx.AllEqualityOperation() {
		assignmentOperator := operator.Accept(v).(string)
		node := &BinaryOperator{}
		node.Left = exp.Right
		node.Operator = assignmentOperator
		node.Right = comparisons[i+1].Accept(v).(Node)
		exp.Right = node
		exp = node
	}
	return exp
}

func (v *AstBuilder) VisitComparison(ctx *ComparisonContext) interface{} {
	// TODO: impl
	return ctx.NamedInfix(0).Accept(v)
}

func (v *AstBuilder) VisitNamedInfix(ctx *NamedInfixContext) interface{} {
	return ctx.ElvisExpression(0).Accept(v)
}

func (v *AstBuilder) VisitElvisExpression(ctx *ElvisExpressionContext) interface{} {
	return ctx.InfixFunctionCall(0).Accept(v)
}

func (v *AstBuilder) VisitInfixFunctionCall(ctx *InfixFunctionCallContext) interface{} {
	return ctx.RangeExpression(0).Accept(v)
}

func (v *AstBuilder) VisitRangeExpression(ctx *RangeExpressionContext) interface{} {
	return ctx.AdditiveExpression(0).Accept(v)
}

func (v *AstBuilder) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	if len(ctx.AllMultiplicativeExpression()) == 1 {
		return ctx.MultiplicativeExpression(0).Accept(v)
	}
	exp := &BinaryOperator{
		Right: ctx.MultiplicativeExpression(0).Accept(v).(Node),
	}
	comparisons := ctx.AllMultiplicativeExpression()
	for i, operator := range ctx.AllAdditiveOperator() {
		assignmentOperator := operator.Accept(v).(string)
		node := &BinaryOperator{}
		node.Left = exp.Right
		node.Operator = assignmentOperator
		node.Right = comparisons[i+1].Accept(v).(Node)
		exp.Right = node
		exp = node
	}
	return exp
}

func (v *AstBuilder) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return ctx.TypeRHS(0).Accept(v)
}

func (v *AstBuilder) VisitTypeRHS(ctx *TypeRHSContext) interface{} {
	return ctx.PrefixUnaryExpression(0).Accept(v)
}

func (v *AstBuilder) VisitPrefixUnaryExpression(ctx *PrefixUnaryExpressionContext) interface{} {
	return ctx.PostfixUnaryExpression().Accept(v)
}

func (v *AstBuilder) VisitPostfixUnaryExpression(ctx *PostfixUnaryExpressionContext) interface{} {
	var expression Node
	if e := ctx.AtomicExpression(); e != nil {
		expression = e.Accept(v).(Node)
	}
	if len(ctx.AllPostfixUnaryOperation()) == 0 {
		return expression
	}
	if len(ctx.AllPostfixUnaryOperation()) > 1 {
		panic("not pass")
	}
	operator := ctx.PostfixUnaryOperation(0).Accept(v)
	if invoke, ok := operator.(*MethodInvocation); ok {
		if invoke.Expression != nil {
			expression = &MemberAccess{
				Left:  expression,
				Right: invoke.Expression,
			}
		}
		invoke.Expression = expression
		return invoke
	}
	return &MemberAccess{
		Left:  expression,
		Right: operator.(Node),
	}
}

func (v *AstBuilder) VisitAtomicExpression(ctx *AtomicExpressionContext) interface{} {
	// TODO: impl
	if lit := ctx.LiteralConstant(); lit != nil {
		return lit.Accept(v)
	}
	if ident := ctx.SimpleIdentifier(); ident != nil {
		return &Identifier{
			Value: ident.Accept(v).(string),
		}
	}
	if exp := ctx.ThisExpression(); exp != nil {
		return &Identifier{
			Value: exp.GetText(),
		}
	}
	if exp := ctx.ConditionalExpression(); exp != nil {
		return exp.Accept(v)
	}
	if exp := ctx.LoopExpression(); exp != nil {
		return exp.Accept(v)
	}
	if exp := ctx.FunctionLiteral(); exp != nil {
		return exp.Accept(v)
	}
	debug(ctx.GetStart().GetColumn())
	debug(ctx.GetStart().GetLine())
	panic("not pass")
}

func (v *AstBuilder) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}


type callSuffix struct {
	Type []*TypeRef
	Args []Node
}

func (v *AstBuilder) VisitCallSuffix(ctx *CallSuffixContext) interface{} {
	var typeArgs []*TypeRef
	var valueArgs []Node
	if args := ctx.TypeArguments(); args != nil {
		typeArgs = args.Accept(v).([]*TypeRef)
	}
	if vArgs := ctx.ValueArguments(); vArgs != nil {
		valueArgs = vArgs.Accept(v).([]Node)
	}
	return callSuffix{typeArgs, valueArgs}
}

func (v *AstBuilder) VisitAnnotatedLambda(ctx *AnnotatedLambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitValueArguments(ctx *ValueArgumentsContext) interface{} {
	args := make([]Node, len(ctx.AllValueArgument()))
	for i, arg := range ctx.AllValueArgument() {
		args[i] = arg.Accept(v).(Node)
	}
	return args
}

func (v *AstBuilder) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	tps := make([]*TypeRef, len(ctx.AllTypeProjection()))
	for i, t := range ctx.AllTypeProjection() {
		tps[i] = t.Accept(v).(*TypeRef)
	}
	return tps
}

func (v *AstBuilder) VisitTypeProjection(ctx *TypeProjectionContext) interface{} {
	return ctx.Ktype().Accept(v)
}

func (v *AstBuilder) VisitTypeProjectionModifierList(ctx *TypeProjectionModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitValueArgument(ctx *ValueArgumentContext) interface{} {
	if ident := ctx.SimpleIdentifier(); ident != nil {
		panic("not impl")
	}
	return ctx.Expression().Accept(v)
}

func (v *AstBuilder) VisitLiteralConstant(ctx *LiteralConstantContext) interface{} {
	if l := ctx.BooleanLiteral(); l != nil {
		v := l.GetText()
		return &Boolean{
			Value: strings.ToLower(v) == "true",
		}
	}
	if l := ctx.IntegerLiteral(); l != nil {
		v, err := strconv.Atoi(l.GetText())
		if err != nil {
			panic(err)
		}
		return &Integer{
			Value: v,
		}
	}
	if l := ctx.StringLiteral(); l != nil {
		return l.Accept(v)
	}
	panic("not pass")
}

func (v *AstBuilder) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	if l := ctx.LineStringLiteral(); l != nil {
		return l.Accept(v)
	}
	if l := ctx.MultiLineStringLiteral(); l != nil {
		return l.Accept(v)
	}
	panic("not pass")
}

func (v *AstBuilder) VisitLineStringLiteral(ctx *LineStringLiteralContext) interface{} {
	parts := make([]string, len(ctx.AllLineStringPart()))
	for i, p := range ctx.AllLineStringPart() {
		parts[i] = p.GetText()
	}
	return &String{
		Value: strings.Join(parts, ""),
	}
}

func (v *AstBuilder) VisitLineStringPart(ctx *LineStringPartContext) interface{} {
	if l := ctx.LineStringContent(); l != nil {
		return l.Accept(v)
	}
	if l := ctx.LineStringExpression(); l != nil {
		return l.Accept(v)
	}
	panic("not pass")
}

func (v *AstBuilder) VisitMultiLineStringLiteral(ctx *MultiLineStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLineStringContent(ctx *LineStringContentContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitLineStringExpression(ctx *LineStringExpressionContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitMultiLineStringContent(ctx *MultiLineStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMultiLineStringExpression(ctx *MultiLineStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionLiteral(ctx *FunctionLiteralContext) interface{} {
	stmts := ctx.Statements().Accept(v).([]Node)
	var parameters []*Parameter
	if p := ctx.LambdaParameters(); p != nil {
		parameters = p.Accept(v).([]*Parameter)
	}
	return &Lambda{
		Parameters: parameters,
		Statements: stmts,
	}
}

func (v *AstBuilder) VisitLambdaParameters(ctx *LambdaParametersContext) interface{} {
	parameters := make([]*Parameter, len(ctx.AllLambdaParameter()))
	for i, p := range ctx.AllLambdaParameter() {
		parameters[i] = p.Accept(v).(*Parameter)
	}
	return parameters
}

func (v *AstBuilder) VisitLambdaParameter(ctx *LambdaParameterContext) interface{} {
	if decl := ctx.VariableDeclaration(); decl != nil {
		declaration := decl.Accept(v).(*VariableDeclaration)
		return &Parameter{
			Identifier: declaration.Identifier,
			TypeRef:    declaration.TypeRef,
		}
	}
	panic("not impl")
}

func (v *AstBuilder) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCollectionLiteral(ctx *CollectionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSuperExpression(ctx *SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	if exp := ctx.IfExpression(); exp != nil {
		return exp.Accept(v)
	}
	return ctx.WhenExpression().Accept(v)
}

func (v *AstBuilder) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	cond := ctx.Expression().Accept(v).(Node)
	ifStmt := ctx.ControlStructureBody(0).Accept(v).(*Block)
	var elseStmt *Block
	if ctx.ELSE() != nil {
		elseStmt = ctx.ControlStructureBody(1).Accept(v).(*Block)
	}
	return &If{
		Condition:     cond,
		IfStatement:   ifStmt,
		ElseStatement: elseStmt,
	}
}

func (v *AstBuilder) VisitControlStructureBody(ctx *ControlStructureBodyContext) interface{} {
	if b := ctx.Block(); b != nil {
		return b.Accept(v)
	}
	return &Block{
		Statements: []Node{
			ctx.Expression().Accept(v).(Node),
		},
	}
}

func (v *AstBuilder) VisitWhenExpression(ctx *WhenExpressionContext) interface{} {
	condition := ctx.Expression().Accept(v).(Node)
	sw := &Switch{}
	whens := []*When{}
	for _, we := range ctx.AllWhenEntry() {
		entry := we.Accept(v)
		switch e := entry.(type) {
		case *When:
			whens = append(whens, e)
		case *Block:
			sw.Else = e
		}
	}
	sw.Condition = condition
	sw.Whens = whens
	return sw
}

func (v *AstBuilder) VisitWhenEntry(ctx *WhenEntryContext) interface{} {
	if len(ctx.AllWhenCondition()) > 0 {
		conditions := make([]Node, len(ctx.AllWhenCondition()))
		for i, c := range ctx.AllWhenCondition() {
			conditions[i] = c.Accept(v).(Node)
		}
		block := ctx.ControlStructureBody().Accept(v).(*Block)
		return &When{
			Conditions: conditions,
			Block:      block,
		}
	}
	return ctx.ControlStructureBody().Accept(v)
}

func (v *AstBuilder) VisitWhenCondition(ctx *WhenConditionContext) interface{} {
	if exp := ctx.Expression(); exp != nil {
		return exp.Accept(v)
	}
	if exp := ctx.RangeTest(); exp != nil {
		return exp.Accept(v)
	}
	return ctx.TypeTest().Accept(v)
}

func (v *AstBuilder) VisitRangeTest(ctx *RangeTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeTest(ctx *TypeTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTryExpression(ctx *TryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCatchBlock(ctx *CatchBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFinallyBlock(ctx *FinallyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLoopExpression(ctx *LoopExpressionContext) interface{} {
	if f := ctx.ForExpression(); f != nil {
		return f.Accept(v)
	}
	if w := ctx.WhileExpression(); w != nil {
		return w.Accept(v)
	}
	return ctx.DoWhileExpression().Accept(v)
}

func (v *AstBuilder) VisitForExpression(ctx *ForExpressionContext) interface{} {
	expression := ctx.Expression().Accept(v)
	decl := ctx.VariableDeclaration().Accept(v)
	identifier := decl.(*VariableDeclaration).Identifier
	block := ctx.ControlStructureBody().Accept(v).(*Block)
	return &For{
		Identifier: identifier,
		Expression: expression.(Node),
		Block:      block,
	}
}

func (v *AstBuilder) VisitWhileExpression(ctx *WhileExpressionContext) interface{} {
	cond := ctx.Expression().Accept(v).(Node)
	block := ctx.ControlStructureBody().Accept(v).(*Block)
	return &While{
		Condition: cond,
		Block:     block,
	}
}

func (v *AstBuilder) VisitDoWhileExpression(ctx *DoWhileExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitJumpExpression(ctx *JumpExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCallableReference(ctx *CallableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitEqualityOperation(ctx *EqualityOperationContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitInOperator(ctx *InOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitIsOperator(ctx *IsOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAdditiveOperator(ctx *AdditiveOperatorContext) interface{} {
	if ctx.ADD() != nil {
		return ctx.ADD().GetText()
	}
	return ctx.SUB().GetText()
}

func (v *AstBuilder) VisitMultiplicativeOperation(ctx *MultiplicativeOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeOperation(ctx *TypeOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPrefixUnaryOperation(ctx *PrefixUnaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPostfixUnaryOperation(ctx *PostfixUnaryOperationContext) interface{} {
	if accessOperator := ctx.MemberAccessOperator(); accessOperator != nil {
		accessOperator.Accept(v)
		node := ctx.PostfixUnaryExpression().Accept(v)
		return node
	}
	callSuffix := ctx.CallSuffix().Accept(v).(callSuffix)
	return &MethodInvocation{
		Parameters: callSuffix.Args,
		Type: callSuffix.Type,
	}
}

func (v *AstBuilder) VisitMemberAccessOperator(ctx *MemberAccessOperatorContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitModifierList(ctx *ModifierListContext) interface{} {
	modifiers := make([]*Modifier, len(ctx.AllAnnotations())+len(ctx.AllModifier()))
	if allAnnotations := ctx.AllAnnotations(); allAnnotations != nil {
		for i, annotation := range allAnnotations {
			modifiers[i] = annotation.Accept(v).(*Modifier)
		}
	}
	if allModifiers := ctx.AllModifier(); allModifiers != nil {
		for i, modifier := range allModifiers {
			modifiers[i] = modifier.Accept(v).(*Modifier)
		}
	}
	return modifiers
}

func (v *AstBuilder) VisitModifier(ctx *ModifierContext) interface{} {
	if m := ctx.ClassModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.MemberModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.VisibilityModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.VarianceAnnotation(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.FunctionModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.PropertyModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.InheritanceModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.ParameterModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	if m := ctx.TypeParameterModifier(); m != nil {
		return &Modifier{
			Name: m.GetText(),
		}
	}
	panic("not pass")
}

func (v *AstBuilder) VisitClassModifier(ctx *ClassModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMemberModifier(ctx *MemberModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVisibilityModifier(ctx *VisibilityModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVarianceAnnotation(ctx *VarianceAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionModifier(ctx *FunctionModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPropertyModifier(ctx *PropertyModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInheritanceModifier(ctx *InheritanceModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitParameterModifier(ctx *ParameterModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeParameterModifier(ctx *TypeParameterModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLabelDefinition(ctx *LabelDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnnotations(ctx *AnnotationsContext) interface{} {
	if a := ctx.Annotation(); a != nil {
		return a.Accept(v)
	}
	return ctx.AnnotationList().Accept(v)
}

func (v *AstBuilder) VisitAnnotation(ctx *AnnotationContext) interface{} {
	// TODO: impl
	return nil
}

func (v *AstBuilder) VisitAnnotationList(ctx *AnnotationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnnotationUseSiteTarget(ctx *AnnotationUseSiteTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitUnescapedAnnotation(ctx *UnescapedAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitIdentifier(ctx *IdentifierContext) interface{} {
	identifier := make([]string, len(ctx.AllSimpleIdentifier()))
	for i, ident := range ctx.AllSimpleIdentifier() {
		identifier[i] = ident.Accept(v).(string)
	}
	return identifier
}

func (v *AstBuilder) VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{} {
	return ctx.Identifier().GetText()
}

func (v *AstBuilder) VisitSemi(ctx *SemiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnysemi(ctx *AnysemiContext) interface{} {
	return v.VisitChildren(ctx)
}

func debug(args ...interface{}) {
	pp.Println(args...)
}
