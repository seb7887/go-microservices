package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/seb7887/go-microservices/auth"
)

type ErrResponse struct {
	Error string
}

func ReadBody (r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	return body, err
}

func ParseBody(b []byte, v interface{}) error {
	err := json.Unmarshal(b, v)
	return err
}

func HandleError(w http.ResponseWriter, message string) {
	resp := ErrResponse{Error: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}

func PrepareResponse(w http.ResponseWriter, call map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(call)
}

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()
			if error != nil {
				log.Println(error)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				resp := ErrResponse{Error: "Internal server error"}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func unauthorizerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	resp := ErrResponse{Error: "Unauthorized"}
	json.NewEncoder(w).Encode(resp)
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			unauthorizerError(w)
			return
		}
		userId, valid := auth.ValidateToken(tokenString)
		if !valid {
			unauthorizerError(w)
			return
		}

		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}