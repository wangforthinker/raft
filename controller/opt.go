package controller

import "time"

type DataOpt struct {
	//if version == 0,means no version
	version int64
	//expire time
	expire time.Duration
}
