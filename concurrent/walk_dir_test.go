package concurrent

import (
	"fmt"
	"testing"
)

func TestWalkDirectory(t *testing.T) {
	dirs := []string{"."}
	fmt.Println(WalkDirectory(dirs...))
}
