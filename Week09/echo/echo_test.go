package echo

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

const addr = ":8080"

func TestEcho(t *testing.T) {
	c := NewClient(addr)
	ctx, cancel := context.WithTimeout(context.Background(), 1 *time.Second)

	err := c.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = c.Stop(ctx)
		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err := c.Send(fmt.Sprintf("i am msg: %d\n", i))
			if err != nil {
				t.Fatal(err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <- c.Recv():
				t.Logf("recv msg: %s", msg)
			case <- ctx.Done():
				return
			}
		}
	}()
	wg.Wait()
}