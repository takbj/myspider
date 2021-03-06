package spiderIns

import (
	"fmt"

	"github.com/takbj/myspider/3rd/go_spider/core/downloader"
	//	"bytes"
	"io/ioutil"
	"net/http"

	//	"github.com/PuerkitoBio/goquery"
	//	"github.com/bitly/go-simplejson"
	//    iconv "github.com/djimenez/iconv-go"
	"github.com/takbj/myspider/3rd/go_spider/core/common/mlog"
	"github.com/takbj/myspider/3rd/go_spider/core/common/page"
	"github.com/takbj/myspider/3rd/go_spider/core/common/request"
	//	"github.com/myspider/3rd/go_spider/core/common/util"
	//    "golang.org/x/text/encoding/simplifiedchinese"
	//    "golang.org/x/text/transform"
	/*	"io"
		"io/ioutil"
		"net/http"
		"net/url"
		//"fmt"
		"golang.org/x/net/html/charset"
		//    "regexp"
		//    "golang.org/x/net/html"
		"compress/gzip"
		"strings"*/)

type tDownloadEx struct {
	downloader.HttpDownloader
}

func (this *tDownloadEx) Download(req *request.Request) *page.Page {
	var mtype string
	var p = page.NewPage(req)
	mtype = req.GetResponceType()
	switch mtype {
	case "html", "json", "jsonp", "text":
		return this.HttpDownloader.Download(req)
	case "bin":
		return this.downloadBin(p, req)
	default:
		mlog.LogInst().LogError("error request type:" + mtype)
	}

	return p
}

func (this *tDownloadEx) downloadBin(p *page.Page, req *request.Request) *page.Page {
	resp, err := http.Get(req.Url)
	//	io.Copy(file, resp.Body)

	//	p, destbody := this.DownloadFile(p, req)
	//	if !p.IsSucc() {
	//		return p
	//	}
	if err != nil {
		fmt.Println("downloadBin  err=", err)
		return p
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	p.SetHeader(resp.Header)
	p.SetResp(resp)

	//	fmt.Println("downloadBin  req.Url=", req.Url, ",len(respBody)=", len(respBody))

	p.SetBodyStr(string(respBody)).SetStatus(false, "")
	return p
}
