package helper

import (
	"gorm.io/gorm"
)

// membuat function untuk melakukan commit atau rollback transaction
func CommitOrRollback(tx *gorm.DB, err error) {
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
