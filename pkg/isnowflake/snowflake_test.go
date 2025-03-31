package isnowflake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	node, err := New(1)
	assert.NoError(t, err)
	t.Log(node.Generate().Int64())
}
