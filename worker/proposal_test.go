/*
 * Copyright 2017-2019 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package worker

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"sync/atomic"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func PropposeAndWaitDup() error {
// 	// some time succed/fail
// 	propose := func(timeout time.Duration) error {
// 		num := int(rand.Int31n(10))
// 		if num%2 == 0 {
// 			return errInternalRetry
// 		}

// 		time.Sleep(1 * time.Second)

// 		return nil
// 	}

// 	runPropose := func(i int) error {
// 		// Each retry creates a new proposal, which adds to the number of pending proposals. We
// 		// should consider this into account, when adding new proposals to the system.
// 		if err := limiter.incr(context.Background(), i); err != nil {
// 			return err
// 		}
// 		defer limiter.decr(i)
// 		return propose(newTimeout(i))
// 	}

// 	for i := 0; i < 3; i++ {
// 		if err := runPropose(i); err != errInternalRetry {
// 			return err
// 		}
// 	}
// 	return errUnableToServe
// }

// var total, pending, completed, aborted int64

// func TestLimiter(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano())
// 	limiter = rateLimiter{c: sync.NewCond(&sync.Mutex{}), max: 256}
// 	go limiter.c.Broadcast()

// 	go func() {
// 		now := time.Now()
// 		for range time.Tick(1 * time.Second) {
// 			fmt.Println("Seconds elapsed :", int64(time.Since(now).Seconds()),
// 				"Total proposals: ", atomic.LoadInt64(&total),
// 				"Pending proposal: ", atomic.LoadInt64(&pending),
// 				"Completed Proposals: ", atomic.LoadInt64(&completed),
// 				"Aboted Proposals: ", atomic.LoadInt64(&aborted),
// 				"IOU: ", limiter.iou)
// 		}
// 	}()

// 	var wg sync.WaitGroup
// 	for i := 0; i < 500; i++ {
// 		wg.Add(1)
// 		go func(no int) {
// 			defer wg.Done()

// 			for count := 0; ; count++ {
// 				atomic.AddInt64(&total, 1)
// 				atomic.AddInt64(&pending, 1)
// 				if err := PropposeAndWaitDup(); err != nil {
// 					// fmt.Println("Got error while propose and wait: ", err)
// 					atomic.AddInt64(&aborted, 1)
// 				} else {
// 					atomic.AddInt64(&completed, 1)
// 				}
// 				// fmt.Println("Routine: ", no, " Iteration ", count, "propose and wait successful")
// 				atomic.AddInt64(&pending, -1)
// 				// time.Sleep(100 * time.Millisecond)
// 			}
// 		}(i)
// 	}

// 	wg.Wait()
// }
