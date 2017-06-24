//
package resource_manage_test

import (
	"testing"

	"github.com/takbj/myspider/3rd/go_spider/core/common/resource_manage"
)

func TestResourceManage(t *testing.T) {
	var mc *resource_manage.ResourceManageChan
	mc = resource_manage.NewResourceManageChan(1)
	mc.GetOne()
	println("incr")
	mc.FreeOne()
	println("decr")
	mc.GetOne()
	println("incr")
}
