package middleware

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Definir os cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")               // Permitir todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT") // Métodos permitidos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")   // Cabeçalhos permitidos

		// Se for uma requisição OPTIONS, responder imediatamente (útil para pré-você de CORS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Chama o próximo handler (aplicação real)
		next.ServeHTTP(w, r)
	})
}
