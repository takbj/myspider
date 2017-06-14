//
package main

/*
Packages must be imported:
    "core/common/page"
    "core/spider"
Pckages may be imported:
    "core/pipeline": scawler result persistent;
    "github.com/PuerkitoBio/goquery": html dom parser.
*/
import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	//	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/spider"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

var (
	host    = "ilxdh.com"
	xx      = "http://"
	urlHead = xx + host + "/"
)

// Parse html dom here and record the parse result that we want to Page.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	fmt.Println("*MyPageProcesser.Process.000")
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	url := p.GetRequest().GetUrl()
	query := p.GetHtmlParser()
	var urls []string
	query.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.Index(href, "://") >= 0 {
			return
		}

		index := strings.LastIndex(url, "/")
		if index == -1 {
			index = len(url)
		}
		path := url[0 : index+1]

		if _, exist := existUrls[urlHead+href]; !exist {
			existUrls[path+href] = true
			urls = append(urls, urlHead+href)
		}
	})
	// these urls will be saved and crawed by other coroutines.
	p.AddTargetRequests(urls, "html")

	//	url := p.GetRequest().GetUrl()
	//	index := strings.Index(url, urlHead)
	//	var path string
	//	if index >= 0 {
	//		path = url[len(urlHead):len(url)]
	//	}
	//	if path == "" {
	//		path = "index.html"
	//	}
	//	fmt.Println("*MyPageProcesser.Process.111,path=", path, ",urls=", urls)
}

func (this *MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

type MyPipeline struct {
}

func (this *MyPipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	println("----------------------------------------------------------------------------------------------")
	println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	println("Crawled result : ")
	for key, value := range items.GetAll() {
		println(key + "\t:\t" + value)
	}
	println("==============================================================================================")
}

var existUrls map[string]bool = map[string]bool{}

func main() {
	// Spider input:
	//  PageProcesser ;
	//  Task name used in Pipeline for record;
	spider.NewSpider(NewMyPageProcesser(), "TaskName").
		AddUrl(urlHead+"index.html", "html"). // Start url, html is the responce type ("html" or "json" or "jsonp" or "text")
		AddPipeline(&MyPipeline{}).           // Print result on screen
		SetThreadnum(3).                      // Crawl request by three Coroutines
		Run()
	existUrls[urlHead+"index.html"] = true
}
