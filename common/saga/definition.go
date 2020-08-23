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
	nameToType map[string]map[string]reflect.Type
	typeToName map[string]map[reflect.Type]string
}

func (r *paramTypeRegister) addParams(subTxId string, f interface{}) {
	var ok bool
	if _, ok = r.nameToType[subTxId]; !ok {
		r.nameToType[subTxId] = make(map[string]reflect.Type)
	}
	if _, ok = r.typeToName[subTxId]; !ok {
		r.typeToName[subTxId] = make(map[reflect.Type]string)
	}

	funcValue := subTxMethod(f)
	funcType := funcValue.Type()
	for i := 0; i < funcType.NumIn(); i++ {
		paramType := funcType.In(i)
		r.nameToType[subTxId][paramType.Name()] = paramType
		r.typeToName[subTxId][paramType] = paramType.Name()
	}
	for i := 0; i < funcType.NumOut(); i++ {
		returnType := funcType.Out(i)
		r.nameToType[subTxId][returnType.Name()] = returnType
	}
}

func (r *paramTypeRegister) findTypeName(subTxId string, typ reflect.Type) (string, bool) {
	f, ok := r.typeToName[subTxId][typ]
	return f, ok
}

func (r *paramTypeRegister) findType(subTxId, typeName string) (reflect.Type, bool) {
	f, ok := r.nameToType[subTxId][typeName]
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
