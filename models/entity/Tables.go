package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Table struct {
	TableId     int64  `json:"table"`
	GuestName   string `json:"guest_name"`
	Capacity    int64  `json:"capacity"`
	Fill        int64  `json:"fill"`
	Status      string `json:"status"`
	ArrivalTime string `json:"arrival_time"`
	CreatedAt   string `json:"created_at"`
}

type TableInfo struct {
	TableId   int64  `json:"table"`
	GuestName string `json:"guest_name"`
	Fill      int64  `json:"fill"`
}

type GuestInfo struct {
	GuestName   string `json:"guest_name"`
	Fill        int64  `json:"fill"`
	ArrivalTime string `json:"arrival_time"`
}

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "ishan:water1211@/GetGroundParty")
	if err != nil {
		panic(err.Error())
	}

	DB = db
}

func UpdateTable(tableId int64, accompanyingGuests int64) error {
	result, err := DB.Exec(
		fmt.Sprintf("UPDATE tables SET fill = ? WHERE capacity > ? AND id = ?",
			accompanyingGuests,
			accompanyingGuests,
			tableId,
		))

	if err != nil {
		return err
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if RowsAffected < 0 {
		return errors.New("no capacity")
	}

	return nil
}

func GetGuests(tableInfo *[]TableInfo) error {
	var tableInfoDTO []TableInfo
	rows, err := DB.Query("SELECT guest_name, capacity, id FROM tables")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table TableInfo
		if err := rows.Scan(&table.GuestName, &table.TableId, &table.Fill); err != nil {
			panic(err)
		}
		tableInfoDTO = append(tableInfoDTO, table)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	tableInfo = &tableInfoDTO
	return nil
}

func UpdateTableArrival(tableId int64, accompanyingGuests int64, guestName string) error {
	t := time.Now()
	result, err := DB.Exec(fmt.Sprintf("UPDATE tables SET fill = ?, arrival_time = ?, status = 1 WHERE capacity > ? AND id = ?",
		accompanyingGuests,
		t.String(),
		accompanyingGuests,
		tableId,
	))

	if err != nil {
		return err
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if RowsAffected < 0 {
		return errors.New("no capacity")
	}

	return nil
}

func GetGuestCurrent(guestInfo *[]GuestInfo) error {
	var guestInfoDTO []GuestInfo
	rows, err := DB.Query("SELECT guest_name, fill, arrival_time FROM tables where status = 1")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var guest GuestInfo
		if err := rows.Scan(&guest.GuestName, &guest.Fill, &guest.ArrivalTime); err != nil {
			panic(err)
		}
		guestInfoDTO = append(guestInfoDTO, guest)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return nil
}

func GetEmptySeats() (int64, error) {
	row := DB.QueryRow("SELECT (SELECT sum(capacity) from tables) - (SELECT sum(fill) from tables) as difference")
	var emptySeats int64
	if err := row.Scan(&emptySeats); err != nil {
		return 0, err
	}
	return emptySeats, nil
}

func DeleteGuest(guestName string) error {
	result, err := DB.Exec(
		fmt.Sprintf("DELETE FROM tables WHERE guest_name =? ",
			guestName,
		))

	RowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	} else if RowsAffected < 0 {
		return errors.New("error in deletion")
	}

	return nil
}
