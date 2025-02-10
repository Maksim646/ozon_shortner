package handler

import "net/http"

func (h *Handler) WsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// if req.Method == "OPTIONS" {
		// 	w.Header().Set("Access-Control-Allow-Origin", "*")
		// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// 	w.Header().Set("Access-Control-Allow-Headers", "*")
		// 	return
		// }
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, req)
	})
}
