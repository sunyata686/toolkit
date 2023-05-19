package toolkit

import (
	"github.com/spf13/cast"
	"time"
)

func ToString(any interface{}) string {
	if t, ok := any.(time.Time); ok {
		return t.Format("2006-01-02 15:04:05")
	} else {
		return cast.ToString(any)
	}
}
