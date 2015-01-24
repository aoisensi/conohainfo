package conohainfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	urlList   = "https://cp.conoha.jp/Information.aspx"
	urlGetter = "https://cp.conoha.jp/GetInforMation.aspx?mid="
)

type Page struct {
	Title, Id string
}

type Info struct {
	Category, CategoryCssClass, Subject, Date, Body string
}

func GetList() ([]Page, error) {
	doc, err := goquery.NewDocument(urlList)
	if err != nil {
		return nil, err
	}
	list := doc.Find("dl.newsList a")
	result := make([]Page, list.Length())
	list.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		page := Page{Id: href[1:], Title: s.Text()}
		result[i] = page
	})
	return result, nil
}

func (p *Page) GetInfo() (*Info, error) {
	url := urlGetter + p.Id
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := new(Info)

	return data, json.Unmarshal(body, data)
}
