package syncapply

import "testing"

func TestAtomic(t *testing.T) {
	AtomicAdd()
}

func TestAtomicCompareSwrap(t *testing.T) {
	AtomicCompareSwrap()
}

func TestAtomicStore(t *testing.T) {
	AtomicStore()
}

func TestAtomicExport(t *testing.T) {
	AtomicExport()
}

func TestAtomicSwap(t *testing.T) {
	AtomicSwap()
}