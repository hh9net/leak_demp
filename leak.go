package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail/ratelimiter"
)

func main() {

	bucket := ratelimiter.NewLeakyBucket(60, time.Second)
	bucket.Lastupdate = time.Unix(0, 0)

	bucket.Now = func() time.Time { return time.Unix(1, 0) }

	if bucket.Pour(61) {
		fmt.Println("---------------", 61, "Expected false")
	}

	if bucket.Pour(10) {
		fmt.Println(10, "Expected true")
	}

	if bucket.Pour(49) {
		fmt.Println(49, "Expected true")
	}

	if bucket.Pour(2) {
		fmt.Println(2, "Expected false")
	}

	bucket.Now = func() time.Time { return time.Unix(61, 0) }
	if bucket.Pour(60) {
		fmt.Println(60, "Expected true")
	}

	if bucket.Pour(1) {
		fmt.Println(1, "Expected false")
	}

	bucket.Now = func() time.Time { return time.Unix(70, 0) }

	if bucket.Pour(1) {
		fmt.Println(1, "Expected true")
	}
	if bucket.Pour(8) {
		fmt.Println(10, "Expected true")
	}
}
