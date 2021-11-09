package apiserver

import "net/http"

type responseWriter struct {
	 http.ResponseWriter
	 code int
}

func (w *responseWriter) WriterHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}