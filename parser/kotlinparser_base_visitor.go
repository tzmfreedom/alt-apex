// Code generated from KotlinParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // KotlinParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseKotlinParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKotlinParserVisitor) VisitKotlinFile(ctx *KotlinFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPreamble(ctx *PreambleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFileAnnotations(ctx *FileAnnotationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFileAnnotation(ctx *FileAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPackageHeader(ctx *PackageHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitImportList(ctx *ImportListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitImportHeader(ctx *ImportHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitImportAlias(ctx *ImportAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTopLevelObject(ctx *TopLevelObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPrimaryConstructor(ctx *PrimaryConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitClassParameters(ctx *ClassParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitClassParameter(ctx *ClassParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitDelegationSpecifiers(ctx *DelegationSpecifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitDelegationSpecifier(ctx *DelegationSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitConstructorInvocation(ctx *ConstructorInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitExplicitDelegation(ctx *ExplicitDelegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnonymousInitializer(ctx *AnonymousInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitSecondaryConstructor(ctx *SecondaryConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitConstructorDelegationCall(ctx *ConstructorDelegationCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitEnumClassBody(ctx *EnumClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitEnumEntries(ctx *EnumEntriesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitEnumEntry(ctx *EnumEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionValueParameters(ctx *FunctionValueParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionValueParameter(ctx *FunctionValueParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitObjectDeclaration(ctx *ObjectDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitCompanionObject(ctx *CompanionObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitGetter(ctx *GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitSetter(ctx *SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeAlias(ctx *TypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitKtype(ctx *KtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeModifierList(ctx *TypeModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitParenthesizedType(ctx *ParenthesizedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitNullableType(ctx *NullableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionTypeReceiver(ctx *FunctionTypeReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitUserType(ctx *UserTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitSimpleUserType(ctx *SimpleUserTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionTypeParameters(ctx *FunctionTypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeConstraints(ctx *TypeConstraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeConstraint(ctx *TypeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitBlockLevelExpression(ctx *BlockLevelExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitDisjunction(ctx *DisjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitConjunction(ctx *ConjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitEqualityComparison(ctx *EqualityComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitNamedInfix(ctx *NamedInfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitElvisExpression(ctx *ElvisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitInfixFunctionCall(ctx *InfixFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitRangeExpression(ctx *RangeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeRHS(ctx *TypeRHSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPrefixUnaryExpression(ctx *PrefixUnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPostfixUnaryExpression(ctx *PostfixUnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAtomicExpression(ctx *AtomicExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitCallSuffix(ctx *CallSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnnotatedLambda(ctx *AnnotatedLambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitValueArguments(ctx *ValueArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeProjection(ctx *TypeProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeProjectionModifierList(ctx *TypeProjectionModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitValueArgument(ctx *ValueArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLiteralConstant(ctx *LiteralConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLineStringLiteral(ctx *LineStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLineStringPart(ctx *LineStringPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMultiLineStringLiteral(ctx *MultiLineStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLineStringContent(ctx *LineStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLineStringExpression(ctx *LineStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMultiLineStringContent(ctx *MultiLineStringContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMultiLineStringExpression(ctx *MultiLineStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionLiteral(ctx *FunctionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLambdaParameters(ctx *LambdaParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLambdaParameter(ctx *LambdaParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitCollectionLiteral(ctx *CollectionLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitSuperExpression(ctx *SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitControlStructureBody(ctx *ControlStructureBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitWhenExpression(ctx *WhenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitWhenEntry(ctx *WhenEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitWhenCondition(ctx *WhenConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitRangeTest(ctx *RangeTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeTest(ctx *TypeTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTryExpression(ctx *TryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitCatchBlock(ctx *CatchBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFinallyBlock(ctx *FinallyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLoopExpression(ctx *LoopExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitForExpression(ctx *ForExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitWhileExpression(ctx *WhileExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitDoWhileExpression(ctx *DoWhileExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitJumpExpression(ctx *JumpExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitCallableReference(ctx *CallableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitEqualityOperation(ctx *EqualityOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitInOperator(ctx *InOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitIsOperator(ctx *IsOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAdditiveOperator(ctx *AdditiveOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMultiplicativeOperation(ctx *MultiplicativeOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeOperation(ctx *TypeOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPrefixUnaryOperation(ctx *PrefixUnaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPostfixUnaryOperation(ctx *PostfixUnaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMemberAccessOperator(ctx *MemberAccessOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitModifierList(ctx *ModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitClassModifier(ctx *ClassModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitMemberModifier(ctx *MemberModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitVisibilityModifier(ctx *VisibilityModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitVarianceAnnotation(ctx *VarianceAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitFunctionModifier(ctx *FunctionModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitPropertyModifier(ctx *PropertyModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitInheritanceModifier(ctx *InheritanceModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitParameterModifier(ctx *ParameterModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitTypeParameterModifier(ctx *TypeParameterModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitLabelDefinition(ctx *LabelDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnnotations(ctx *AnnotationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnnotationList(ctx *AnnotationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnnotationUseSiteTarget(ctx *AnnotationUseSiteTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitUnescapedAnnotation(ctx *UnescapedAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitSemi(ctx *SemiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKotlinParserVisitor) VisitAnysemi(ctx *AnysemiContext) interface{} {
	return v.VisitChildren(ctx)
}
