package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error :", err)
	}

}
func wFile(data, fileName string) {
	file, err := os.Create(fileName)
	checkErr(err)
	defer file.Close()

	file.WriteString(data)
}

func main() {
	url := "https://techcrunch.com/"

	res, err := http.Get(url)
	checkErr(err)
	defer res.Body.Close()

	if res.StatusCode > 400 {
		fmt.Println("Status code :", res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	csvfile, err := os.Create("sData.csv")
	checkErr(err)
	writer := csv.NewWriter(csvfile)

	doc.Find("div.river").Find("div.post-block").Each(func(i int, item *goquery.Selection) {
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url, _ := h2.Find("a").Attr("href")
		//fmt.Println(url)
		text := strings.TrimSpace(item.Find("div.post-block__content").Text())
		posts := []string{title, url, text}
		writer.Write(posts)
	})
	writer.Flush()
	//	checkErr(errr)
	//wFile(river, "htmlFile")
}
