package main

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestTaskControl(t *testing.T) {
	dataChan := make(chan int)

	taskNum := 3

	wg := sync.WaitGroup{}
	wg.Add(taskNum)

	// 起多个协程，data关闭时退出
	for i := 0; i < taskNum; i++ {
		go func(taskNo int) {
			defer wg.Done()
			t.Logf("Task %d run\n", taskNo)

			for {
				select {
				case _, ok := <-dataChan:
					if !ok {
						t.Logf("Task %d notify to stop\n", taskNo)
						return
					}
				}
			}
		}(i)
	}

	// 通知退出
	go func() {
		time.Sleep(3 * time.Second)
		close(dataChan)
	}()

	// 等待退出完成
	wg.Wait()
}

func TestTcp(t *testing.T) {
	_, err := net.Dial("tcp", "127.0.0.1:50555")
	if err != nil {
		t.Logf("conn error %s", err)
		return
	}
}

func TestUdp(t *testing.T) {
	conn, err := net.Dial("udp", ":50555")
	if err != nil {
		return
	}
	defer conn.Close()
	conn.Write([]byte("close"))
}

func TestConfig(t *testing.T) {
	multiport.LoadFromYaml()
}
