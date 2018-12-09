package manual

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManual_GetStepDuration(t *testing.T) {
	m := Manual{}
	assert.Equal(t, 1, m.GetStepDuration("a"))
	assert.Equal(t, 2, m.GetStepDuration("b"))
	assert.Equal(t, 3, m.GetStepDuration("c"))
}
