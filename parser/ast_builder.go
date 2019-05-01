package parser // KotlinParser

// A complete Visitor for a parse tree produced by KotlinParser.
type AstBuilder struct {
	*BaseKotlinParserVisitor
}

type File struct {
	Header *Header
	Classes []*Class
}

type Header struct {
	Value string
}

type Class struct {
	Name string
}

func (v *AstBuilder) VisitKotlinFile(ctx *KotlinFileContext) interface{} {
	preamble := ctx.Preamble().Accept(v)
	klass := ctx.AllTopLevelObject()[0].Accept(v)
	return &File{
		Header: preamble.(*Header),
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
		Value: ctx.Identifier().Accept(v).(string),
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
	return &Class{
		Name: ctx.SimpleIdentifier().Accept(v).(string),
	}
}

func (v *AstBuilder) VisitPrimaryConstructor(ctx *PrimaryConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassParameters(ctx *ClassParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassParameter(ctx *ClassParameterContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnonymousInitializer(ctx *AnonymousInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSecondaryConstructor(ctx *SecondaryConstructorContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionValueParameters(ctx *FunctionValueParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionValueParameter(ctx *FunctionValueParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitObjectDeclaration(ctx *ObjectDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCompanionObject(ctx *CompanionObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionTypeReceiver(ctx *FunctionTypeReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitUserType(ctx *UserTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSimpleUserType(ctx *SimpleUserTypeContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitBlockLevelExpression(ctx *BlockLevelExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitDisjunction(ctx *DisjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConjunction(ctx *ConjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEqualityComparison(ctx *EqualityComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitNamedInfix(ctx *NamedInfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitElvisExpression(ctx *ElvisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInfixFunctionCall(ctx *InfixFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitRangeExpression(ctx *RangeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeRHS(ctx *TypeRHSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPrefixUnaryExpression(ctx *PrefixUnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPostfixUnaryExpression(ctx *PostfixUnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAtomicExpression(ctx *AtomicExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCallSuffix(ctx *CallSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnnotatedLambda(ctx *AnnotatedLambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitValueArguments(ctx *ValueArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeProjection(ctx *TypeProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeProjectionModifierList(ctx *TypeProjectionModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitValueArgument(ctx *ValueArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLiteralConstant(ctx *LiteralConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLineStringLiteral(ctx *LineStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMultiLineStringLiteral(ctx *MultiLineStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLineStringContent(ctx *LineStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLineStringExpression(ctx *LineStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMultiLineStringContent(ctx *MultiLineStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMultiLineStringExpression(ctx *MultiLineStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFunctionLiteral(ctx *FunctionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLambdaParameters(ctx *LambdaParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLambdaParameter(ctx *LambdaParameterContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitControlStructureBody(ctx *ControlStructureBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhenExpression(ctx *WhenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhenEntry(ctx *WhenEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhenCondition(ctx *WhenConditionContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitForExpression(ctx *ForExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhileExpression(ctx *WhileExpressionContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEqualityOperation(ctx *EqualityOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInOperator(ctx *InOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitIsOperator(ctx *IsOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAdditiveOperator(ctx *AdditiveOperatorContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMemberAccessOperator(ctx *MemberAccessOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitModifierList(ctx *ModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
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
	return ctx.SimpleIdentifier(0).Accept(v)
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
