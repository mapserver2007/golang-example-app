// This package is a sample application of the Saga pattern of distributed transactions.
// nolint: gochecknoinits
package saga

import (
	"reflect"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/gorp.v1"

	appLog "github.com/mapserver2007/golang-example-app/common/log"
	"github.com/mapserver2007/golang-example-app/common/saga/storage"
)

const LogPrefix = "saga_"

var StorageProvider storage.StorageProvider
var StorageConfig storage.StorageConfig

func LogStorage() storage.Storage {
	return StorageProvider(StorageConfig)
}

type Saga struct {
	logId    string
	serverId string
	context  context.Context
	conn     *gorp.DbMap
	sec      *ExecutionCoodinator
}

func (s *Saga) StartSaga() *Saga {
	log := &Log{
		Type:     SagaStart,
		ServerId: s.serverId,
		Time:     time.Now(),
	}
	err := LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga log: " + log.mustMarshal())

	return s
}

func (s *Saga) ExecSub(subTxId string, args ...interface{}) *Saga {
	subTxDef := s.sec.MustFindSubTxDef(subTxId)
	log := &Log{
		Type:     ActionStart,
		ServerId: s.serverId,
		SubTxId:  subTxId,
		Time:     time.Now(),
		Params:   MarshalParam(s.sec, subTxId, args),
	}

	// Create parameters for executing the reflection method
	params := make([]reflect.Value, 0, len(args)+1)
	params = append(params, reflect.ValueOf(s.context), reflect.ValueOf(s.conn))
	for _, arg := range args {
		params = append(params, reflect.ValueOf(arg))
	}
	result := subTxDef.action.Call(params)
	log.ResultParams = result[0].Interface().([]int64)

	err := LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga SubTxId: " + subTxId)
	appLog.Info("saga log: " + log.mustMarshal())

	if len(result) == 2 && !result[1].IsNil() {
		s.Abort()
		return s
	}

	log = &Log{
		Type:     ActionEnd,
		ServerId: s.serverId,
		SubTxId:  subTxId,
		Time:     time.Now(),
	}
	err = LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga SubTxId: " + subTxId)
	appLog.Info("saga log: " + log.mustMarshal())

	return s
}

func (s *Saga) EndSaga() {
	log := &Log{
		Type:     SagaEnd,
		ServerId: s.serverId,
		Time:     time.Now(),
	}
	err := LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga log: " + log.mustMarshal())

	err = LogStorage().Cleanup(s.logId)
	if err != nil {
		panic("Cleanup log failure")
	}

	appLog.Info("Cleanup logId: " + s.logId)
}

// When "Abort" is called, the log is traced backwards and a compensation transaction rollback is started.
func (s *Saga) Abort() {
	logs, err := LogStorage().Lookup(s.logId)

	if err != nil {
		panic("Abort panic")
	}
	abortLog := &Log{
		Type:     SagaAbort,
		ServerId: s.serverId,
		Time:     time.Now(),
	}
	err = LogStorage().AppendLog(s.logId, abortLog.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga log: " + abortLog.mustMarshal())

	for j := len(logs) - 1; j >= 0; j-- {
		logStr := logs[j]
		log := mustUnmarshalLog(logStr)
		if log.Type == ActionStart {
			s.compensate(&log)
		}
	}
}

func (s *Saga) compensate(txLog *Log) {
	cLog := &Log{
		Type:     CompensateStart,
		ServerId: s.serverId,
		SubTxId:  txLog.SubTxId,
		Time:     time.Now(),
	}
	err := LogStorage().AppendLog(s.logId, cLog.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga SubTxId: " + txLog.SubTxId)
	appLog.Info("saga log: " + cLog.mustMarshal())

	refArgs := []reflect.Value{reflect.ValueOf(s.context), reflect.ValueOf(s.conn)}
	args := UnmarshalParam(s.sec, txLog.SubTxId, txLog.Params)
	lastInsertValues := reflect.ValueOf(txLog.ResultParams)
	params := make([]reflect.Value, 0, len(refArgs)+len(args)+1)
	params = append(params, refArgs...)
	params = append(params, lastInsertValues)
	params = append(params, args...)

	subDef := s.sec.MustFindSubTxDef(txLog.SubTxId)
	result := subDef.compensate.Call(params)

	if len(result) == 1 && !result[0].IsNil() {
		s.Abort()
		return
	}

	cLog = &Log{
		Type:     CompensateEnd,
		ServerId: s.serverId,
		SubTxId:  txLog.SubTxId,
		Time:     time.Now(),
	}
	err = LogStorage().AppendLog(s.logId, cLog.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	appLog.Info("saga logId: " + s.logId)
	appLog.Info("saga SubTxId: " + txLog.SubTxId)
	appLog.Info("saga log: " + cLog.mustMarshal())
}
