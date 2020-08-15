package saga

import (
	"reflect"
	"strconv"

	"golang.org/x/net/context"
)

var DefaultSagaExecutionCoordinator = NewSagaExecutionCoordinator()

type ExecutionCoodinator struct {
	subTxDefinitions  subTxDefinitions
	paramTypeRegister *paramTypeRegister
}

func NewSagaExecutionCoordinator() ExecutionCoodinator {
	return ExecutionCoodinator{
		subTxDefinitions: make(subTxDefinitions),
		paramTypeRegister: &paramTypeRegister{
			nameToType: make(map[string]reflect.Type),
			typeToName: make(map[reflect.Type]string),
		},
	}
}

func AddSubTxDef(subTxId string, action, compensate interface{}) *ExecutionCoodinator {
	return DefaultSagaExecutionCoordinator.AddSubTxDef(subTxId, action, compensate)
}

func (e *ExecutionCoodinator) AddSubTxDef(subTxId string, action, compensate interface{}) *ExecutionCoodinator {
	e.paramTypeRegister.addParams(action)
	e.paramTypeRegister.addParams(compensate)
	e.subTxDefinitions.addDefinition(subTxId, action, compensate)
	return e
}

func (e *ExecutionCoodinator) InitSaga(ctx context.Context, id uint64) *Saga {
	return &Saga{
		id:      id,
		context: ctx,
		sec:     e,
		logId:   LogPrefix + strconv.FormatInt(int64(id), 10),
	}
}

func (e *ExecutionCoodinator) MustFindParamName(typ reflect.Type) string {
	name, ok := e.paramTypeRegister.findTypeName(typ)
	if !ok {
		panic("Can not find param name: " + typ.String())
	}
	return name
}

func (e *ExecutionCoodinator) MustFindParamType(name string) reflect.Type {
	typ, ok := e.paramTypeRegister.findType(name)
	if !ok {
		panic("Can not find param type: " + name)
	}
	return typ
}

func (e *ExecutionCoodinator) MustFindSubTxDef(subTxId string) subTxDefinition {
	define, ok := e.subTxDefinitions.findDefinition(subTxId)
	if !ok {
		panic("SubTxId not found in context: " + subTxId)
	}
	return define
}
