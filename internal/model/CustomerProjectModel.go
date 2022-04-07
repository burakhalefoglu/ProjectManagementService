package model

import "time"

type CustomerProject struct {
	Id         int64
	ProjectId  int64
	IndustryId int8
	ProductId  int8
	CreatedAt  time.Time
	Status     bool
}
