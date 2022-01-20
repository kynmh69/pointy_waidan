package control

import (
	"database/sql"
	"fmt"
	"pointy/model"
	"pointy/utils"

	_ "github.com/lib/pq"
)

const (
	DRIVER_NAME = "postgres"
	IP_ADDRESS  = "postgres"
	USER        = "postgres"
	PASS        = "postgres"
	DB_NAME     = "postgres"
)

func GetAllData() *[]*model.Article {
	db := connectDb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM pointy order by id")
	utils.CheckError(err)
	articles := createArticleList(rows)
	// fmt.Printf("%v\n", &articles)
	return articles
}

func InsertData(a *[]*model.Article) {
	db := connectDb()
	defer db.Close()
	for _, v := range *a {
		r, err := db.Exec("INSERT INTO pointy (title) VALUES ($1);", v.Title)
		utils.CheckError(err)
		fmt.Println("insert result: ", r)
	}

}

func UpdateData(a *[]*model.Article) {
	db := connectDb()
	defer db.Close()
	for _, v := range *a {
		r, err := db.Exec("UPDATE pointy SET title = $1 WHERE id = $2;", v.Title, v.Number)
		utils.CheckError(err)
		fmt.Println("result: ", r)
	}
}

func GetLastRow() *model.Article {
	db := connectDb()
	defer db.Close()

	row := db.QueryRow("select * from pointy order by id desc limit 1;")
	if row.Err() != nil {
		return nil
	}
	a := scanRow(row)

	fmt.Println("a: ", *a)
	return a
}

func createArticleList(rows *sql.Rows) *[]*model.Article {
	var articles []*model.Article
	for rows.Next() {
		a := scanRows(rows)
		articles = append(articles, a)
	}
	return &articles
}

func scanRow(row *sql.Row) *model.Article {
	var a model.Article
	row.Scan(&a.Number, &a.Title)
	return &a
}

func scanRows(rows *sql.Rows) *model.Article {
	var a model.Article
	rows.Scan(&a.Number, &a.Title)
	return &a
}

func connectDb() *sql.DB {
	connectionInfo := GetConnectionInfo()
	db, err := sql.Open(DRIVER_NAME, connectionInfo)
	utils.CheckError(err)
	return db
}

func CreateTable() {
	db := connectDb()
	_, err := db.Exec("create table pointy (id serial, title text);")
	utils.CheckError(err)
}

func GetConnectionInfo() string {
	return fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", IP_ADDRESS, USER, PASS, DB_NAME)
}
