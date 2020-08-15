package storage

type Storage interface {
	AppendLog(logId string, data string) error

	Lookup(logId string) ([]string, error)

	Close() error

	LogIds() ([]string, error)

	Cleanup(logId string) error

	LastLog(logId string) (string, error)
}

type StorageProvider func(cfg StorageConfig) Storage

type StorageConfig struct {
	Redis struct {
		Address string
	}
}
