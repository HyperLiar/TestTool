package main

import (
	"fmt"
	"sync"
)

// sync.map 通过read,dirty分别作为读 写map
// 延迟删除，删除一个key只是在dirty打标记，提升dirty为read时，才清理数据
// 优先从read读取，更新，删除，对read读取不需要锁
// 双检查，即，如果不在read内，增加miss次数，改为从dirty获取
func main() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")

	fmt.Println(m.Len())
}
