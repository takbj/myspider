package spiderIns

import (
	"fmt"
	//	"io"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/takbj/myspider/config"
)

type MyPageProcesser struct {
	configer interface{} //SiteCfg
}

func NewMyPageProcesser(configerIn interface{}) *MyPageProcesser {
	return &MyPageProcesser{configer: configerIn}
}

func (this *MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

type SiteCfg interface {
	GetStartUrls() []string                      //起始页面
	GetDefaultFileName() string                  //站点的默认索引文件名，ex: index.html
	GetHostList() []string                       //爬取的Host列表
	CheckHost(host string) bool                  //检查一个host是否在爬取的Host列表内
	GetSearchNodes() map[string]*config.TSiteUrl //获取需要爬取的节点,ex: map[string]string{"a":"href","link":"href","script":"src"}
}

// Parse html dom here and record the parse result that we want to Page.
// Package goquery (http://godoc.org/github50.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		fmt.Println(p.Errormsg())
		return
	}

	urlStr := p.GetRequest().GetUrl()
	curUrl, _ := url.Parse(urlStr)

	query := p.GetHtmlParser()
	if !save(curUrl, p.GetBodyStr()) {
		if config.C_GlobalCfg.DebugFlag {
			fmt.Println("\n\n\nsave \"", urlStr, "\" failed!")
		}
	}

	lastSepIndex := strings.LastIndex(curUrl.Path, "/")
	relativePath := curUrl.Path[:lastSepIndex+1]
	var urls = map[string]string{}
	var herfAttrName *config.TSiteUrl
	cbFun := func(i int, s *goquery.Selection) {
		href, exist := s.Attr(herfAttrName.AttrName)
		if !exist {
			return
		}

		hrefUrl, _ := url.Parse(strings.Trim(href, " "))
		if !hrefUrl.IsAbs() {
			hrefUrl.Host = curUrl.Host
			hrefUrl.Path = relativePath + hrefUrl.Path
			hrefUrl.Scheme = curUrl.Scheme
		}
		if !this.configer.(SiteCfg).CheckHost(hrefUrl.Host) {
			return
		}
		if hrefUrl.Path == "" || hrefUrl.Path == "\\" || hrefUrl.Path == "/" {
			hrefUrl.Path = this.configer.(SiteCfg).GetDefaultFileName()
		}

		urlTmp := hrefUrl.String()

		if config.C_GlobalCfg.DebugFlag {
			if strings.Index(urlTmp, "%") >= 0 {
				fmt.Println("\n", urlTmp, href)
			}
		}

		if _, exist := existUrls[urlTmp]; !exist {
			existUrls[urlTmp] = true
			urls[urlTmp] = herfAttrName.AttrType
			//			urls = append(urls, urlTmp)
		}
	}

	for nodeName, nodeAttr := range this.configer.(SiteCfg).GetSearchNodes() {
		herfAttrName = nodeAttr //"href","src"
		query.Find(nodeName).Each(cbFun)
	}

	// these urls will be saved and crawed by other coroutines.
	//	fmt.Println("*MyPageProcesser.Process", urls)
	for url, respType := range urls {
		p.AddTargetRequest(url, respType)
	}
}

func save(curUrl *url.URL, bodyString string) bool {
	if len(bodyString) <= 0 {
		if config.C_GlobalCfg.DebugFlag {
			fmt.Println("save.000000")
		}
		return false
	}
	filePath, fileName := getFilePath(curUrl)

	os.MkdirAll(filePath, os.ModeDir)
	absPath := path.Join(filePath, fileName)
	file, err := os.Create(absPath)
	if err != nil {
		fmt.Println("save.111  out file[", absPath, "] err:", err)
		return false
	}
	defer file.Close()

	//	io.Copy(file, []byte(bodyString))
	file.Write([]byte(bodyString))

	return true
}

func getFilePath(curUrl *url.URL) (filePath, fileName string) {
	pathTmp := []string{cstSaveRootPath, curUrl.Host}

	pathTmp = append(pathTmp, strings.Split(curUrl.Path, "/")...)
	//	fmt.Println(pathTmp[:len(pathTmp)-1])

	filePath = path.Join(pathTmp[:len(pathTmp)-1]...)
	return filePath, pathTmp[len(pathTmp)-1]
}
