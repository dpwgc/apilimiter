# API Limiter

## Go接口限流工具包

***

### 简介

#### apilimiter是一个简易接口限流工具包，其包含两种限流方式。

#### （1）基于令牌桶限流算法+CAS实现的毫秒级全局限流器 `tokenLimiter`

#### （2）基于滑动窗口限流算法实现的毫秒级IP访问限流器 `addrLimiter`

***

### 导入工具

* 引入包：`go get github.com/dpwgc/apilimiter`

```
import (
    "github/dpwgc/apilimiter"
)
```

***

### 使用方法

#### 令牌桶全局流量限制 `tokenLimiter`

* 函数说明

```
// Bucket 令牌桶配置
type Bucket struct {
    Max   int64 	//令牌桶的最大存储上限
    Cycle int64 	//生成一批令牌的周期（每{cycle}毫秒生产一批令牌）
    Batch int64 	//每批令牌的数量
}

// NewTokenLimiter 初始化令牌桶全局限流器
func (bucket *Bucket) NewTokenLimiter()

// GetToken 获取令牌 @num:本次请求需要拿取的令牌数
func (bucket *Bucket) GetToken(num int64) bool 
```

* 初始化令牌桶全局限流器

```
//设置令牌桶最大容量为100，每500毫秒生产50个令牌，相当于每1000毫秒最多只能取出100个令牌
bucket := apilimiter.Bucket{
    Max:   100,
    Cycle: 500,
    Batch: 50,
}

//初始化令牌桶限流器
bucket.NewTokenLimiter()
```

* 获取令牌

```
//取出1个令牌
isOk := bucket.GetToken(1)
if isOk {
    // 获取成功
    fmt.Println("Access successful")
} else {
    // 获取失败 
    fmt.Println("Access failed.", "Token bucket is empty")
}
```

***

#### 滑动窗口地址限流器 `addrLimiter`

* 函数说明

```
// Period 滑动窗口配置
type Period struct {
    Max  int64 		//单个用户在滑动窗口内的最大访问次数
    Size int64 		//滑动窗口大小（单位：毫秒）
}

// NewAddrLimiter 初始化滑动窗口地址限流器
func (period *Period) NewAddrLimiter()

// GetPermit 用户申请访问
func (period *Period) GetPermit(addr string) bool
```

* 初始化滑动窗口地址限流器

```
//设置单个用户每1000毫秒内最多访问5次
period := apilimiter.Period{
    Max:  5,
    Size: 1000,
}

//初始化地址限流器
period.NewAddrLimiter()
```

* 获取访问许可

```
//用户10.110.2.1尝试获取访问许可
isOk := period.GetPermit("10.110.2.1")
if isOk {
    // 获取成功
    fmt.Println("Access successful")
} else {
    // 获取失败 
    fmt.Println("Access failed")
}
```