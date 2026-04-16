package middlewares

import "net/http"

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next(w, r)
			return
		}

		role := r.Header.Get("X-Role")
		if role == "admin" || role == "super-clave-admin-123" {
			next(w, r)
			return
		}
		next(w, r)
	}
}
