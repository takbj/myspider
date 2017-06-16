package spiderIns

import (
	"github.com/hu17889/go_spider/core/spider"
)

const (
	cstHost    string = "ilxdh.com"
	cstShecm   string = "http"
	cstUrlHead string = cstShecm + "://" + cstHost + "/"

	// cstSaveRootPath string = "E:/pro/go_ex/src/github.com/takbj/myspider/xx/"
	cstSaveRootPath string = "E:/zsy/spider_out/"
)

var existUrls map[string]bool = map[string]bool{}

func Run() {
	// Spider input:
	//  PageProcesser ;
	//  Task name used in Pipeline for record;
	//
	url := cstUrlHead + "index.html"
	existUrls[url] = true

	spider.NewSpider(NewMyPageProcesser(), "TaskName").
		AddUrl(url, "html").        // Start url, html is the responce type ("html" or "json" or "jsonp" or "text")
		AddPipeline(&MyPipeline{}). // Print result on screen
		SetThreadnum(3).            // Crawl request by three Coroutines
		Run()
}
