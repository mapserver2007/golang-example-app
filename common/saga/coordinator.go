package saga

import (
	"reflect"

	"golang.org/x/net/context"
	"gopkg.in/gorp.v1"
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
			nameToType: make(map[string]map[string]reflect.Type),
			typeToName: make(map[string]map[reflect.Type]string),
		},
	}
}

func CreateSubTx(ctx context.Context, conn *gorp.DbMap, serverId, id string) *Saga {
	return DefaultSagaExecutionCoordinator.CreateSubTx(ctx, conn, serverId, id)
}

func AddSubTxDef(subTxId string, action, compensate interface{}) *ExecutionCoodinator {
	return DefaultSagaExecutionCoordinator.AddSubTxDef(subTxId, action, compensate)
}

func (e *ExecutionCoodinator) AddSubTxDef(subTxId string, action, compensate interface{}) *ExecutionCoodinator {
	e.paramTypeRegister.addParams(subTxId, action)
	e.paramTypeRegister.addParams(subTxId, compensate)
	e.subTxDefinitions.addDefinition(subTxId, action, compensate)

	return e
}

func (e *ExecutionCoodinator) CreateSubTx(ctx context.Context, conn *gorp.DbMap, serverId, id string) *Saga {
	return &Saga{
		logId:    LogPrefix + id,
		serverId: serverId,
		context:  ctx,
		conn:     conn,
		sec:      e,
	}
}

func (e *ExecutionCoodinator) MustFindParamName(subTxId string, typ reflect.Type) string {
	name, ok := e.paramTypeRegister.findTypeName(subTxId, typ)
	if !ok {
		panic("Can not find param name: " + typ.String())
	}
	return name
}

func (e *ExecutionCoodinator) MustFindParamType(subTxId, name string) reflect.Type {
	typ, ok := e.paramTypeRegister.findType(subTxId, name)
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
