package routes

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shi-gg/githook/config"
	"github.com/shi-gg/githook/utils"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	conf := config.Get()

	url := r.URL.Query().Get("url")

	id, err := utils.Encrypt(url, []byte(conf.Secret))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
	w.WriteHeader(http.StatusOK)
}
