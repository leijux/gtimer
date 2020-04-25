package gtimer

import (
	"sync"
	"sync/atomic"
	"time"
)

type task struct { //包含任务状态 还有多久开始执行任务 任务的方法
	//任务状态 1 等待执行  0执行结束  任务等待被回收 -1
	sync.RWMutex
	name string
	f    func()

	state int32 //项目的状态

	createdOn time.Time //缓存项目的创建时间戳。

	t time.Duration //多少点执行
}

func newTask(name string, t time.Duration, f func()) *task {
	if t < 0 {
		return nil
	}

	return &task{
		name:      name,
		f:         f,
		t:         t,
		createdOn: time.Now(),
	}
}

func (t *task) setState(i int32) error {
	switch i {
	case -1, 0, 1:
		atomic.StoreInt32(&t.state, i)
		return nil
	default:
		return nil
	}
}

func (t *task) getState() int32 {
	s := atomic.LoadInt32(&t.state)
	return s
}
