package config

//	"fmt"
// 	"misc/mylog"
// 	"os"

var (
	C_SiteCfgs TSites
)

type TParseNode struct {
	AttrName string
	AttrType string
}

type TSiteCfg struct {
	Id              int  `json:"sites"`
	SwitchOn        bool `json:"switch_on"`
	StartUrls       []string
	DefaultFileName string
	HostList        []string
	SearchNodes     map[string]*TParseNode

	HostMaps map[string]bool
}

type TSites struct {
	Sites []*TSiteCfg `json:"sites"`
}

func (this *TSiteCfg) GetStartUrls() []string { //起始页面
	return this.StartUrls
}

func (this *TSiteCfg) GetDefaultFileName() string { //站点的默认索引文件名，ex: index.html
	return this.DefaultFileName
}

func (this *TSiteCfg) GetHostList() []string { //爬取的Host列表
	return this.HostList
}

func (this *TSiteCfg) CheckHost(host string) bool { //检查一个host是否在爬取的Host列表内
	ok, exist := this.HostMaps[host]
	return exist && ok
}

func (this *TSiteCfg) GetSearchNodes() map[string]*TParseNode { //获取需要爬取的节点,ex: map[string]string{"a":"href","link":"href","script":"src"}
	return this.SearchNodes
}

func (this *TSiteCfg) ForEachSearchNodes(param interface{}, cbFun func(nodeName string, attrName, attrType string, param interface{})) {
	for nodeName, nodeAttr := range this.SearchNodes {
		cbFun(nodeName, nodeAttr.AttrName, nodeAttr.AttrType, param)
	}
}

func (this *TSites) OnBeforeLoad() {

}

func (this *TSites) OnAfterLoad() {
	for _, siteCfg := range this.Sites {
		siteCfg.HostMaps = map[string]bool{}
		for _, host := range siteCfg.HostList {
			siteCfg.HostMaps[host] = true
		}
	}
}

func init() {
	registerCfg("site", "config/site_cfg.json", &C_SiteCfgs)
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
