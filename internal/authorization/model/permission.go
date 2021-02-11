package model

type Permission struct {
	ID     int64
	Code   int64
	Object string
	Action string
}
