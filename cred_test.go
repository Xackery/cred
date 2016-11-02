package cred

import (
	"testing"
	"time"
)

type NoErrProfiler struct{}

func (p *NoErrProfiler) Load(path string, profile string) (err error) {
	return
}

func (p *NoErrProfiler) Save(path string, filepath string) (err error) {
	return
}

func (p *NoErrProfiler) SetValue(key string, value string) (err error) {
	return
}

func (p *NoErrProfiler) String(key string) (value string) {
	return
}

func (p *NoErrProfiler) Duration(key string) (value time.Duration) {
	return
}

func (p *NoErrProfiler) Int64(key string) (value int64) {
	return
}

func (p *NoErrProfiler) Int(key string) (value int) {
	return
}

func TestNoErrProfiler(t *testing.T) {
	p := NoErrProfiler{}
	Interface(&p)
	p.Load("test", "test")
	p.Save("test", "test")
	p.SetValue("test", "test2")
	p.String("test")
	p.Duration("test")
	p.Int("test")
	p.Int64("test")
}

func Interface(p Profiler) {
	return
}
