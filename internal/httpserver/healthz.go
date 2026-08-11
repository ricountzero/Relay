package httpserver

import (
	"log/slog"
	"net/http"
)

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		slog.Error("write header error", "err", err)
	}
}
