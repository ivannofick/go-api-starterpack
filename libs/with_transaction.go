package libs

import (
	"log"
	"net/http"

	"go-crud/config"

	"gorm.io/gorm"
)

func WithTransaction(
	fn func(w http.ResponseWriter, r *http.Request, tx *gorm.DB),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tx := config.DB.Begin()
		if tx.Error != nil {
			log.Println("[ERROR] Failed to start transaction:", tx.Error)
			ResponseAPI(w, []any{}, nil, "Server error....", 1)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				tx.Rollback()
				log.Println("Transaction Rolled Back Due To Error:", err)
				ResponseAPI(w, []any{}, nil, "Server error....", 1)
			}
		}()

		fn(w, r, tx)

		if err := tx.Commit().Error; err != nil {
			log.Println("[ERROR] Commit failed:", err)
			ResponseAPI(w, []any{}, nil, "Server error....", 1)
		}
	}
}
