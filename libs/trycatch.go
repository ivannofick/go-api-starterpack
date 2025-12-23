package libs

import (
	"log"
	"net/http"
)

func TryCatch(
	w http.ResponseWriter,
	fn func(),
) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[ERROR]", err)

			ResponseAPI(
				w,
				[]any{},
				nil,
				"Server error....",
				1,
			)
		}
	}()

	fn()
}
