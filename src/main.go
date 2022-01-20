package main

import (
	"fmt"
	"pointy/control"
	"pointy/utils"
)

func updateDB() {
	max_id := 0
	article := control.GetLastRow()
	if article != nil {
		max_id = article.Number
	}
	fmt.Println("max id: ", max_id)

	// htmlから情報吸い出し
	article_list := *utils.GetArticle(max_id)
	if len(article_list) <= 0 {
		fmt.Println("空リスト")
	}

	// テーブルに挿入
	control.InsertData(&article_list)
	// 全てのデータ取得
	article_list = *control.GetAllData()
	for _, v := range article_list {
		fmt.Println(v.GetUrl())
	}

}
func main() {
	updateDB()
}
