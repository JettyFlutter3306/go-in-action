package concurrent

import "testing"

func TestSelectStmt(t *testing.T) {
	for i := 0; i < 10; i++ {
		println(i, "============")
		SelectStmt()
	}
}

func TestSelectFor(t *testing.T) {
	SelectFor()
}

func TestSelectBlock(t *testing.T) {
	SelectBlock()
}

func TestSelectNilChannel(t *testing.T) {
	SelectNilChannel()
}

func TestSelectNonBlock(t *testing.T) {
	SelectNonBlock()
}

func TestSelectRace(t *testing.T) {
	SelectRace()
}

func TestSelectAll(t *testing.T) {
	SelectAll()
}

func TestSelectChannelCloseSignal(t *testing.T) {
	SelectChannelCloseSignal()
}

func TestSelectSignal(t *testing.T) {
	SelectSignal()
}
