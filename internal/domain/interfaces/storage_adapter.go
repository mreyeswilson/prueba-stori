package interfaces

import "io"

type IStorageAdapter interface {
	GetObject(bucketName string, key string) (io.ReadCloser, error)
}
