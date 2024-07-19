package domain

import "time"

type Task struct {
	Id     int
	UserId int
	Start  time.Time
	End    time.Time
}
