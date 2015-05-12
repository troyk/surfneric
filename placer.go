package surfneric

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

var siteURL = "http://www.placer.courts.ca.gov"

const (
	searchPath = "/CaseCalSearch/CaseIndex2/CaseIndex_list.asp"
)

// Request searches the site
type Request struct {
	CaseYear  string
	FirstName string
	LastName  string

	body string
}

// ToSearchURL converts the struct into a valid url
func (r *Request) ToSearchURL() string {
	return fmt.Sprintf("%s%s?s_F5=%s&s_CN=&s_F3=Criminal&s_F7=%s&s_F4=%s", siteURL, searchPath, sanitizeParam(r.FirstName), sanitizeParam(r.CaseYear), sanitizeParam(r.LastName))
}

func sanitizeParam(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

// Case is a case from the court system
type Case struct {
	Number     string
	Type       string
	LastName   string
	FirstName  string
	Caption    string
	PartyType  string
	Date       string
	ParseError error
}

// Search gets a request from the website
func Search(r *Request) ([]Case, error) {
	b := surf.NewBrowser()
	err := b.Open(r.ToSearchURL())
	if err != nil {
		return nil, err
	}
	r.body = b.Body()
	rows := b.Find("tr.Row")
	cases := make([]Case, rows.Length())
	b.Find("tr.Row").Each(func(trNum int, s *goquery.Selection) {
		c := &cases[trNum]
		tds := s.Find("td")
		if tds.Length() != 6 {
			c.ParseError = fmt.Errorf("row:%d tds is %d (expected 6)", trNum, tds.Length())
		}
		tds.Each(func(tdNum int, s *goquery.Selection) {
			txt := strings.TrimSpace(s.Text())
			switch tdNum {
			case 0:
				c.Number = txt
			case 4:
				c.Caption = txt
			}
			fmt.Println(tdNum, s.Text(), s)
		})
	})
	return cases, err
}
