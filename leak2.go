package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail/ratelimiter"
)

func main() {

	bucket := ratelimiter.NewLeakyBucket(10, time.Second*60)
	bucket.Lastupdate = time.Now()
	var i = 0
	for {
		i++

		bucket.Now = func() time.Time { return time.Now() }

		if bucket.Pour(1) {
			fmt.Println(time.Now(), "Expected true", i)
		} else {
			fmt.Println(time.Now(), "Expected false", i)

		}
		time.Sleep(time.Second)
	}

}
