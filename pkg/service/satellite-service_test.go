package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	testService := NewSatelliteService()

	postX, postY := testService.GetLocation(100, 115.5, 142.7)

	assert.Equal(t, postX, 100)
	assert.Equal(t, postY, 200)
}
