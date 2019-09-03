package downloader

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func Downlaod(baseDir string, articleUrl string) error {
	resp, err := soup.Get(articleUrl)
	if err != nil {
		return fmt.Errorf("please make sure the article URL is correct. %v", articleUrl)
	}

	doc := soup.HTMLParse(resp)
	title := doc.Find("h2", "id", "activity-name").Text()
	title = strings.TrimSpace(title)
	resourceDir := baseDir + string(filepath.Separator) + title

	log.Println("创建资源目录: " + resourceDir)
	createResourceDir(resourceDir)

	log.Println("保存视频地址...")
	handleVideoResource(resourceDir, doc)

	log.Println("下载图片...")
	handleImageResource(resourceDir, doc)

	return nil
}

func createResourceDir(d string) {
	_ = os.Mkdir(d, 0777)
	_ = os.Mkdir(d+"/images", 0777)
}

func handleVideoResource(rd string, doc soup.Root) {
	videos := doc.FindAll("iframe", "class", "video_iframe")
	if len(videos) > 0 {
		f, _ := os.Create(rd + "/视频地址.txt")
		videoUrls := make([]string, 0, len(videos))
		for _, video := range videos {
			dataSrc := video.Attrs()["data-src"]
			if dataSrc == "" {
				continue
			}
			videoUrls = append(videoUrls, dataSrc)
		}

		for _, u := range videoUrls {
			_, _ = f.WriteString(u + "\n")
		}
		defer f.Close()
	}
}

func handleImageResource(rd string, doc soup.Root) {
	var wg sync.WaitGroup

	imgs := doc.FindAll("img")
	imgUrls := make([]string, 0, len(imgs))
	for _, img := range imgs {
		dataSrc := img.Attrs()["data-src"]
		if dataSrc == "" {
			continue
		}
		wg.Add(1)
		imgUrls = append(imgUrls, dataSrc)
	}

	for i, u := range imgUrls {
		go func(rd string, filename string, imgUrl string) {
			defer wg.Done()
			u, err := url.Parse(imgUrl)
			if err != nil {
				panic(err)
			}
			fileType := u.Query().Get("wx_fmt")
			if fileType == "" {
				fileType = "jpg"
			}

			resp, err := http.Get(imgUrl)
			if err != nil {
				panic(err)
			}

			imgFile, err := os.Create(rd + "/images/" + filename + "." + fileType)
			if err != nil {
				panic(err)
			}
			d, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				panic(err)
			}

			_, err = imgFile.Write(d)
			if err != nil {
				panic(err)
			}
			defer imgFile.Close()
		}(rd, strconv.Itoa(i+1), u)
	}
	wg.Wait()
}
