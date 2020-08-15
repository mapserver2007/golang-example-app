package storage

type Storage interface {
	AppendLog(logId string, data string) error

	Lookup(logId string) ([]string, error)

	Close() error

	Cleanup(logId string) error
}

type StorageProvider func(cfg StorageConfig) Storage

type StorageConfig struct {
	Redis struct {
		Host     string
		Port     string
		Password string
	}
}
