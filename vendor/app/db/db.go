package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	_ "strconv"

)

type Env struct {
	db DBContext
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}

type User struct {
	userid int64
	name string
}

type DBContext interface {
	Users() ([]*User, error)
}

func (db *DB) Users() ([]*User, error) {
	log.Println("USERS FUNC")
	rows, err := db.Query("SELECT * FROM users")
    if err != nil {
		log.Println(err.Error())
        return nil, err
	}
	log.Println("USERS B CLOSE")
	defer rows.Close()
	log.Println("USERS AFTER CLOSE")
	
	users := make([]*User, 0)
	log.Println("ROWS")
	for rows.Next() {
		var user User
		err = rows.Scan(&user.userid, &user.name)
		if err != nil{
			log.Println(err.Error())
		}
		users = append(users, &user)
	}

	return users, nil
}