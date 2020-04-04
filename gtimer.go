package gtimer

import (
	"time"

	log "github.com/sirupsen/logrus"
)

//Timer ...
//最近一次执行的时间点
//计时任务数量
//log记录器
//
type Timer struct {
	t      time.Time
	logger *log.Logger
	count  int
}

const ()

//New new timer
func New() *Timer {
	return new(Timer)
}

//Count return timer count
func (t *Timer) Count() int {
	return t.count
}

//CycleTask ...
func (t *Timer) CycleTask(cyclet int, time time.Time, f func()) {

}

func (t *Timer) setLogger(logger *log.Logger) {
	t.logger = logger
	t.Println("set logger success")
}

//Println ...
func (t *Timer) Println(args ...interface{}) {
	if t.logger != nil {
		t.logger.Println(args...)
		return
	}
	log.Println(args...)
}
