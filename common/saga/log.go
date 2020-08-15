package saga

import (
	"encoding/json"
	"time"
)

type LogType int

const (
	SagaStart LogType = iota + 1
	SagaEnd
	SagaAbort
	ActionStart
	ActionEnd
	CompensateStart
	CompensateEnd
)

type Log struct {
	Type    LogType     `json:"type,omitempty"`
	SubTxId string      `json:"subTxId,omitempty"`
	Time    time.Time   `json:"time,omitempty"`
	Params  []ParamData `json:"params,omitempty"`
}

func (l *Log) mustMarshal() string {
	return mustMarshal(l)
}

func mustMarshal(value interface{}) string {
	s, err := json.Marshal(value)
	if err != nil {
		panic("Marshal Failure")
	}
	return string(s)
}

func mustUnmarshal(data []byte, value interface{}) {
	err := json.Unmarshal(data, value)
	if err != nil {
		panic("Unmarshal Failure")
	}
}

func mustUnmarshalLog(data string) Log {
	var log Log
	mustUnmarshal([]byte(data), &log)
	return log
}
