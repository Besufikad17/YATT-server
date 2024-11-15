package handlers

import (
	"encoding/json"
	"io"

	"net/http"

	"github.com/Besufikad17/YATT-server/db"
	"github.com/Besufikad17/YATT-server/internal"
	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusBadRequest, "Invalid payload")
		return
	}

	var actionPayload internal.LoginParams
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusInternalServerError, "Error unmarshalling payload")
		return
	}

	if err := internal.Validate.Struct(actionPayload); err != nil {
		internal.ThrowHttpError(w, err, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	queries := db.New(h.DBConn)
	user, err := queries.GetUserByEmail(h.Ctx, actionPayload.Email)
	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusInternalServerError, err.Error())
		return
	}

	valid := internal.Compare(user.Password, actionPayload.Password)
	if valid {
		token, err := internal.CreateToken(user, actionPayload.RememberMe)
		if err != nil {
			internal.ThrowHttpError(w, err, http.StatusInternalServerError, "Error creating token"+err.Error())
			return
		}
		w.Write(internal.Marshall(nil, map[string]interface{}{
			"token": token,
		}, true))
	} else {
		internal.ThrowHttpError(w, err, http.StatusInternalServerError, "Invalid credentails")
		return
	}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusBadRequest, "Invalid payload")
		return
	}

	var actionPayload internal.SignupParams
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusInternalServerError, "Error unmarshalling payload")
		return
	}

	if err := internal.Validate.Struct(actionPayload); err != nil {
		internal.ThrowHttpError(w, err, http.StatusBadRequest, "Validation error: "+err.Error())
		return

	}

	queries := db.New(h.DBConn)
	hashedPassword, err := internal.Hash(actionPayload.Password)
	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusInternalServerError, "Error hashing password"+err.Error())
		return
	}

	var newUser db.User
	newUser, err = queries.CreateUser(h.Ctx, db.CreateUserParams{
		FirstName: actionPayload.FirstName,
		LastName:  actionPayload.LastName,
		Email:     actionPayload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		internal.ThrowHttpError(w, err, http.StatusInternalServerError, "Error adding new user "+err.Error())
		return
	}

	token, err := internal.CreateToken(newUser, actionPayload.RememberMe)
	w.Write(internal.Marshall(nil, map[string]interface{}{
		"token": token,
	}, true))
}
