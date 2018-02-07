package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	_"go/doc"
	"github.com/jinzhu/gorm"
	"strings"
)

type Book struct {

	Id int64 `json:"id" gorm:"AUTO_INCREMENT"`

	// author avatar
	AuthorAvatar string `json:"authorAvatar"`

	// author name
	AuthorName string `json:"authorName"`

	// clips number
	ClipsNumber int64 `json:"clipsNumber"`

	// icon
	Icon string `json:"icon"`

	// name
	Name string `json:"name"`

	// play count
	PlayCount int64 `json:"playCount"`

	// show icon
	ShowIcon bool `json:"showIcon"`

	// status
	Status int64 `json:"status"`

	// sub category Id
	SubCategoryID int64 `json:"subCategoryId"`

	// sub title
	SubTitle string `json:"subTitle"`

	// summary
	Summary string `json:"summary"`

	// time
	Time int64 `json:"time"`
}

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

	db,err := OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(db.Value)

	d.Find(".name").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Println("s is",s.Text())
		//band := s.Find("h4").Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
		var book Book
		book.Summary = strings.TrimSpace(s.Text())
		book.Name = strings.TrimSpace(s.Text())
		book.AuthorName = "佚名"
		book.Status = 0
		book.Icon = "http://tingting-resource.bitekun.xin/resource/image/icon/hlm.jpg"
		fmt.Println("book.Name is",book.Name)
		//book.User_id = *(Params.MemberID)
		//db.Table("books").Create(&book)
	})

}

func OpenConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:root811123@tcp(106.14.2.153:3306)/tingting?charset=utf8&parseTime=True&loc=Local")
	return db,err
}