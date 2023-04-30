package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	gorm 步骤
 	1，创建映射模型
	2. 连接数据库、池
	3,	 创建表（做一次就行不要重复做
	4 增删改查
*/
// 定义数据库连接参数
const (
	USERNAME = "root"
	PASSWORD = "yiyy2011code"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "douban_movie"
)

var db *gorm.DB

// 数据模型
type MovieData struct {
	Title    string
	ImgSrc   string `gorm:"column:img_src" `
	Director string
	Act      string
	Year     string
	Score    string
	Quote    string

	gorm.Model //模型
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

			var movieData MovieData
			//打印
			if ok {
				director, act, year := InfoSpite(info)
				movieData.Title = tital
				movieData.ImgSrc = imgSrc
				movieData.Director = director
				movieData.Act = act
				movieData.Year = year
				movieData.Score = score
				movieData.Quote = quote
				movieData.CreatedAt = time.Now()
				//插入数据
				ok := InsertSql(&movieData)
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

	var err error
	//连接数据库
	dns := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}))
	if err != nil {
		panic("failed to connect database")
	}
	// 数据库连接池设置
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("success to link mysql")

	//创建表结构
	/*err = db.Table("movie_data2").AutoMigrate(&MovieData{})
	if err != nil {
		panic("failed to create table")
	} else {
		fmt.Println("success to create table")
	}*/

}

func InsertSql(movieData *MovieData) bool { //返回true 插入成功
	//1开启事务 返回任何错误都会回滚事务
	err := db.Transaction(func(tx *gorm.DB) error {
		//2插入语句

		if err := tx.Table("movie_data2").Create(movieData).Error; err != nil {
			return err
		}
		return nil //提交事务
	})

	if err != nil {
		return false
	}
	return true
}
