package main

import (
	"flag"
	"github.com/jwma/wechat_article_res_downloader/pkg/downloader"
	"log"
)

func main() {
	var baseDir, articleUrl string
	flag.StringVar(&baseDir, "folder", "", "资源存放目录")
	flag.StringVar(&articleUrl, "url", "", "微信文章地址")
	flag.Parse()

	if baseDir == "" {
		log.Fatal("请传入资源存放目录")
	}
	if articleUrl == "" {
		log.Fatal("请传入微信文章地址")
	}
	err := downloader.Download(baseDir, articleUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("资源下载成功")
}
