package utils

import (
	"fmt"
	"io"
	"net/http"
	"pointy/model"
	"regexp"
)

const TITLE_REGX = `<title>.*</title>`
const BREAK_TITLE_REGX = `<title>THE猥談募集フォーム</title>`

func GetArticle(count int) *[]*model.Article {
	artice_list := make([]*model.Article, 0)
	r := regexp.MustCompile(TITLE_REGX)
	r_break := regexp.MustCompile(BREAK_TITLE_REGX)
	//最終記事に+1する。
	count++

	for {
		// get html
		rsp := Get(count)
		defer rsp.Body.Close()
		body_byte, _ := io.ReadAll(rsp.Body)
		body := string(body_byte)

		// get title
		m_str := r.FindString(body)
		fmt.Println("match string: ", m_str)

		// 記事のタイトルじゃなかったら終了
		if r_break.MatchString(body) {
			break
		}
		// 構造体生成
		new_article := *model.New(count, m_str)
		// タイトルタグ削除
		new_article = *ReplaceTitle(&new_article)
		// スライスに追加
		artice_list = append(artice_list, &new_article)
		count++
	}
	for _, v := range artice_list {
		fmt.Println(v.GetUrl())
	}
	return &artice_list
}

func Get(num int) *http.Response {
	url := fmt.Sprintf("https://thewaidan.studio.site/%d", num)
	rsp, err := http.Get(url)
	if err != nil {
		CheckError(err)
		return nil
	}
	return rsp
}
