package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex ป้อนกันการเปลี่ยแปลงตัวแปรที่แชร์กันระหว่าง routine(thred) ที่มันแชร์กันอยู่อย่างไม่สมบูรณ์ lock/unlock
func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increase("somekey") //1000 routine
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Increase(key string) {
	c.mux.Lock()
	//Lock ให้ routine เดียวเท่านั้นที่เข้าถึงได้
	c.v[key]++
	c.mux.Unlock()
}

// อ่านค่าก่อนแก้ไข
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
