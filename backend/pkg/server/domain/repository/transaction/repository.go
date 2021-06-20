package transaction

import "database/sql"

type TxRepo interface {
	Transaction(f func(tx *sql.Tx) error) error
}
