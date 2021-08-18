package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     Keys   `json:"keys"`
}

/*
func main() {
	db, err := InitDB("hi.db")
	if err != nil {
		log.Fatalln(err)
	}
	AddUser(db, "p1", "0000")
	AddUser(db, "p2", "0200")
	RemoveUser(db, "p1")
	GetAllUser(db)

	defer db.Close()
}
*/

func InitDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
	create TABLE IF NOT EXISTS useracount (
		id INTEGER PRIMARY KEY autoincrement,
		endpoint TEXT,
		auth TEXT,
		pdh TEXT,
		UNIQUE (id, auth)
	)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("=+==+===")
		return nil, err
	}

	return db, nil
}

func AddUser(db *sql.DB, sub Subscription) error {
	var exists bool
	row := db.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM useracount WHERE auth='%s')", sub.Keys.Auth))
	if err := row.Scan(&exists); err != nil {
		return err
	} else if !exists {
		tx, _ := db.Begin()
		stmt, _ := tx.Prepare("INSERT INTO useracount (endpoint, auth, pdh) values (?,?,?)")
		_, err := stmt.Exec(sub.Endpoint, sub.Keys.Auth, sub.Keys.P256dh)
		if err != nil {
			fmt.Println("=++=====")
			log.Println(err.Error())
			return err
		}
		tx.Commit()
		return nil
	}
	fmt.Println("already exits!")
	return nil
}

func RemoveUser(db *sql.DB, id string) error {
	tx, _ := db.Begin()
	stmt, err := tx.Prepare("DELETE FROM useracount WHERE auth=?")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	affect, err := res.RowsAffected()
	fmt.Println(affect)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func GetAllUser(db *sql.DB) ([]Subscription, error) {
	var endpoint string
	var auth string
	var p256dh string
	rows, err := db.Query("SELECT endpoint, auth, pdh FROM useracount")
	if err != nil {
		fmt.Println("=+==_===")
		return nil, err
	}
	defer rows.Close()

	var list []Subscription

	for rows.Next() {
		err := rows.Scan(&endpoint, &auth, &p256dh)
		if err != nil {
			fmt.Println("=+=__====")
			return nil, err
		}
		fmt.Println(auth)
		list = append(list, Subscription{
			Endpoint: endpoint,
			Keys: Keys{
				Auth:   auth,
				P256dh: p256dh,
			},
		})
	}
	return list, nil
}
