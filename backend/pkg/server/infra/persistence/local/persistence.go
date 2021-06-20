package local

import (
	"os"

	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/file"
)

const path = "/app/main"

type localPersistence struct {
}

func NewPersistence() file.FileRepo {
	return &localPersistence{}
}

// TODO: 返り値にstringは必要か？
// 画像をサーバに保存する処理
func (lp localPersistence) SaveFile(filename string, file []byte) (string, error) {
	// 保存先を作成
	dst, err := os.Create(path + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	// TODO: 正しく保存されているか確認する
	if _, err = dst.Write(file); err != nil {
		return "", err
	}
	return filename, nil
}
