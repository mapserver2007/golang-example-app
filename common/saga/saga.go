// This package is a sample application of the Saga pattern of distributed transactions.
// nolint: gochecknoinits
package saga

import (
	"reflect"
	"time"

	"golang.org/x/net/context"

	"github.com/mapserver2007/golang-example-app/common/saga/storage"
)

const LogPrefix = "saga_"

var StorageProvider storage.StorageProvider
var StorageConfig storage.StorageConfig

func init() {
	// TODO logとか
}

func LogStorage() storage.Storage {
	return StorageProvider(StorageConfig)
}

type Saga struct {
	id      uint64
	logId   string
	context context.Context
	sec     *ExecutionCoodinator
}

func (s *Saga) StartSaga() *Saga {
	log := &Log{
		Type: SagaStart,
		Time: time.Now(),
	}
	err := LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}
	return s
}

func (s *Saga) ExecSub(subTxId string, args ...interface{}) *Saga {
	subTxDef := s.sec.MustFindSubTxDef(subTxId)
	log := &Log{
		Type:    ActionStart,
		SubTxId: subTxId,
		Time:    time.Now(),
		Params:  MarshalParam(s.sec, args),
	}
	err := LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	// Create parameters for executing the reflection method
	params := make([]reflect.Value, 0, len(args)+1)
	params = append(params, reflect.ValueOf(s.context))
	for _, arg := range args {
		params = append(params, reflect.ValueOf(arg))
	}
	result := subTxDef.action.Call(params)

	if len(result) == 1 && !result[0].IsNil() {
		s.Abort()
		return s
	}

	log = &Log{
		Type:    ActionEnd,
		SubTxId: subTxId,
		Time:    time.Now(),
	}
	err = LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	return s
}

func (s *Saga) EndSaga() {
	log := &Log{
		Type: SagaEnd,
		Time: time.Now(),
	}
	err := LogStorage().AppendLog(s.logId, log.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}
	err = LogStorage().Cleanup(s.logId)
	if err != nil {
		panic("Clean up topic failure")
	}
}

// When "Abort" is called, the log is traced backwards and a compensation transaction rollback is started.
func (s *Saga) Abort() {
	logs, err := LogStorage().Lookup(s.logId)

	if err != nil {
		panic("Abort panic")
	}
	abortLog := &Log{
		Type: SagaAbort,
		Time: time.Now(),
	}
	err = LogStorage().AppendLog(s.logId, abortLog.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}
	for i := len(logs) - 1; i >= 0; i-- {
		logStr := logs[i]
		log := mustUnmarshalLog(logStr)
		if log.Type == ActionStart {
			s.compensate(log)
		}
	}
}

func (s *Saga) compensate(txLog Log) {
	cLog := &Log{
		Type:    CompensateStart,
		SubTxId: txLog.SubTxId,
		Time:    time.Now(),
	}
	err := LogStorage().AppendLog(s.logId, cLog.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}

	args := UnmarshalParam(s.sec, txLog.Params)
	params := make([]reflect.Value, 0, len(args)+1)
	params = append(params, reflect.ValueOf(s.context))
	params = append(params, args...)

	subDef := s.sec.MustFindSubTxDef(txLog.SubTxId)
	result := subDef.compensate.Call(params)

	if len(result) == 1 && !result[0].IsNil() {
		s.Abort()
		return
	}

	cLog = &Log{
		Type:    CompensateEnd,
		SubTxId: txLog.SubTxId,
		Time:    time.Now(),
	}
	err = LogStorage().AppendLog(s.logId, cLog.mustMarshal())
	if err != nil {
		panic("Add log failure")
	}
}
