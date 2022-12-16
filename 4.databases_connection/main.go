package main

import (
	"context"
	"dbconn/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// type Books struct {
// 	id        int64
// 	title     string
// 	author_id int64
// 	new_book  string
// }

// type NewBook struct {
// 	title     string
// 	author_id int64
// 	new_book  string
// }

func main() {
	ctx := context.Background()

	db, err := database.CreateConnection()

	if err != nil {
		log.Fatal(err)
	}

	err = database.QueryBooks(ctx, db)
	if err != nil {
		panic(err)
	}

	b := database.NewBook{
		Title:     "Sebas el mejor",
		Author_id: 69,
		New_book:  "new",
	}

	err = database.AddBook(ctx, db, b)
	if err != nil {
		log.Fatal(err)
	}

	err = database.QueryBook(ctx, db, 6)
	if err != nil {
		log.Fatal(err)
	}
	// connectionString := "golang:sebastian12345@tcp(localhost:3306)/golang"
	// // connectionString := "golang:sebastian12345@tcp(localhost:3306)/books"

	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// db.SetMaxOpenConns(5)

	// err = db.Ping()
	// if err != nil {
	// 	log.Panic(err)
	// }

	db.Close()

}

// func createConnection() (*sql.DB, error) {
// 	connectionString := "golang:sebastian12345@tcp(localhost:3306)/golang?parseTime=True"
// 	// connectionString := "golang:sebastian12345@tcp(localhost:3306)/books"

// 	db, err := sql.Open("mysql", connectionString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	db.SetMaxOpenConns(5)

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func queryBooks(ctx context.Context, db *sql.DB) error {
// 	// PARA VARIOS

// 	qry := `SELECT * FROM BOOKS`

// 	rows, err := db.QueryContext(ctx, qry)
// 	if err != nil {
// 		return err
// 	}

// 	books := []Books{}

// 	for rows.Next() {
// 		b := Books{}

// 		err = rows.Scan(&b.id, &b.title, &b.author_id, &b.new_book)
// 		if err != nil {
// 			return err
// 		}

// 		books = append(books, b)
// 	}

// 	fmt.Println(books)

// 	return nil
// }

// // PARA UNO SOLO
// func queryBook(ctx context.Context, db *sql.DB, id_qry int) error {

// 	qry := `SELECT * FROM BOOKS WHERE id = ?`

// 	row := db.QueryRowContext(ctx, qry, id_qry)

// 	var id int64
// 	var title string
// 	var author_id int64
// 	var new_book string

// 	err := row.Scan(&id, &title, &author_id, &new_book)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("ROW:", id, title, author_id, new_book)
// 	return nil
// }

// func addBook(ctx context.Context, db *sql.DB, b NewBook) error {
// 	qryadd := `INSERT INTO books (title,author_id,new_book)
// 		VALUES (?,?,?)`

// 	result, err := db.ExecContext(ctx, qryadd, b.title, b.author_id, b.new_book)
// 	if err != nil {
// 		return err
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return nil
// 	}

// 	fmt.Println("INSERT ID:", id)

// 	return nil
// }
