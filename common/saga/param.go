package saga

import "reflect"

type ParamData struct {
	ParamType string `json:"paramType,omitempty"`
	Data      string `json:"data,omitempty"`
}

func MarshalParam(sec *ExecutionCoodinator, args []interface{}) []ParamData {
	p := make([]ParamData, 0, len(args))
	for _, arg := range args {
		typ := sec.MustFindParamName(reflect.ValueOf(arg).Type())
		p = append(p, ParamData{
			ParamType: typ,
			Data:      mustMarshal(arg),
		})
	}
	return p
}

func UnmarshalParam(sec *ExecutionCoodinator, params []ParamData) []reflect.Value {
	var values []reflect.Value
	for _, param := range params {
		ptyp := sec.MustFindParamType(param.ParamType)
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
