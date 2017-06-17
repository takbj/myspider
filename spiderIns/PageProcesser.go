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
	//	urlDir := fmt.Sprintf("%s://%s%s", curUrl.Scheme, curUrl.Host, relativePath)

	var urls = []string{}
	cbFun := func(i int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		if !exist {
			href, _ = s.Attr("src")
		}
		//		fmt.Println("href=", href)
		hrefUrl, _ := url.Parse(href)
		if !hrefUrl.IsAbs() {
			hrefUrl.Host = curUrl.Host
			hrefUrl.Path = relativePath + hrefUrl.Path
			hrefUrl.Scheme = curUrl.Scheme
		}
		if ok, exist := allHosts[hrefUrl.Host]; !ok || !exist {
			return
		}
		if hrefUrl.Path == "" || hrefUrl.Path == "\\" || hrefUrl.Path == "/" {
			hrefUrl.Path = "/index.html"
		}

		urlTmp := hrefUrl.String()
		//		fmt.Println("href=", href, "urlTmp=", urlTmp)

		if _, exist := existUrls[urlTmp]; !exist {
			existUrls[urlTmp] = true
			urls = append(urls, urlTmp)
		}
	}

	query.Find("a").Each(cbFun)
	query.Find("link").Each(cbFun)
	query.Find("script").Each(cbFun)

	// these urls will be saved and crawed by other coroutines.
	fmt.Println("*MyPageProcesser.Process", urls)
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
	fmt.Println(pathTmp[:len(pathTmp)-1])

	filePath = path.Join(pathTmp[:len(pathTmp)-1]...)
	fmt.Println("filePath=", filePath, ",pathTmp[len(pathTmp)-1]=", pathTmp[len(pathTmp)-1])
	return filePath, pathTmp[len(pathTmp)-1]
}
