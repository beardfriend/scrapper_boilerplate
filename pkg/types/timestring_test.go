package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	s := time.Now()
	a1 := TimeString(s)
	a2 := TimeString(s)
	assert.Equal(t, a1, a2)
}
