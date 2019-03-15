package javbus

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/lyraeve/lyrae-go/contracts"
	"github.com/lyraeve/lyrae-go/finder"
	"gopkg.in/h2non/gentleman.v2"
	"log"
	"reflect"
	"strings"
)

var (
	baseUrl = "https://www.javbus.com"

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

type Javbus struct {

}

func New() Javbus {
	b := new(Javbus)
	return *b
}

func (_ Javbus) FindByNumber(number string) (contracts.Lyr, error) {
	var lyr finder.Lyr

	client := gentleman.New()
	client.URL(baseUrl)

	req := client.Request()
	req.Path("/" + number)

	res, err := req.Send()

	if err != nil {
		return lyr, err
	}

	if !res.Ok {
		log.Fatal("Invalid server response: " + string(res.StatusCode))
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

	lyr = parseScreenShots(doc, lyr)

	return lyr, nil
}

func parseScreenShots(doc *goquery.Document, lyr finder.Lyr) finder.Lyr {
	doc.Find("#sample-waterfall .photo-frame img").Each(func(i int, s *goquery.Selection) {
		src, exist := s.Attr("src")

		if exist {
			lyr.ScreenShots = append(lyr.ScreenShots, src)
		}
	})

	return lyr
}
