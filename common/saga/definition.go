package saga

import (
	"reflect"

	"golang.org/x/net/context"
)

type subTxDefinitions map[string]subTxDefinition

type subTxDefinition struct {
	subTxId    string
	action     reflect.Value
	compensate reflect.Value
}

func (s subTxDefinitions) addDefinition(subTxId string, action, compensate interface{}) subTxDefinitions {
	actionMethod := subTxMethod(action)
	compensateMethod := subTxMethod(compensate)
	s[subTxId] = subTxDefinition{
		subTxId:    subTxId,
		action:     actionMethod,
		compensate: compensateMethod,
	}

	return s
}

func (s subTxDefinitions) findDefinition(subTxId string) (subTxDefinition, bool) {
	define, ok := s[subTxId]
	return define, ok
}

type paramTypeRegister struct {
	nameToType map[string]reflect.Type
	typeToName map[reflect.Type]string
}

func (r *paramTypeRegister) addParams(f interface{}) {
	funcValue := subTxMethod(f)
	funcType := funcValue.Type()
	for i := 0; i < funcType.NumIn(); i++ {
		paramType := funcType.In(i)
		r.nameToType[paramType.Name()] = paramType
		r.typeToName[paramType] = paramType.Name()
	}
	for i := 0; i < funcType.NumOut(); i++ {
		returnType := funcType.Out(i)
		r.nameToType[returnType.Name()] = returnType
	}
}

func (r *paramTypeRegister) findTypeName(typ reflect.Type) (string, bool) {
	f, ok := r.typeToName[typ]
	return f, ok
}

func (r *paramTypeRegister) findType(typeName string) (reflect.Type, bool) {
	f, ok := r.nameToType[typeName]
	return f, ok
}

func subTxMethod(obj interface{}) reflect.Value {
	funcValue := reflect.ValueOf(obj)
	if funcValue.Kind() != reflect.Func {
		panic("Register object must be type of func")
	}
	if funcValue.Type().NumIn() < 1 || funcValue.Type().In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
		panic("First argument must use context.Context")
	}

	return funcValue
}
