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
	IP_ADDRESS  = "127.0.0.1"
	USER        = "postgres"
	PASS        = "postgres"
	DB_NAME     = "postgres"
)

func GetAllData() {
	db := connectDb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM pointy")
	utils.CheckError(err)
	articles := createArticleList(rows)
	fmt.Printf("%v", articles)
}

func InsertData(a *[]model.Article) {
	db := connectDb()
	defer db.Close()
	for _, v := range *a {
		r, err := db.Exec("INSERT INTO pointy (title) VALUES ($1);", v.Title)
		utils.CheckError(err)
		fmt.Println("insert result: ", r)
	}

}

func createArticleList(rows *sql.Rows) []model.Article {
	var articles []model.Article
	for rows.Next() {
		var a model.Article
		rows.Scan(&a.Number, &a.Title)
		articles = append(articles, a)
	}
	return articles
}

func connectDb() *sql.DB {
	connectionInfo := GetConnectionInfo()
	db, err := sql.Open(DRIVER_NAME, connectionInfo)
	utils.CheckError(err)
	return db
}

func GetConnectionInfo() string {
	return fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", IP_ADDRESS, USER, PASS, DB_NAME)
}
