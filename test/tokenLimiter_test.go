package test

import (
	"fmt"
	"github/dpwgc/apilimiter"
	"testing"
	"time"
)

//令牌桶全局限流器测试
func TestTokenLimiter(t *testing.T) {

	fmt.Println("--- tokenLimiter ---")

	//设置令牌桶最大容量为100，每500毫秒生产50个令牌，相当于每1000毫秒最多只能取出100个令牌
	bucket := apilimiter.Bucket{
		Max:   100,
		Cycle: 500,
		Batch: 50,
	}

	//初始化令牌桶限流器
	bucket.NewTokenLimiter()

	//模拟200次请求
	for i := 0; i < 200; i++ {
		//每次访问至取出1个令牌
		isOk := bucket.GetToken(1)
		if isOk {
			fmt.Println(i, "Access successful", "[Time]:", time.Now().Unix())
		} else {
			fmt.Println(i, "Access failed.", "Token bucket is empty", "[Time]:", time.Now().Unix())
		}
	}

	time.Sleep(time.Second * 1)

	//模拟200次请求
	for i := 200; i < 400; i++ {
		//每次访问至取出1个令牌
		isOk := bucket.GetToken(1)
		if isOk {
			fmt.Println(i, "Access successful", "[Time]:", time.Now().Unix())
		} else {
			fmt.Println(i, "Access failed.", "Token bucket is empty", "[Time]:", time.Now().Unix())
		}
	}
}
