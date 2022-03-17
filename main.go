package main
import(
  "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
 )

var schema = `
CREATE TABLE user(
	user_name text,
)`

type User struct {
	UserName string `db:"user_name"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO user(user_name) VALUES($1)", "Olena")
	tx.MustExec("INSERT INTO user(user_name) VALUES($1)", "Jordan")
	tx.MustExec("INSERT INTO user(user_name) VALUES($1)", "Popa")
	tx.Commit()

}
