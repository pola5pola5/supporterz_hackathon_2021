package transaction

import (
	"database/sql"
	"log"

	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/transaction"
)

type txPersistence struct {
	db *sql.DB
}

func NewPersistence(db *sql.DB) transaction.TxRepo {
	return &txPersistence{
		db: db,
	}
}

// Transaction トランザクション処理
func (tp txPersistence) Transaction(function func(tx *sql.Tx) error) error {
	tx, err := tp.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		// panic
		if err := recover(); err != nil {
			log.Println("!! PANIC !!")
			log.Println(err)
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Println("failed to Rollback")
				log.Println(rollbackErr)
			}
		}
	}()

	// トランザクション内での実行
	// TODO: エラーの書き方
	if err = function(tx); err != nil {
		log.Println(err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Println("failed to Rollback")
			log.Println(rollbackErr)
			return rollbackErr
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
