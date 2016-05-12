package task

import (
	"time"
)

type TimerTask struct {
	InitTime  int64
	DelayTime int64
	C         chan int
	execute   func()
}

func NewTimerTask(initTime int64, delayTime int64, execute func()) *TimerTask {
	c := make(chan int)
	timerTask := &TimerTask{initTime, delayTime, c, execute}
	go func() {

	}()
	return timerTask
}
func NewSchedule(init int64, delay int64, schedule func()) {
	go func() {
		temp := init
		for {
			select {
			case <-time.After(time.Duration(temp) * time.Second):
				if init != delay {
					temp = delay
				}
				go schedule()
			}
		}
	}()
}
