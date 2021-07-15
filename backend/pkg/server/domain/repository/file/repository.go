package file

type FileRepo interface {
	SaveFile(filename string, file []byte) (string, error)
	DeleteFile(filename string) error
}
