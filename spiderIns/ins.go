package spiderIns

import (
	"github.com/hu17889/go_spider/core/spider"
	"github.com/takbj/myspider/config"
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
	//	url := cstUrlHead + "index.html"
	//	existUrls[url] = true

	spiderIns := spider.NewSpider(NewMyPageProcesser(&config.C_SiteCfg), "TaskName")
	for _, startUrl := range config.C_SiteCfg.GetStartUrls() {
		existUrls[startUrl] = true
		spiderIns.AddUrl(startUrl, "html") // Start url, html is the responce type ("html" or "json" or "jsonp" or "text")
	}
	spiderIns.AddPipeline(&MyPipeline{}). // Print result on screen
						SetThreadnum(8). // Crawl request by three Coroutines
						Run()

}
