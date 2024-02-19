package random

import "testing"

func TestSleep(t *testing.T) {
	Sleep(1, 3)
	Sleep(1, 9)
	Sleep(0.1, 0.9)
	Sleep(0.1, 3)
}
