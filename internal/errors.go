package internal

import (
	"log/slog"
	"net/http"
)

func ThrowHttpError(w http.ResponseWriter, err error, statusCode int, message string) {
	slog.With(message).Error(err.Error())
	w.WriteHeader(statusCode)
	w.Write(Marshall(&message, nil, false))
	return
}
