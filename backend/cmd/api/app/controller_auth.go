package app

import (
	"encoding/json"
	"net/http"
)

// CsrfController issues a CSRF cookie compatible with the frontend ApiService expectations
func (i *Instance) CsrfController(w http.ResponseWriter, r *http.Request) {
    // simple static token sufficient for demo; in real apps use a random secret per session
    http.SetCookie(w, &http.Cookie{
        Name:     "mock_store_csrftoken",
        Value:    "demo-token",
        Path:     "/",
        HttpOnly: false,
        Secure:   false,
        SameSite: http.SameSiteLaxMode,
    })

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (i *Instance) WhoAmIController(w http.ResponseWriter, r *http.Request) {
	type user struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).
		Encode(user{
			ID:        1,
			FirstName: "Иван",
			LastName:  "Иванов",
			Email:     "momolover@mail.ru",
		})
}
