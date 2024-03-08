package controllers

type Game struct {
	ID        int    `json:"gameId"`
	Name      string `json:"gameName"`
	MaxPlayer int    `json:"maxPlayer"`
}

type Rooms struct {
	ID        int    `json:"roomId"`
	RoomName  string `json:"roomName"`
	GameID    int    `json:"gameId"`
	MaxPlayer int    `json:"maxPlayer"`
}

type Participant struct {
	ID        int `json:"participantId"`
	RoomID    int `json:"roomId"`
	AccountID int `json:"accountId"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
