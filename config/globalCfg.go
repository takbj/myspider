package config

// import (
// 	"fmt"
// 	"misc/mylog"
// 	"os"
// )

var (
	C_GlobalCfg TGlobalCfg
)

type TGlobalCfg struct {
	DebugFlag bool
}

func (this *TGlobalCfg) OnBeforeLoad() {

}

func (this *TGlobalCfg) OnAfterLoad() {
}

func init() {
	registerCfg("global", "config/global_cfg.json", &C_GlobalCfg)
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
