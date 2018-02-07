package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	_"go/doc"
)


func main() {
	fmt.Println("ok")
	ExampleScrape()
}

func ExampleScrape() {
	// doc, err := goquery.NewDocument("http://metalsucks.net")
	doc, err := goquery.NewDocument("http://m.ximalaya.com/35878101/album/3475911")

	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	/*doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})*/

	doc.Find("h4").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//fmt.Println("s is",s.Text())
		//band := s.Find("h4").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})

	d,err1 := goquery.NewDocument("http://m.ximalaya.com/album-quan/kid-绘本/rank")
	if err1 != nil{
		log.Fatal(err1)
	}

	d.Find(".name").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Println("s is",s.Text())
		//band := s.Find("h4").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}