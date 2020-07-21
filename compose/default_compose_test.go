package compose

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDefaultSequence(t *testing.T) {
	got := CreateDefaultSequence()
	assert.Equal(t, 16, len(got), "Processor count should be 14")
}
