package finder

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/h2non/gentleman.v2"
	"log"
	"strings"
)

var (
	sites = map[string]string{
		"JavBus": "https://www.javbus.com",
	}

	removeString = map[string]string{
		"number":     "識別碼:",
		"publish_at": "發行日期:",
		"length":     "長度:",
		"director":   "導演:",
		"producer":   "製作商:",
		"publisher":  "發行商:",
		"series":     "系列:",
	}
)

func FindByNumber(number string) (lyr Lyr, err error) {
	client := gentleman.New()
	client.URL(sites["JavBus"])

	req := client.Request()
	req.Path("/" + number)

	res, err := req.Send()

	if err != nil {
		return lyr, err
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return lyr, nil
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res.String()))

	if err != nil {
		log.Fatal(err)

		return lyr, err
	}

	lyr.Title = doc.Find("h3").First().Text()
	lyr.Cover = doc.Find(".movie .screencap img").First().AttrOr("src", "")

	data := map[string]string{}

	doc.Find(".movie .info p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		for header, replaceString := range removeString {
			if strings.Contains(text, replaceString) {
				text = strings.TrimSpace(strings.Replace(text, replaceString, "", 1))
				data[header] = text
			}
		}
	})

	for header, content := range data {
		fmt.Println(header + ": " + content)
	}

	doc.Find(".movie .info p .genre").Each(func(i int, s *goquery.Selection) {
		if strings.Compare("", s.AttrOr("onmouseover", "")) == 0 {
			lyr.Actors = append(lyr.Actors, strings.TrimSpace(s.Text()))
		} else {
			lyr.Genres = append(lyr.Genres, strings.TrimSpace(s.Text()))
		}
	})

	doc.Find("#sample-waterfall .photo-frame img").Each(func(i int, s *goquery.Selection) {
		lyr.ScreenShots = append(lyr.ScreenShots, s.AttrOr("src", ""))
	})

	return lyr, nil
}
