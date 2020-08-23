package subtx

import (
	"github.com/mapserver2007/golang-example-app/common/saga"
	users "github.com/mapserver2007/golang-example-app/grpc-service1-server/services/saga"
	items "github.com/mapserver2007/golang-example-app/grpc-service2-server/services/saga"
)

var SubTxDefinitions = NewSubTxDefinitionsFactory()

func NewSubTxDefinitionsFactory() *saga.ExecutionCoodinator {
	return saga.AddSubTxDef("createUser", users.CreateUserAction, users.CreateUserCompensate).
		AddSubTxDef("createItem", items.CreateItemAction, items.CreateItemCompensate)
}
