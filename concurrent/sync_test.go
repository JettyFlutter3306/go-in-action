package concurrent

import "testing"

func TestSyncErr(t *testing.T) {
	SyncErr()
}

func TestSyncLock(t *testing.T) {
	SyncLock()
}

func TestSyncMutex(t *testing.T) {
	SyncMutex()
}

func TestSyncLockAndNo(t *testing.T) {
	SyncLockAndNo()
}

func TestSyncRLock(t *testing.T) {
	SyncRLock()
}

func TestSyncMapErr(t *testing.T) {
	SyncMapErr()
}

func TestSyncMap(t *testing.T) {
	SyncMap()
}

func TestSyncMapMethod(t *testing.T) {
	SyncMapMethod()
}

func TestSyncPool(t *testing.T) {
	SyncPool()
}

func TestSyncOnce(t *testing.T) {
	SyncOnce()
}

func TestSyncCond(t *testing.T) {
	SyncCond()
}
