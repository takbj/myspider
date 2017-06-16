package spiderIns

import (
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page_items"
)

type MyPipeline struct {
}

func (this *MyPipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	// println("----------------------------------------------------------------------------------------------")
	// println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	// //	println("",items.)
	// println("Crawled result : ")
	// for key, value := range items.GetAll() {
	// 	println(key + "\t:\t" + value)
	// }
	// println("==============================================================================================")
}
