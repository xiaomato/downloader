package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
)

func main() {
	reg, err := regexp.Compile("http.*?\\.m3u8")
	if err != nil {
		return
	}
	exsit := make(map[string]bool)
	for true {
		time.Sleep(time.Second)

		d, err := clipboard.ReadAll()
		if err != nil {
			println(err.Error())
			continue
		}

		text := string(d)
		if !strings.HasPrefix(text, "http") {
			continue
		}
		if exsit[text] {
			continue
		}
		exsit[text] = true

		rsp, err := http.Get(text)
		if err != nil {
			continue
		}
		html, _ := ioutil.ReadAll(rsp.Body)
		rsp.Body.Close()
		m := reg.Find(html)
		doc, _ := goquery.NewDocument(text)
		title := trim(doc.Find("h1.py-2").Text())
		go func(t, mu string) {
			d, _ := NewM3u8Downloader(mu, t, 3)
			_ = d.Download()

		}(title, string(m))
	}
}

func trim(t string) string {
	t = strings.ReplaceAll(t, " ", "")
	t = strings.ReplaceAll(t, "\n", "")
	t = strings.ReplaceAll(t, "\t", "")
	return t
}
