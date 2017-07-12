package spiderIns

import (
	"github.com/takbj/myspider/3rd/go_spider/core/spider"
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

	for _, siteCfg := range config.C_SiteCfgs.Sites {
		if !siteCfg.SwitchOn {
			continue
		}
		spiderIns := spider.NewSpider(NewMyPageProcesser(siteCfg), "TaskName")
		for _, startUrl := range siteCfg.GetStartUrls() {
			existUrls[startUrl] = true
			spiderIns.AddUrl(startUrl, "html") // Start url, html is the responce type ("html" or "json" or "jsonp" or "text")
		}
		spiderIns.AddPipeline(&MyPipeline{}). // Print result on screen
							SetThreadnum(8). // Crawl request by three Coroutines
							SetDownloader(&tDownloadEx{}).
							Run()
	}

}
