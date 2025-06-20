package object

import (
	"math/big"
	"strconv"
)

type ObjKind int

func (o ObjKind) Equal(x ObjKind) bool {
	return o == x
}

const (
	O_NUM ObjKind = iota
	O_STR
	O_BOOL
	O_FUNC
	O_NULL
	O_BUILTIN
	O_LITERAL
	O_ARR
	O_MAP
	O_RET
	O_QUOTE
	O_TRAIT
	O_ENUM
	O_CLASS
	O_OBJ
)

type Type interface {
	Name() string

	Text() string

	New(values ...Value) Value

	// Implements reports whether the type implements the interface type u.
	Implements(u Type) bool

	AssignableTo(u Type) bool

	ConvertibleTo(u Type) bool

	Kind() ObjKind

	NumMethod() int

	Method(i int) Method

	MethodByName(name string) Method

	NumField() int

	Field(i int) Type

	FieldByName(name string) Type
}

type Value interface {
	Type() Type

	Text() string

	Interface() any
}

type Function interface {
	Signature() []Type
	Call(args ...Value) Value
	Match(args []Value) bool

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn() int

	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(i int) Type
}

type Method interface {
	Type
	Function
}

type ObjectType struct {
	name    string
	kind    ObjKind
	pkgPath string
	methods map[string]Method
	fields  map[string]Type
}

func NewObjectType(name string, kind ObjKind, pkgPath string) *ObjectType {
	return &ObjectType{
		name:    name,
		kind:    kind,
		pkgPath: pkgPath,
		methods: make(map[string]Method),
		fields:  make(map[string]Type),
	}
}

func (t *ObjectType) Name() string    { return t.name }
func (t *ObjectType) Kind() ObjKind   { return t.kind }
func (t *ObjectType) PkgPath() string { return t.pkgPath }

func (t *ObjectType) NumMethod() int {
	return len(t.methods)
}

func (t *ObjectType) Method(i int) Method {
	if i < 0 || i >= len(t.methods) {
		return nil
	}
	methods := make([]Method, 0, len(t.methods))
	for _, method := range t.methods {
		methods = append(methods, method)
	}
	return methods[i]
}

func (t *ObjectType) MethodByName(name string) Method {
	return t.methods[name]
}

func (t *ObjectType) NumField() int {
	return len(t.fields)
}

func (t *ObjectType) Field(i int) Type {
	if i < 0 || i >= len(t.fields) {
		return nil
	}
	fields := make([]Type, 0, len(t.fields))
	for _, method := range t.fields {
		fields = append(fields, method)
	}
	return fields[i]
}

func (t *ObjectType) FieldByName(name string) Type {
	return t.fields[name]
}

func (t *ObjectType) Implements(u Type) bool {
	return false
}

func (t *ObjectType) AssignableTo(u Type) bool {
	return false
}

func (t *ObjectType) ConvertibleTo(u Type) bool {
	return false
}

func (t *ObjectType) New(values ...Value) Value {
	return nil
}

func (t *ObjectType) Text() string {
	return t.name
}

type NullValue struct{}

func (n *NullValue) Text() string {
	return "null"
}

func (n *NullValue) Interface() any {
	return nil
}

func (n *NullValue) Type() Type {
	return nil
}

var NULL = &NullValue{}

type NumberValue struct {
	// TODO 根据精度选择存储模型
	Value *big.Float
}

func (n *NumberValue) Text() string {
	return n.Value.String()
}

func (n *NumberValue) Interface() any {
	return n.Value
}

func (n *NumberValue) Type() Type {
	return NumberType
}

var NumberType = NewObjectType("number", O_NUM, "std")

type String struct {
	Value string
}

func (s *String) Text() string {
	return s.Value
}

func (s *String) Interface() any {
	return s.Value
}

func (s *String) Type() Type {
	return StringType
}

var StringType = NewObjectType("string", O_STR, "std")

type Boolean struct {
	Value bool
}

func (b *Boolean) Text() string {
	return strconv.FormatBool(b.Value)
}

func (b *Boolean) Interface() any {
	return b.Value
}

func (b *Boolean) Type() Type {
	return BooleanType
}

var BooleanType = NewObjectType("bool", O_BOOL, "std")

var TRUE = &Boolean{Value: true}
var FALSE = &Boolean{Value: false}

type ErrorValue struct {
	Value string
}

func (e *ErrorValue) Text() string {
	return e.Value
}

func (e *ErrorValue) Interface() any {
	return e.Value
}

func (e *ErrorValue) Type() Type {
	return ErrorType
}

var ErrorType = NewObjectType("error", O_OBJ, "std")

type OverloadedFunction struct {
	name  string
	funcs []Function
}

func (o *OverloadedFunction) Text() string {
	return "overloaded function"
}

func (o *OverloadedFunction) Interface() any {
	return o.funcs
}

func (o *OverloadedFunction) Type() Type {
	return FunctionType
}

var FunctionType = NewObjectType("function", O_FUNC, "std")

type BuiltinFunction func(args ...Value) Value

func (b *BuiltinFunction) Text() string {
	return "builtin function"
}

func (b *BuiltinFunction) Interface() any {
	return b
}

func (b *BuiltinFunction) Type() Type {
	return FunctionType
}

type UserFunction struct {
}

type OperatorFunction struct{}

type Operator int

const (
	OpIllegal      Operator = iota
	OpAdd                   // +
	OpSub                   // -
	OpMul                   // *
	OpDiv                   // /
	OpMod                   // %
	OpConcat                // ..
	OpAnd                   // &&
	OpOr                    // ||
	OpEqual                 // ==
	OpNot                   // !
	OpLess                  // <
	OpLessEqual             // <=
	OpGreater               // >
	OpGreaterEqual          // >=
	OpAssign                // =
	OpIndex                 // []
	OpCall                  // ()

	Opa // a""
	Opb // b""
	Opc // c""
	Opd // d""
	Ope // e""
	Opf // f""
	Opg // g""
	Oph // h""
	Opi // i""
	Opj // j""
	Opk // k""
	Opl // l""
	Opm // m""
	Opn // n""
	Opo // o""
	Opp // p""
	Opq // q""
	Opr // r""
	Ops // s""
	Ott // t""
	Opu // u""
	Opv // v""
	Opw // w""
	Opz // z""

	OpA // A""
	OpB // B""
	OpC // C""
	OpD // D""
	OpE // E""
	OpF // F""
	OpG // G""
	OpH // H""
	OpI // I""
	OpJ // J""
	OpK // K""
	OpL // L""
	OpM // M""
	OpN // N""
	OpO // O""
	OpP // P""
	OpQ // Q""
	OpR // R""
	OpS // S""
	OpT // T""
	OpU // U""
	OpV // V""
	OpW // W""
	OpX // X""
	OpY // Y""
	OpZ // Z""
)
