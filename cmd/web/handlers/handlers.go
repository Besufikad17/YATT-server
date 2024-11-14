package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Ctx    context.Context
	DBConn *pgxpool.Pool
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "hi :)",
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(data)
}

func NewHandler(ctx context.Context, dbConn *pgxpool.Pool) *Handler {
	return &Handler{
		Ctx:    ctx,
		DBConn: dbConn,
	}
}
