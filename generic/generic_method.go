package generic

import "sync"

// 定义泛型类型
type List[T int | float64] struct {
	data     []T
	max, min T
	m        sync.Mutex
}

// 定义泛型方法
func (l *List[T]) Add(e T) *List[T] {
	// 加锁为了并发安全
	l.m.Lock()
	defer l.m.Unlock()

	l.data = append(l.data, e)

	// 统计 min max
	if len(l.data) == 1 {
		l.min = e
		l.max = e
	}

	if e > l.max {
		l.max = e
	}

	if e < l.min {
		l.min = e
	}

	return l
}

func (l *List[T]) Data() []T {
	return l.data
}

func (l *List[T]) Max() T {
	return l.max
}

func (l *List[T]) Min() T {
	return l.min
}
