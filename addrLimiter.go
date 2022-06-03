package apilimiter

import (
	"container/list"
	"sync"
	"time"
)

// Period 滑动窗口配置
type Period struct {
	Max  int64 //单个用户在滑动窗口内的最大访问次数
	Size int64 //滑动窗口大小（单位：毫秒）

	addrMap map[string]*list.List //用户访问记录桶，用于存储各个用户的访问记录
	m       sync.Mutex
}

//访问记录
type record struct {
	Addr      string //用户的地址
	Timestamp int64  //本次访问时间（毫秒级时间戳）
}

// NewAddrLimiter 初始化滑动窗口地址限流器
func (period *Period) NewAddrLimiter() {

	period.addrMap = make(map[string]*list.List)
}

// GetPermit 用户申请访问
func (period *Period) GetPermit(addr string) bool {

	//加锁，确保线程安全
	period.m.Lock()
	defer period.m.Unlock()

	oldSet, isOk := period.addrMap[addr]

	//如果没有该addr的记录
	if !isOk {

		set := list.New()

		//新建访问记录
		rec := record{
			Addr:      addr,
			Timestamp: time.Now().UnixNano() / 1e6,
		}

		//将本次访问记录插入list尾部
		set.PushBack(rec)

		//更新addrMap
		period.addrMap[addr] = set
		return true
	}

	//如果addrMap中已有该addr的访问记录
	set := oldSet

	//获取当前时间的毫秒级时间戳
	now := time.Now().UnixNano() / 1e6

	//累加数
	var num int64 = 0

	//从list头部开始遍历
	for log := set.Front(); log != nil; log = log.Next() {
		timestamp := log.Value.(record).Timestamp
		//如果该记录的时间在滑动窗口的范围内
		if timestamp+period.Size >= now {
			//累加
			num++
		} else {
			//如果记录在滑动窗口的范围外，删除该记录
			set.Remove(log)
		}
	}

	//如果该addr超过窗口周期内的访问次数限制，拒绝访问
	if num >= period.Max {
		return false
	}

	//允许访问
	rec := record{
		Addr:      addr,
		Timestamp: time.Now().UnixNano() / 1e6,
	}

	//将本次访问记录插入list尾部
	set.PushBack(rec)

	//更新addrMap
	period.addrMap[addr] = set
	return true
}
