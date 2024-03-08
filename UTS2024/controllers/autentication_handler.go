package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Player represents the structure of a player in a room
type Player struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// RoomDetail represents the detailed information of a room
type RoomDetail struct {
	ID           int      `json:"id"`
	RoomName     string   `json:"room_name"`
	Participants []Player `json:"participants"`
}

// RoomDetailResponse represents the response structure for room detail
type RoomDetailResponse struct {
	Status int        `json:"status"`
	Data   RoomDetail `json:"data"`
}

// GetRoomDetail retrieves detailed information of a room
func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	// Assuming room detail data is retrieved from somewhere
	// Here, we mock the data for demonstration purpose
	roomID := mux.Vars(r)["id"]
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	// Mock room detail data
	roomDetail := RoomDetail{
		ID:       roomIDInt,
		RoomName: "Room " + roomID,
		Participants: []Player{
			{ID: 1, Username: "Player 1"},
			{ID: 2, Username: "Player 2"},
			// Add more players as needed
		},
	}

	// Create response object
	response := RoomDetailResponse{
		Status: 200,
		Data:   roomDetail,
	}

	// Convert response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
