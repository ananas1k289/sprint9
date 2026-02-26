package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	size := 1000
	data := generateRandomElements(size)

	assert.NotNil(t, data)
	assert.Len(t, data, size)
}

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	data := generateRandomElements(0)
	assert.Nil(t, data)
}

func TestMaximum_EmptySlice(t *testing.T) {
	result := maximum([]int{})
	assert.Equal(t, 0, result)
}

func TestMaximum_SingleElement(t *testing.T) {
	result := maximum([]int{42})
	assert.Equal(t, 42, result)
}

func TestMaximum_MultipleElements(t *testing.T) {
	data := []int{1, 5, 3, 9, 2}
	result := maximum(data)

	assert.Equal(t, 9, result)
}

func TestMaxChunks_EmptySlice(t *testing.T) {
	result := maxChunks([]int{})
	assert.Equal(t, 0, result)
}

func TestMaxChunks_LessThanChunks(t *testing.T) {
	data := []int{3, 7, 2}
	result := maxChunks(data)

	assert.Equal(t, 7, result)
}

func TestMaxChunks_EqualsSingleThreadResult(t *testing.T) {
	data := []int{1, 100, 50, 999, 23, 888, 77, 42}

	single := maximum(data)
	parallel := maxChunks(data)

	assert.Equal(t, single, parallel)
}

func TestMaxChunks_LargeData(t *testing.T) {
	data := generateRandomElements(100_000)

	single := maximum(data)
	parallel := maxChunks(data)

	assert.Equal(t, single, parallel)
}
