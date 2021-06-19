package file

type FileRepo interface {
	SaveFile(filename string, file []byte) error
}
