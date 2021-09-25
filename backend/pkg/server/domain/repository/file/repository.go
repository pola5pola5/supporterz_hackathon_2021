package file

import "io"

type FileRepo interface {
	SaveFile(filename string, file io.Reader) (string, error)
	DeleteFile(filename string) error
	CreatepreSignedURL(filename string) (string, error)
}
