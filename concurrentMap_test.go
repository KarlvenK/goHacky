package goHacky

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var m = NewConcurrentMap()
var wg = sync.WaitGroup{}
var ans []int

func TestConcurrentMap(t *testing.T) {
	wg.Add(2)
	go testconcurrentmapGet(t)
	go testconcurrentmapPut(t)
	wg.Wait()
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		v, _ := m.Get(i, time.Millisecond*100)
		if v.(int) != ans[i] {
			t.Errorf(".....shit")
			fmt.Println(i, ":  ", ans[i], " ", v)
		}
	}
	fmt.Println(m.data)
	fmt.Println(ans)
}

func testconcurrentmapGet(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Put(i, r.Int())
		}(i)
	}
	wg.Done()
}

func testconcurrentmapPut(t *testing.T) {
	ans = make([]int, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			if v, ok := m.Get(i, 10*time.Millisecond); ok {
				ans[i] = v.(int)
			} else {
				t.Log("shit.........", i)
			}
		}(i)
	}
	wg.Done()
}
