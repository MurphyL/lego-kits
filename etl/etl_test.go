package etl

import (
	"log"
	"strconv"
	"testing"
)

func TestRunJobNode(t *testing.T) {
	node := JobNode[int, string]{
		Read: func(ch chan string, cnt int) {
			for i := 0; i < cnt; i++ {
				ch <- strconv.FormatInt(int64(i), 10)
			}
			defer close(ch)
		},
		Write: func(ch chan string) {
			for line := range ch {
				log.Println("write:", line)
			}
		},
	}
	node.Run(10000)
}
