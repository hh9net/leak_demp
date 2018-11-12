package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/CrowdStrike/ratelimiter"
)

func main() {
	runtime.GOMAXPROCS(4)
	maxCapacity := 1000
	for i := 1; i < 5000; i++ {
		ratePeriod := 1 * time.Minute
		rl, err := ratelimiter.New(maxCapacity, ratePeriod)
		if err != nil {
			fmt.Printf("Unable to create cache")
		}
		userKey2 := strconv.Itoa(i)
		maxCount2 := 40 // the maximum number of items I want from this user in one hour
		uu := i
		fmt.Println("================", uu, userKey2)
		go func(userKey string, maxCount int, k int) {
			for {
				cnt, underRateLimit := rl.Incr(userKey, maxCount)
				if underRateLimit {
					// allow further access
					fmt.Println(k, userKey, "keyi", cnt)
				} else {
					fmt.Printf("%d   User [%s] is over rate limit, denying for now, current count [%d]\n", k, userKey, cnt)
				}
				//				time.Sleep(time.Second / 1000)
			}
		}(userKey2, maxCount2, uu)
	}
	for {
		time.Sleep(time.Second)
	}
}
