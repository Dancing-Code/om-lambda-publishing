package main

import (
	"time"
)

type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreateAt time.Time
}
