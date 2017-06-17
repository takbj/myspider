package config

// import (
// 	"fmt"
// 	"misc/mylog"
// 	"os"
// )

var (
	C_SiteCfg SiteCfg
)

type SiteCfg struct {
	StartUrl        string
	DefaultFileName string
	HostList        []string
	SearchNodes     map[string]string

	hostMaps map[string]bool
}

func (this *SiteCfg) GetStartUrl() string { //起始页面
	return this.StartUrl
}

func (this *SiteCfg) GetDefaultFileName() string { //站点的默认索引文件名，ex: index.html
	return this.DefaultFileName
}

func (this *SiteCfg) GetHostList() []string { //爬取的Host列表
	return this.HostList
}

func (this *SiteCfg) CheckHost(host string) bool { //检查一个host是否在爬取的Host列表内
	ok, exist := this.hostMaps[host]
	return exist && ok
}

func (this *SiteCfg) GetSearchNodes() map[string]string { //获取需要爬取的节点,ex: map[string]string{"a":"href","link":"href","script":"src"}
	return this.SearchNodes
}

func (this *SiteCfg) OnBeforeLoad() {

}

func (this *SiteCfg) OnAfterLoad() {
	this.hostMaps = map[string]bool{}
	for _, host := range this.HostList {
		this.hostMaps[host] = true
	}
}

func init() {
	registerCfg("site", "config/site_cfg.json", &C_SiteCfg)
}

// func init() {
// 	tmpCfg := ExCfg{}
// 	err := ReloadCfg("config/excfg.json", &tmpCfg)
// 	if err != nil {
// 		os.Exit(-1)
// 		mylog.Error("server start error:ExCfg json file read failed,", err)
// 		fmt.Println("server start error:ExCfg json file read failed,", err)
// 	}
// 	C_ExCfg = tmpCfg
// }
