package parser // KotlinParser

type Node interface {
	Accept(v Visitor) (interface{}, error)
}

type File struct {
	Header  *Header
	Classes []*Class
}

type Header struct {
	Value []string
}

type Class struct {
	Modifiers []*Modifier
	Name                  string
	Properties            map[string]*Property
	Methods               map[string]*Method
	PrimaryConstructor    *Constructor
	SecondalyConstructors []*Constructor
	Declarations          []Node
}

type Property struct {
	Modifiers  []*Modifier
	TypeRef    *TypeRef
	Name       string
	Expression Node
}

type Constructor struct {
	Name string
	Parameters []*Parameter
	Block      *Block
}

type Parameter struct {
	Modifiers  []*Modifier
	TypeRef    *TypeRef
	Identifier string
	Expression Node
}

type Method struct {
	AccessModifiers []*Modifier
	Identifier []string
	Parameters []*Parameter
	Statements      *Block
}

type Modifier struct {
	Name string
}

type If struct {
	Condition     Node
	IfStatement   *Block
	ElseStatement *Block
}

type For struct {
	Expression Node
	Identifier string
	Block      *Block
}

type While struct {
	Condition Node
	Block     *Block
}

type Switch struct {
	Condition Node
	Whens     []When
	Else      *Block
}

type When struct {
	Expression Node
	Block      *Block
}

type VariableDeclaration struct {
	Identifier string
	TypeRef    *TypeRef
}

var publicModifier = &Modifier{"public"}
var privateModifier = &Modifier{"private"}
var protectedModifier = &Modifier{"protected"}
var globalModifier = &Modifier{"global"}

type Block struct {
	Statements []Node
}

type TypeRef struct {
	Name []string
}

type Boolean struct {
	Value bool
}

type Integer struct {
	Value int
}

type String struct {
	Value string
}

type Identifier struct {
	Value string
}

type MemberAccess struct {
	Left Node
	Right Node
}

type BinaryOperator struct {
	Operator string
	Left Node
	Right Node
}

type MethodInvocation struct {
	Expression Node
	Parameters []Node
}

type Name struct {
	Value []string
}

type Visitor interface {
	VisitFile(*File) (interface{}, error)
	VisitHeader(*Header) (interface{}, error)
	VisitClass(*Class) (interface{}, error)
	VisitProperty(*Property) (interface{}, error)
	VisitConstructor(*Constructor) (interface{}, error)
	VisitParameter(*Parameter) (interface{}, error)
	VisitMethod(*Method) (interface{}, error)
	VisitModifier(*Modifier) (interface{}, error)
	VisitIf(*If) (interface{}, error)
	VisitFor(*For) (interface{}, error)
	VisitSwitch(*Switch) (interface{}, error)
	VisitWhen(*When) (interface{}, error)
	VisitVariableDeclaration(*VariableDeclaration) (interface{}, error)
	VisitBlock(*Block) (interface{}, error)
	VisitTypeRef(*TypeRef) (interface{}, error)
	VisitBoolean(*Boolean) (interface{}, error)
	VisitInteger(*Integer) (interface{}, error)
	VisitString(*String) (interface{}, error)
	VisitIdentifier(*Identifier) (interface{}, error)
	VisitMemberAccess(*MemberAccess) (interface{}, error)
	VisitBinaryOperator(*BinaryOperator) (interface{}, error)
	VisitName(*Name) (interface{}, error)
	VisitMethodInvocation(*MethodInvocation) (interface{}, error)
}
func (t *File) Accept(v Visitor) (interface{}, error) {
	return v.VisitFile(t)
}

func (t *Header) Accept(v Visitor) (interface{}, error) {
	return v.VisitHeader(t)
}

func (t *Class) Accept(v Visitor) (interface{}, error) {
	return v.VisitClass(t)
}

func (t *Property) Accept(v Visitor) (interface{}, error) {
	return v.VisitProperty(t)
}

func (t *Constructor) Accept(v Visitor) (interface{}, error) {
	return v.VisitConstructor(t)
}

func (t *Parameter) Accept(v Visitor) (interface{}, error) {
	return v.VisitParameter(t)
}

func (t *Method) Accept(v Visitor) (interface{}, error) {
	return v.VisitMethod(t)
}

func (t *Modifier) Accept(v Visitor) (interface{}, error) {
	return v.VisitModifier(t)
}

func (t *If) Accept(v Visitor) (interface{}, error) {
	return v.VisitIf(t)
}

func (t *For) Accept(v Visitor) (interface{}, error) {
	return v.VisitFor(t)
}

func (t *Switch) Accept(v Visitor) (interface{}, error) {
	return v.VisitSwitch(t)
}

func (t *When) Accept(v Visitor) (interface{}, error) {
	return v.VisitWhen(t)
}

func (t *VariableDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitVariableDeclaration(t)
}

func (t *Block) Accept(v Visitor) (interface{}, error) {
	return v.VisitBlock(t)
}

func (t *TypeRef) Accept(v Visitor) (interface{}, error) {
	return v.VisitTypeRef(t)
}

func (t *Boolean) Accept(v Visitor) (interface{}, error) {
	return v.VisitBoolean(t)
}

func (t *Integer) Accept(v Visitor) (interface{}, error) {
	return v.VisitInteger(t)
}

func (t *String) Accept(v Visitor) (interface{}, error) {
	return v.VisitString(t)
}

func (t *Identifier) Accept(v Visitor) (interface{}, error) {
	return v.VisitIdentifier(t)
}

func (t *MemberAccess) Accept(v Visitor) (interface{}, error) {
	return v.VisitMemberAccess(t)
}

func (t *BinaryOperator) Accept(v Visitor) (interface{}, error) {
	return v.VisitBinaryOperator(t)
}

func (t *Name) Accept(v Visitor) (interface{}, error) {
	return v.VisitName(t)
}

func (t *MethodInvocation) Accept(v Visitor) (interface{}, error) {
	return v.VisitMethodInvocation(t)
}
