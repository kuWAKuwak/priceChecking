package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://backpack.tf/stats/"

func main() {
	mainMenu()
}

var (
	pageURL    string
	itemName   string
	prices     string
	itemURL    string
	comItemURL string
)

func mainMenu() string {
	var selQual int
	fmt.Println("쿠왘이의 시세 확인 프로그램입니다.")
	fmt.Println("아이템의 이름을 입력해주세요(영어로)")
	fmt.Scanln(&itemName)
	fmt.Println("검색하시려는 제품의 품질 번호는 무엇인가요?")
	fmt.Println("[1]진품 [2]골동품 [3]범상찮은 [4]유니크 [5]이상한 [6]귀신들린 [7]수집가")
	fmt.Scan(&selQual)

	switch selQual {
	case 1:
		pageURL = baseURL + "Genuine/" + strings.Title(itemName) + "/Tradable/"

	case 2:
		pageURL = baseURL + "Vintage/" + strings.Title(itemName) + "/Tradable/"

	case 3:
		pageURL = baseURL + "Unusual/" + strings.Title(itemName) + "/Tradable/"

	case 4:
		pageURL = baseURL + "Unique/" + strings.Title(itemName) + "/Tradable/"

	case 5:
		pageURL = baseURL + "Strange/" + strings.Title(itemName) + "/Tradable/"

	case 6:
		pageURL = baseURL + "Haunted/" + strings.Title(itemName) + "/Tradable/"

	case 7:
		pageURL = baseURL + "Collector/" + strings.Title(itemName) + "/Tradable/"

	}
	checkPrice(pageURL)
	return pageURL
}

func checkPrice(pageURL string) string {
	pageURL += "Craftable"
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)

	doc.Find(".tag.bottom-right").Each(func(i int, s *goquery.Selection) {
		prices = s.Find("span").Text()
		if strings.Contains(prices, "~") {
			prices = strings.Trim(prices, prices)
		}
		fmt.Print(prices)

	})

	return prices
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("요청실패 코드 :", res.StatusCode)
	}
}
