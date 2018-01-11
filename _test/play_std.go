package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mediocregopher/radix.v2/pool"
)

func main() {
	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	p, err := pool.New("tcp", "localhost:6379", 8)
	check(err)

	start := time.Now()
	var wg sync.WaitGroup
	N, K := 800, 8000
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			for j := 0; j < K; j++ {
				conn, err := p.Get()
				check(err)
				res, err := conn.Cmd("GET", "asdf").Str()
				check(err)
				p.Put(conn)

				//if !bytes.Equal(fut.Result.([]byte), []byte{49}) {
				if res != "1" {
					panic("mismatch")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%d*%d: %s\n", N, K, time.Now().Sub(start))
}