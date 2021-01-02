package game

import "time"

// Timer is the
type Timer struct {
	TimeLeft time.Duration
}

// SetInitTime set the time allowed for a player in a game
func (t *Timer) SetInitTime(seconds int64) {

	t.TimeLeft = time.Duration(seconds) * time.Second

}
