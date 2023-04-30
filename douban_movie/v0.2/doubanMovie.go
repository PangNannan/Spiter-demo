package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql" //这个必须要有
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// 定义数据库连接参数
const (
	USERNAME = "root"
	PASSWORD = "yiyy2011code"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "douban_movie"
)

var DB *sql.DB

// 映射数据
type movieData struct {
	title    string `json:"title"`
	imgSrc   string `json:"imgSrc"`
	director string `json:"director"`
	act      string `json:"act"`
	year     string `json:"year"`
	score    string `json:"score"`
	quote    string `json:"quote"`
}

func main() {

	InitDB()

	//爬取每一页
	for i := 0; i < 10; i++ {
		fmt.Printf("正在查询第 %d 页 \n", i+1)
		Spiter(strconv.Itoa(i * 25))
	}

}

func Spiter(page string) {
	// 创建一个客户端
	client := http.Client{}

	//2发送请求
	req, err := http.NewRequest("GET", "https://movie.douban.com/top250"+"?start="+page, nil)
	if err != nil {
		fmt.Println("")
		return
	}

	//增加请求头
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/chart")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resq, err := client.Do(req) //请求
	if err != nil {
		fmt.Println("请求失败", err)
		return
	}
	//3 解析网页
	doc, err := goquery.NewDocumentFromReader(resq.Body)
	if err != nil {
		fmt.Println("解析失败：", err)
		return
	}

	//4获取节点
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.hd > a > span:nth-child(1)
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.bd > p:nth-child(1)
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.bd > div > span.rating_num
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.bd > p.quote > span
	//#content > div > div.article > ol > li:nth-child(1) > div > div.pic > a > img

	doc.Find("#content > div > div.article > ol > li").Each( //返回一个列表
		func(i int, selection *goquery.Selection) { //在列表里面继续找
			tital := selection.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
			img := selection.Find("div > div.pic > a > img")
			imgSrc, ok := img.Attr("src")

			info := selection.Find("div.info > div.bd > p:nth-child(1)").Text()
			score := selection.Find("div > div.info > div.bd > div > span.rating_num").Text()
			quote := selection.Find("div.bd > p.quote > span").Text()

			var movieData movieData
			//打印
			if ok {
				director, act, year := InfoSpite(info)
				movieData.title = tital
				movieData.imgSrc = imgSrc
				movieData.director = director
				movieData.act = act
				movieData.year = year
				movieData.score = score
				movieData.quote = quote

				//插入数据
				ok := InsertSql(movieData)
				if ok {
					fmt.Println("插入成功！")

				} else {
					fmt.Println("插入失败！")
				}
				//fmt.Println(movieData)
			}
		}) //css选择器语法

	//5 保存数据

}

func InfoSpite(info string) (director, act, year string) {
	directorRe, _ := regexp.Compile(`导演:(.*)主演:`) //匹配年分
	director = string(directorRe.Find([]byte(info)))

	actRe, _ := regexp.Compile(`主演:(.*)`) //匹配年分
	act = string(actRe.Find([]byte(info)))

	yearRe, _ := regexp.Compile(`(\d+)`) //匹配年分
	year = string(yearRe.Find([]byte(info)))

	return
}

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertSql(movieData movieData) bool {
	//1开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx err:", err)
		return false
	}
	//2插入数据
	stmt, err := tx.Prepare("Insert into movie_data(title,imgSrc,director,act,year,score,quote) values (?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("stmt err : ", err)
		return false
	}

	_, err1 := stmt.Exec(movieData.title,
		movieData.imgSrc,
		movieData.director,
		movieData.act,
		movieData.year,
		movieData.score,
		movieData.quote)
	if err1 != nil {
		fmt.Println("exec err :", err)
		return false
	}

	//提交事务
	_ = tx.Commit()

	return true
}
