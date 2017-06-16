package spiderIns

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

func (this *MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

//func queryCBFun(i int, s *goquery.Selection) {
//	href, _ := s.Attr("href")
//	if strings.Index(href, "://") >= 0 {
//		return
//	}

//	index := strings.LastIndex(url, "/")
//	if index == -1 {
//		index = len(url)
//	}
//	path := url[0 : index+1]

//	if _, exist := existUrls[cstUrlHead+href]; !exist {
//		existUrls[path+href] = true
//		urls = append(urls, cstUrlHead+href)
//	}
//}

// Parse html dom here and record the parse result that we want to Page.
// Package goquery (http://godoc.org/github50.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	//fmt.Println("*MyPageProcesser.Process.000")
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	urlStr := p.GetRequest().GetUrl()
	curUrl, _ := url.Parse(urlStr)

	query := p.GetHtmlParser()
	save(curUrl, p.GetBodyStr())

	lastSepIndex := strings.LastIndex(curUrl.Path, "/")
	relativePath := curUrl.Path[:lastSepIndex+1]
	urlPath := fmt.Sprintf("%s://%s%s", curUrl.Scheme, curUrl.Host, relativePath)

	var urls []string
	var urlTmp string
	query.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		hrefUrl, _ := url.Parse(href)
		if ok, exist := allHosts[hrefUrl.Host]; !ok || !exist {
			return
		}

		urlTmp = hrefUrl.String()

		if _, exist := existUrls[urlTmp]; !exist {
			existUrls[urlPath+href] = true
			urls = append(urls, urlTmp)
		}
	})

	// these urls will be saved and crawed by other coroutines.
	p.AddTargetRequests(urls, "html")
}

func save(curUrl *url.URL, bodyString string) {
	if len(bodyString) <= 0 {
		return
	}

	filePath, fileName := getFilePath(curUrl)

	os.MkdirAll(filePath, os.ModeDir)
	absPath := path.Join(filePath, fileName)
	file, err := os.Create(absPath)
	if err != nil {
		println("out file[", absPath, "] err:", err)
		return
	}
	defer file.Close()

	file.WriteString(bodyString)
}

func getFilePath(curUrl *url.URL) (filePath, fileName string) {
	pathTmp := []string{cstSaveRootPath, curUrl.Host}

	pathTmp = append(pathTmp, strings.Split(curUrl.Path, "/")...)

	filePath = path.Join(pathTmp[:len(pathTmp)-1]...)
	return filePath, pathTmp[len(pathTmp)-1]
}
