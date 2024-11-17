package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	if err := recover(); err != nil {
		// Rollback if a panic occurred
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)

		// Re-panic to propagate the original error after rollback
		panic(err)
	} else {
		// Commit if no panic occurred
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
