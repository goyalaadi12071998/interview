package controllers

import (
	"encoding/json"
	errorclass "interview/app/error"
	"net/http"
)

var AppController appcontroller

var isHealthy bool = true

type appcontroller struct {
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

type Error struct {
	Name    string
	Code    string
	Message string
}

func InitializeAppController() {
	AppController = appcontroller{}
}

func (a appcontroller) Health(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		response := map[string]string{
			"message": "Server is healthy",
		}
		Respond(w, r, response, nil)
		return
	}
	Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestError).Wrap("Server is unhealthy"))
	return
}

func (a appcontroller) Get(w http.ResponseWriter, r *http.Request) {
	Respond(w, r, "Welcome to Bitespeed", nil)
}

func (a appcontroller) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	Respond(w, r, nil, errorclass.NewError(errorclass.InternalServerError).Wrap("Endpoint does not exist"))
}

func Respond(w http.ResponseWriter, r *http.Request, payload interface{}, err *errorclass.Error) {
	if err != nil {
		w.WriteHeader(err.StatusCode())
		errorResponse := ErrorResponse{
			Success: false,
			Error: Error{
				Message: err.Description(),
				Code:    err.Code(),
				Name:    err.Name(),
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		w.WriteHeader(200)
		successResponse := SuccessResponse{
			Success: true,
			Data:    payload,
		}
		json.NewEncoder(w).Encode(successResponse)
	}
}
