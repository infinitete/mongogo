package model

import "time"

// 定义一个按时间格式的数据结构
type DateRange[T any] struct {

	/**
	  比如：
	  1. 统计2022年11月1日的数据，那么key为"2022-11-01"，From为2022年11月1日00时00分00秒的时间戳，End为2022年11月1日23时59分59秒的时间戳
	  2. 统计2022年11月月的数据，那么key为"2022-11"，From为2022年11月1日00时00分00秒的时间戳，End为2022年11月31日23时59分59秒的时间戳
	 **/
	Key   string // 唯一的标志
	Title string // 作用描述
	From  int64  // 起始时间戳，第一次统计时的时间戳为准
	End   int64  // 结束世界戳，最后一次统计时的时间戳为准
	Value *T     // 统计值
}

type ServiceStatics struct {
	Total int64 // 总数

	// 按服务模式统计
	// key: 0为事故救援，1为非事故救援
	Mode map[uint8]int64

	// 按服务类型统计
	// key: 0为乘用车救援，1为商用车
	Type map[uint8]int64
}

func (s *ServiceStatics) Inc(serveMode, serveType uint8) {
	if v, ok := s.Mode[serveMode]; ok {
		s.Mode[serveMode] = v + 1
	} else {
		s.Mode[serveMode] = 1
	}

	if v, ok := s.Type[serveType]; ok {
		s.Type[serveMode] = v + 1
	} else {
		s.Type[serveMode] = 1
	}

	s.Total += 1
}

func (s *ServiceStatics) Dec(serveMode, serveType uint8) int64 {
	if v, ok := s.Mode[serveMode]; ok && v > 0 {
		s.Mode[serveMode] = v - 1
	} else {
		s.Mode[serveMode] = 0
	}

	if v, ok := s.Type[serveType]; ok && v > 0 {
		s.Type[serveMode] = v + 1
	} else {
		s.Type[serveMode] = 0
	}

	if s.Total > 0 {
		s.Total = s.Total - 1
	}
}
