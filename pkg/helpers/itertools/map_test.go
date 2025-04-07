package itertools

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	expected := []int{2, 4, 6, 8}
	outputs := Map(inputs, func(i int) int {
		return i * 2
	})
	assert.Equal(t, expected, outputs)
}

func TestMapWithFailure(t *testing.T) {
	t.Run("successful conversion", func(t *testing.T) {
		inputs := []int{1, 2, 3, 4}
		expected := []string{"1", "2", "3", "4"}

		outputs, err := MapWithFailure(inputs, func(i int) (string, error) {
			return string(rune('0' + i)), nil
		})

		assert.NoError(t, err)
		assert.Equal(t, expected, outputs)
	})

	t.Run("failed conversion", func(t *testing.T) {
		inputs := []int{1, 2, 0, 4}
		expectedErr := errors.New("cannot convert zero")

		outputs, err := MapWithFailure(inputs, func(i int) (string, error) {
			if i == 0 {
				return "", expectedErr
			}
			return string(rune('0' + i)), nil
		})

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, outputs)
	})
}
