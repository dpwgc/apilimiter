package test

import (
	"fmt"
	"github/dpwgc/apilimiter"
	"testing"
	"time"
)

//滑动窗口地址限流器测试
func TestAddrLimiter(t *testing.T) {

	fmt.Println("--- addrLimiter ---")

	//设置单个用户每1000毫秒内最多访问5次
	period := apilimiter.Period{
		Max:  5,
		Size: 1000,
	}

	//初始化地址限流器
	period.NewAddrLimiter()

	//模拟100个请求
	for i := 0; i < 100; i++ {
		//用户10.110.2.1尝试获取访问许可
		isOk := period.GetPermit("10.110.2.1")
		if isOk {
			fmt.Println("10.110.2.1", "Access successful", "[Time]:", time.Now().Unix())
		} else {
			fmt.Println("10.110.2.1", "Access failed", "[Time]:", time.Now().Unix())
		}
	}

	//模拟100个请求
	for i := 0; i < 100; i++ {
		//用户10.110.2.2尝试获取访问许可
		isOk := period.GetPermit("10.110.2.2")
		if isOk {
			fmt.Println("10.110.2.2", "Access successful", "[Time]:", time.Now().Unix())
		} else {
			fmt.Println("10.110.2.2", "Access failed", "[Time]:", time.Now().Unix())
		}
	}

	time.Sleep(time.Second * 1)

	//模拟100个请求
	for i := 0; i < 100; i++ {
		//用户10.110.2.1尝试获取访问许可
		isOk := period.GetPermit("10.110.2.1")
		if isOk {
			fmt.Println("10.110.2.1", "Access successful", "[Time]:", time.Now().Unix())
		} else {
			fmt.Println("10.110.2.1", "Access failed", "[Time]:", time.Now().Unix())
		}
	}

	//模拟100个请求
	for i := 0; i < 100; i++ {
		//用户10.110.2.2尝试获取访问许可
		isOk := period.GetPermit("10.110.2.2")
		if isOk {
			fmt.Println("10.110.2.2", "Access successful", "[Time]:", time.Now().Unix())
		} else {
			fmt.Println("10.110.2.2", "Access failed", "[Time]:", time.Now().Unix())
		}
	}
}
