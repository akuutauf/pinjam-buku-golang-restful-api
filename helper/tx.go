package helper

import (
	"gorm.io/gorm"
)

// membuat function untuk melakukan commit atau rollback transaction
func CommitOrRollback(tx *gorm.DB) {
	// melakukan recover error
	err := recover()

		// mengecek jika terjadi error
		if err != nil {
			// maka semua proses transaction di database akan dibatalkan
			errorRollback := tx.Rollback().Error

			// mengeceke error kembali pada saat rollback (overlapping)
			PanicIfError(errorRollback)

			// kemudian akan mengembalikan error untuk diterima di kode atas
			panic(err)
		} else {
			// tapi kalau misalnya tidak error, maka lakukan commit
			errorCommit := tx.Commit().Error

			// mengecek error kembali pada saat commit (overlapping)
			PanicIfError(errorCommit)
		}
}