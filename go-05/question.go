// 参考 Hystrix 实现一个滑动窗口计数器。
// Reference Hystrix-go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}

	return r
}

func (r *Number) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool
	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}

	return bucket
}

func (r *Number) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *Number) Increment(i float64) {
	if i == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	b.Value += i
	r.removeOldBuckets()
}

func (r *Number) UpdateMax(n float64) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	r.removeOldBuckets()
}

func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			sum += bucket.Value
		}
	}

	return sum
}

func (r *Number) Max(now time.Time) float64 {
	var max float64

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

func (r *Number) Avg(now time.Time) float64 {
	return r.Sum(now) / 10
}

func main() {
	r := NewNumber()

	now := time.Now()
	nowN := now.Unix()

	r.Increment(2)
	if v, ok := r.Buckets[nowN]; ok {
		fmt.Println(v.Value)
	} else {
		fmt.Println("Increment null")
	}

	r.UpdateMax(4)
	if v, ok := r.Buckets[nowN]; ok {
		fmt.Println(v.Value)
	} else {
		fmt.Println("UpdateMax null")
	}

	r.Sum(now)
	if v, ok := r.Buckets[nowN]; ok {
		fmt.Println(v.Value)
	} else {
		fmt.Println("Sum null")
	}

	r.Max(now)
	if v, ok := r.Buckets[nowN]; ok {
		fmt.Println(v.Value)
	} else {
		fmt.Println("Max null")
	}

	r.Avg(now)
	if v, ok := r.Buckets[nowN]; ok {
		fmt.Println(v.Value)
	} else {
		fmt.Println("Avg null")
	}
}
