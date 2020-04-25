package gtimer

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

//Timer ...
//最近一次执行的时间点
//计时任务数量
//log记录器
//
type Timer struct {
	sync.RWMutex
	timer    *time.Timer
	t        time.Time        //最近一次清理，
	xt       time.Time        //下一次清理
	logger   *log.Logger      //log记录器
	taskList map[string]*task //存放任务
}

const ()

//New new timer
func New() *Timer {
	l := make(map[string]*task)
	t := time.Now()
	return &Timer{
		t:        t,
		xt:       t,
		taskList: l,
	}
}

//Count return timer count
func (T *Timer) Count() int {
	T.RLock()
	defer T.RUnlock()
	return len(T.taskList)
}

//CycleTask ...
func (T *Timer) CycleTask(name string, duration time.Duration, f func()) error {

	return nil
}

// Task ...
func (T *Timer) Task(name string, t time.Time, f func()) {
	task := newTask(name, t.Sub(time.Now()), f)
	if task == nil {
		T.println("task is nil")
		return
	}
	task.setState(1)

	T.addTask(task)

}

func (T *Timer) addTask(task *task) {
	T.Lock()
	defer T.Unlock()
	T.taskList[task.name] = task
	if T.xt.Sub(time.Now()) <= 0 || time.Now().Add(task.t).Sub(T.xt) < 0 {
		T.xt = time.Now().Add(task.t)
		go T.expirationCheck(task)
	}
	//添加对于定时执行的支持

}
/*待修改*/
func (T *Timer) expirationCheck(t *task) {//执行方法
	// T.Lock()
	// defer T.Unlock()
	T.Lock()
	if T.timer != nil {
		T.timer.Stop()
	}
	T.timer = time.NewTimer(t.t)
	T.Unlock()
	for range T.timer.C {
		t.f()
	}
	smallestDuration := time.Second*0 //储存最小的时间点
	m:= new(task)//最近的时间点
	T.RLock()

	for _, v := range T.taskList {
		if smallestDuration == 0 || v.t < smallestDuration {
			smallestDuration=v.t
			m=v
		}
	}
	if smallestDuration!=0{
		go T.expirationCheck(m)
	}
}

//SetLogger ...
func (T *Timer) SetLogger(logger *log.Logger) {
	if logger == nil {
		T.logger = logger
		T.println("set logger success")
	}
}

//println ...
func (T *Timer) println(args ...interface{}) {
	if T.logger != nil {
		T.logger.Println(args...)
		return
	}
	log.Println(args...)
}
