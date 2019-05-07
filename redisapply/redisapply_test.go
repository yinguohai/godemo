package redisapply

import "testing"

func TestConnect(t *testing.T) {
	Connect()
}

func TestGet(t *testing.T) {
	Get()
}

func TestLpush(t *testing.T) {
	Lpush()
}

func TestRpop(t *testing.T) {
	Rpop()
}

func TestPipeline(t *testing.T) {
	Pipeline()
}