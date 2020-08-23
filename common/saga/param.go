package saga

import (
	"reflect"
)

type ParamData struct {
	ParamType string `json:"paramType,omitempty"`
	Data      string `json:"data,omitempty"`
}

func MarshalParam(sec *ExecutionCoodinator, subTxId string, args []interface{}) []ParamData {
	p := make([]ParamData, 0, len(args))
	for _, arg := range args {
		typ := sec.MustFindParamName(subTxId, reflect.ValueOf(arg).Type())
		p = append(p, ParamData{
			ParamType: typ,
			Data:      mustMarshal(arg),
		})
	}
	return p
}

func MarshalResultParam(args []interface{}) []ParamData {
	p := make([]ParamData, 0, len(args))
	for _, arg := range args {
		p = append(p, ParamData{
			ParamType: reflect.TypeOf(arg).Name(),
			Data:      mustMarshal(arg),
		})
	}
	return p
}

func UnmarshalParam(sec *ExecutionCoodinator, subTxId string, params []ParamData) []reflect.Value {
	var values []reflect.Value
	for _, param := range params {
		ptyp := sec.MustFindParamType(subTxId, param.ParamType)
		obj := reflect.New(ptyp).Interface()
		mustUnmarshal([]byte(param.Data), obj)
		objValue := reflect.ValueOf(obj)
		if objValue.Type().Kind() == reflect.Ptr && objValue.Type() != ptyp {
			objValue = objValue.Elem()
		}
		values = append(values, objValue)
	}
	return values
}
