package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type CompleteArticle struct {
	ArticleUpdates map[int64][]string `json:"updates"`
}

func ParseFR(url string) *CompleteArticle {
	doc, ok := GetDocumentReader(url)
	if !ok {
		return nil
	}

	mainContent := doc.Find(".id-SiteMain")
	if mainContent != nil {
		body := mainContent.Find(".id-Article-body")
		if body == nil {
			return nil
		}
		articels := body.Find(".id-Article-content-item.id-Article-content-item-paragraph")
		if articels == nil {
			return nil
		}
		var compArt CompleteArticle

		compArt.ArticleUpdates = make(map[int64][]string)
		var lastTs int64

		articels.Each(func(i int, article *goquery.Selection) {
			timestring := article.Find("strong").Text()
			if timestring != "" {
				var date time.Time
				var err error
				if strings.Contains(timestring, "Update") {
					date, err = parseTimeLong(timestring)
					if err != nil {
						fmt.Printf("%+v\n", err)
					}

				} else {
					hours, minutes, err := parseTimeToday(timestring)
					if err != nil {
						fmt.Printf("%+v\n", err)
					}
					now := time.Now()
					if now.Hour() < hours {
						now = now.AddDate(0, 0, -1)
					}
					date = time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, now.Location())
				}

				lastTs = date.Unix()
				text := strings.Replace(article.Text(), timestring, "", -1)
				text = strings.Replace(text, ": ", "", 1)
				text = strings.TrimSpace(text)
				compArt.ArticleUpdates[lastTs] = append(compArt.ArticleUpdates[lastTs], text)
			} else {
				compArt.ArticleUpdates[lastTs] = append(compArt.ArticleUpdates[lastTs], article.Text())
			}

		})
		return &compArt
	}

	return &CompleteArticle{}
}

func parseTimeToday(s string) (hours int, minutes int, err error) {
	// (\d+)(?::||.)(\d+)
	reg := regexp.MustCompile(`(\d+)(?::||.)(\d+)`)
	matches := reg.FindAllStringSubmatch(s, -1)
	if len(matches) <= 0 {
		return
	}

	for i, value := range matches[0] {
		if i == 0 {
			continue
		}
		iv, err := strconv.Atoi(value)
		if err != nil {
			return 0, 0, err
		}

		switch i {
		case 1:
			hours = iv
		case 2:
			minutes = iv
		}
	}

	return
}

func parseTimeLong(s string) (time.Time, error) {
	var (
		d, m, y, h, min int
	)

	writtenDateReg := regexp.MustCompile(`((\d+\. )([A-Z]\w+))`)
	wMatches := writtenDateReg.FindAllStringSubmatch(s, -1)

	if len(wMatches) > 0 {
		for i, value := range wMatches[0] {
			if i <= 1 {
				continue
			}

			switch i {
			case 2:
				value = strings.Replace(value, ".", "", -1)
				value = strings.TrimSpace(value)

				t, err := strconv.Atoi(value)
				if err != nil {
					return time.Time{}, err
				}
				d = t
			case 3:
				if value == "Februar" {
					m = 2
				} else if value == "März" || value == "Maerz" {
					m = 3
				}
			}

		}
		y = 2022

		timeReg := regexp.MustCompile(`(\d+)(?::|\.)(\d+)`)
		timeMatch := timeReg.FindAllStringSubmatch(s, -1)

		if len(timeMatch) > 0 {
			for i, value := range timeMatch[0] {
				if i == 0 {
					continue
				}
				iv, err := strconv.Atoi(value)
				if err != nil {
					return time.Time{}, err
				}
				switch i {
				case 1:
					h = iv
				case 2:
					min = iv

				}
			}

		}

	}

	// dateReg := regexp.MustCompile(`(\d+)(?:.)(\d+)(?:.)(\d+)`)
	dateReg := regexp.MustCompile(`((\d+)(?:\.)(\d+)(?:\.)(\d+))|((\d+)(?::)(\d+))`)
	matches := dateReg.FindAllStringSubmatch(s, -1)

	if len(wMatches) == 0 {
		if len(matches) >= 1 {

			// fmt.Printf("%+v\n", matches[0])
			for i, value := range matches[0] {
				if i <= 1 || value == "" {
					continue
				}
				iv, err := strconv.Atoi(value)
				if err != nil {
					return time.Time{}, nil
				}
				switch i {
				case 2:
					d = iv
				case 3:
					m = iv
				case 4:
					y = iv
				}
			}

		}
		if len(matches) == 2 {
			for i, value := range matches[1] {
				if i <= (len(matches[1]) - 2) {
					continue
				}
				iv, err := strconv.Atoi(value)
				if err != nil {
					return time.Time{}, err
				}
				switch i {
				case 1:
					h = iv
				case 2:
					min = iv
				}

			}

		}
	}

	now := time.Now()
	date := time.Date(y, time.Month(m), d, h, min, 0, 0, now.Location())
	return date, nil
}
