//Package config handles configurations
package cred

import (
	"time"
)

type Profiler interface {
	Load(path string, profile string) (err error)
	Save(path string, filepath string) (err error)
	SetValue(key string, value string) (err error)
	String(key string) (value string)
	Duration(key string) (value time.Duration)
	Int64(key string) (value int64)
	Int(key string) (value int)
}
