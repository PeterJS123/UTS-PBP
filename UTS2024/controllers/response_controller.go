package controllers

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, r *http.Request, req interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func SendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
