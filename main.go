package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gocolly/colly"
)

type Currency struct {
	Title    string `json:"title"`
	PriceUsd string `json:"priceUsd"`
	PriceBtc string `json:"PriceBtc"`
	Sum      string `json:"sum"`
}

var count int
var currencies = []Currency{}

func main() {
	scrapPage("https://bitinfocharts.com/ru/spisok-luchshih-kriptovalyut.html")
	println("Find ", count, " currencies")
	writeResultToXls()
}

func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		if e.Attr("class") == "ptr" {
			title := e.DOM.Find("td:nth-child(2)").Text()
			priceUsd := e.DOM.Find("td:nth-child(3) a").Text()
			priceBtc := e.DOM.Find("td:nth-child(4) span:nth-child(1)").Text()
			sum := e.DOM.Find("td:nth-child(6) span:nth-child(1)").Text()
			cur := Currency{title, priceUsd, priceBtc, sum}
			count++
			currencies = append(currencies, cur)
		}
	})

	c.Visit(url)
}

func writeResultToXls() {
	xlsx := excelize.NewFile()

	xlsx.NewSheet("Sheet1")

	for i, cur := range currencies {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+1), cur.Title)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("B%v", i+1), cur.PriceUsd)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%v", i+1), cur.PriceBtc)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("D%v", i+1), cur.Sum)
	}
	err := xlsx.SaveAs("./currencies.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	println("Finish...")
}
