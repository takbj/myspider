package spiderIns

import (
	"fmt"

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

var existUrls map[string]string = map[string]string{}

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
			existUrls[startUrl] = ""
			spiderIns.AddUrl(startUrl, "html") // Start url, html is the responce type ("html" or "json" or "jsonp" or "text")
		}
		spiderIns.AddPipeline(&MyPipeline{}). // Print result on screen
							SetThreadnum(8). // Crawl request by three Coroutines
							SetDownloader(&tDownloadEx{}).
							Run()

		if config.C_GlobalCfg.DebugFlag {
			fmt.Println("-----------------------------------------------------------------")
			for k, parent := range existUrls {
				fmt.Println(k, "\n\t", parent)
			}
		}
	}

}
