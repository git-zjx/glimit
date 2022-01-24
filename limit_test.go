package glimit

import (
	"fmt"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	g := New(3)
	for i := 0; i < 15; i++ {
		g.Add()
		go func(g *Limit, i int) {
			fmt.Println(i)
			time.Sleep(2 * time.Second)
			g.Done()
		}(g, i)
	}
	g.Wait()
}
