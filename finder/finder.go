package finder

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/h2non/gentleman.v2"
	"log"
	"reflect"
	"strings"
)

var (
	sites = map[string]string{
		"JavBus": "https://www.javbus.com",
	}

	removeString = map[string]string{
		"Number":    "識別碼:",
		"PublishAt": "發行日期:",
		"Length":    "長度:",
		"Director":  "導演:",
		"Producer":  "製作商:",
		"Publisher": "發行商:",
		"Series":    "系列:",
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

	t := reflect.ValueOf(&lyr).Elem()

	doc.Find(".movie .info p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		for header, replaceString := range removeString {
			if strings.Contains(text, replaceString) {
				text = strings.TrimSpace(strings.Replace(text, replaceString, "", 1))

				t.FieldByName(header).SetString(text)
			}
		}
	})

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
