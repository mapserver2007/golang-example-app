// nolint: gochecknoinits
package memory

import (
	"errors"
	"sync"

	"github.com/mapserver2007/golang-example-app/common/saga"
	"github.com/mapserver2007/golang-example-app/common/saga/storage"
)

var storageInstance storage.Storage
var memoryInit sync.Once

func init() {
	saga.StorageProvider = func(cfg storage.StorageConfig) storage.Storage {
		memoryInit.Do(func() {
			storageInstance = newMemoryStorage()
		})

		return storageInstance
	}
}

type memoryStorage struct {
	data map[string][]string
}

func newMemoryStorage() storage.Storage {
	return &memoryStorage{
		data: make(map[string][]string),
	}
}

func (s *memoryStorage) AppendLog(logId, data string) error {
	_, ok := s.data[logId]
	if !ok {
		s.data[logId] = []string{}
	}
	s.data[logId] = append(s.data[logId], data)
	return nil
}

func (s *memoryStorage) Lookup(logId string) ([]string, error) {
	return s.data[logId], nil
}

func (s *memoryStorage) Close() error {
	return nil
}

func (s *memoryStorage) LogIds() ([]string, error) {
	ids := make([]string, 0, len(s.data))
	for id := range s.data {
		ids = append(ids, id)
	}
	return ids, nil
}

func (s *memoryStorage) Cleanup(logId string) error {
	delete(s.data, logId)
	return nil
}

func (s *memoryStorage) LastLog(logId string) (string, error) {
	logDataList, ok := s.data[logId]
	if !ok {
		return "", errors.New("LogData is not found: " + logId)
	}
	logSize := len(logDataList)
	if logSize == 0 {
		return "", errors.New("LogData is empty: " + logId)
	}
	lastLog := logDataList[logSize-1]

	return lastLog, nil
}
