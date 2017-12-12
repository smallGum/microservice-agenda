package entities

import (
	"time"
)

type Meeting struct {
	Title         string    `xorm:"title"`
	Participators []string  `xorm:"participator"`
	StartTime     time.Time `xorm:"startTime"`
	EndTime       time.Time `xorm:"endTime"`
	Sponsor       string    `xorm:"sponsor"`
}
