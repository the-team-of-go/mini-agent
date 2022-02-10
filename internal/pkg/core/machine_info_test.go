package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCpuPercent(t *testing.T) {
	t.Run("test cpu percent get function", func(t *testing.T) {
		got := GetCpuPercent()
		assert.Greater(t, got, 0.0)
		assert.LessOrEqual(t, got, 100.0)
	})
}

func TestGetDiskPercent(t *testing.T) {
	t.Run("test disk percent get function", func(t *testing.T) {
		got := GetDiskPercent()
		assert.Greater(t, got, 0.0)
		assert.LessOrEqual(t, got, 100.0)
	})
}

func TestGetMemPercent(t *testing.T) {
	t.Run("test mem percent get function", func(t *testing.T) {
		got := GetMemPercent()
		assert.Greater(t, got, 0.0)
		assert.LessOrEqual(t, got, 100.0)
	})
}
