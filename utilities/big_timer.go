package utilities

import "time"

type TimerPair struct {
	result  string
	timeout time.Duration
}

type BigTimer struct {
	timer_in []TimerPair
	//timedist   time.Duration
	time_queue chan time.Duration
}

func NewBigTimer(timer_in []TimerPair) BigTimer {
	bt := BigTimer{}
	bt.timer_in = make([]TimerPair, len(timer_in))
	for i := 0; i < len(timer_in); i++ {
		bt.timer_in[i] = timer_in[i]
	}
	//bt.timedist = time.Duration(0)
	for tp := range timer_in {

	}
}
