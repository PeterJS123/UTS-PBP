package controllers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InsertDummyData inserts dummy data into the database tables
func InsertDummyData() {
	// Open a database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/game")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert dummy data into Accounts table
	_, err = db.Exec("INSERT INTO Accounts (username) VALUES ('player1'), ('player2'), ('player3'), ('player4')")
	if err != nil {
		log.Fatal(err)
	}

	// Insert dummy data into Games table
	_, err = db.Exec("INSERT INTO Games (name, max_player) VALUES ('Peter', 4), ('Retep', 6), ('Terpe', 8), ('Jaya', 10)")
	if err != nil {
		log.Fatal(err)
	}

	// Insert dummy data into Rooms table
	_, err = db.Exec("INSERT INTO Rooms (room_name, id_game) VALUES ('Room 1', 1), ('Room 2', 2), ('Room 3', 3), ('Room 4', 4)")
	if err != nil {
		log.Fatal(err)
	}

	// Insert dummy data into Participants table
	_, err = db.Exec("INSERT INTO Participants (id_room, id_account) VALUES (1, 1), (1, 2), (2, 3), (3, 4)")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dummy data inserted successfully")
}
