package main

import (
	"fmt"
	"io"
	"net/http"
	"pointy/model"
	"regexp"
)

func main() {
	count := 1
	r := regexp.MustCompile(`<title>.*</title>`)
	r_break := regexp.MustCompile(`<title>THE猥談募集フォーム</title>`)
	artice_list := make([]model.Article, 0)
	for {
		rsp, shouldReturn := Get(count)
		if shouldReturn {
			return
		}
		defer rsp.Body.Close()
		body_byte, _ := io.ReadAll(rsp.Body)
		body := string(body_byte)

		m_str := r.FindString(body)
		fmt.Println("match string: ", m_str)

		if r_break.MatchString(body) {
			break
		}
		artice_list = append(artice_list, *model.New(count, m_str))
		count += 1
	}
	for _, v := range artice_list {
		fmt.Println(v.GetUrl())
	}

}

func Get(num int) (*http.Response, bool) {
	url := fmt.Sprintf("https://thewaidan.studio.site/%d", num)
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, true
	}
	return rsp, false
}
