package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Room represents the structure of a room in the game
type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
}

// RoomsResponse represents the response structure for rooms
type RoomsResponse struct {
	Status int    `json:"status"`
	Data   []Room `json:"data"`
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	// Assuming rooms data is retrieved from somewhere
	rooms := []Room{
		{ID: 1, RoomName: "Room 1"},
		{ID: 2, RoomName: "Room 2"},
		// Add more rooms as needed
	}

	// Create response object
	response := RoomsResponse{
		Status: 200,
		Data:   rooms,
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

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/rooms", GetRooms).Methods("GET")
	return r
}
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	// Parse request body or URL parameters to extract necessary data, e.g., room ID and account ID
	roomID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	accountID, err := strconv.Atoi(r.FormValue("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := sql.Open("mysql", "username:password@tcp(hostname:port)/database")
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query the maximum number of players allowed for the game associated with the room
	var maxPlayers int
	err = db.QueryRow("SELECT max_player FROM Games WHERE id = (SELECT id_game FROM Rooms WHERE id = ?)", roomID).Scan(&maxPlayers)
	if err != nil {
		http.Error(w, "Failed to query maximum players", http.StatusInternalServerError)
		return
	}

	// Query the current number of participants in the room
	var participantCount int
	err = db.QueryRow("SELECT COUNT(*) FROM Participants WHERE id_room = ?", roomID).Scan(&participantCount)
	if err != nil {
		http.Error(w, "Failed to query participant count", http.StatusInternalServerError)
		return
	}

	// Check if the room has reached the maximum number of players
	if participantCount >= maxPlayers {
		http.Error(w, "Room is full", http.StatusForbidden)
		return
	}

	// Insert the participant into the Participants table
	_, err = db.Exec("INSERT INTO Participants (id_room, id_account) VALUES (?, ?)", roomID, accountID)
	if err != nil {
		http.Error(w, "Failed to insert participant into room", http.StatusInternalServerError)
		return
	}

	// Send success response
	response := map[string]string{"message": "Successfully joined the room"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	// Parse request body or URL parameters to extract necessary data, e.g., room ID and participant ID
	roomID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	participantID, err := strconv.Atoi(r.FormValue("participant_id"))
	if err != nil {
		http.Error(w, "Invalid participant ID", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := sql.Open("mysql", "username:password@tcp(hostname:port)/database")
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if the participant exists in the specified room
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM Participants WHERE id_room = ? AND id = ?)", roomID, participantID).Scan(&exists)
	if err != nil {
		http.Error(w, "Failed to check participant existence", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Participant does not exist in the room", http.StatusNotFound)
		return
	}

	// Delete the participant from the Participants table
	_, err = db.Exec("DELETE FROM Participants WHERE id_room = ? AND id = ?", roomID, participantID)
	if err != nil {
		http.Error(w, "Failed to delete participant from room", http.StatusInternalServerError)
		return
	}

	// Send success response
	response := map[string]string{"message": "Successfully left the room"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
