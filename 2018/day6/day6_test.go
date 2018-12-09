package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Idx2Coord(t *testing.T) {
	m := NewMap(5, 4)
	assert.Equal(t, Coord{X: 0, Y: 0}, m.idx2Coord(0))
	assert.Equal(t, Coord{X: 1, Y: 0}, m.idx2Coord(1))
	assert.Equal(t, Coord{X: 4, Y: 0}, m.idx2Coord(4))
	assert.Equal(t, Coord{X: 0, Y: 1}, m.idx2Coord(5))
	assert.Equal(t, Coord{X: 4, Y: 1}, m.idx2Coord(9))
	assert.Equal(t, Coord{X: 0, Y: 2}, m.idx2Coord(10))
	assert.Equal(t, Coord{X: 1, Y: 2}, m.idx2Coord(11))
	assert.Equal(t, Coord{X: 0, Y: 3}, m.idx2Coord(15))
}

func TestMap_IsContour(t *testing.T) {
	m := NewMap(5, 4)
	assert.True(t, m.isContour(0))
	assert.True(t, m.isContour(4))
	assert.True(t, m.isContour(5))
	assert.False(t, m.isContour(6))
	assert.True(t, m.isContour(9))
	assert.False(t, m.isContour(8))
	assert.True(t, m.isContour(10))
	assert.False(t, m.isContour(11))
	assert.True(t, m.isContour(15))
	assert.True(t, m.isContour(16))
	assert.True(t, m.isContour(17))
}

func TestCoord_Distance(t *testing.T) {
	assert.Equal(t, 2, Coord{1, 1}.Distance(Coord{3, 1}))
	assert.Equal(t, 4, Coord{1, 1}.Distance(Coord{3, 3}))
	assert.Equal(t, 4, Coord{3, 3}.Distance(Coord{1, 1}))
}
