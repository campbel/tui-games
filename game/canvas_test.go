package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanvas(t *testing.T) {

	t.Run("single character, no offset", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(0, 0, []string{
			"1",
		})
		expected := []string{
			"1    ",
			"     ",
			"     ",
			"     ",
			"     ",
		}
		assert.Equal(expected, canvas.Content)
	})

	t.Run("single character, with offset", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(1, 1, []string{
			"1",
		})
		expected := []string{
			"     ",
			" 1   ",
			"     ",
			"     ",
			"     ",
		}
		assert.Equal(expected, canvas.Content)
	})

	t.Run("multiline, no offset", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(0, 0, []string{
			"1",
			"2",
			"3",
		})
		expected := []string{
			"1    ",
			"2    ",
			"3    ",
			"     ",
			"     ",
		}
		assert.Equal(expected, canvas.Content)
	})

	t.Run("multiline, with offset", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(1, 1, []string{
			"1",
			"2",
			"3",
		})
		expected := []string{
			"     ",
			" 1   ",
			" 2   ",
			" 3   ",
			"     ",
		}
		assert.Equal(expected, canvas.Content)
	})

	t.Run("multiline, with offset, with overlap", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(1, 1, []string{
			"1",
			"22",
			"333",
		})
		canvas.Draw(2, 2, []string{
			"1",
			"22",
			"333",
		})
		expected := []string{
			"     ",
			" 1   ",
			" 21  ",
			" 322 ",
			"  333",
		}
		assert.Equal(expected, canvas.Content)
	})
}

func TestRenderWindow(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(0, 0, []string{
			"1",
		})
		expected := []string{
			"1    ",
			"     ",
			"     ",
			"     ",
			"     ",
		}
		assert.Equal(expected, canvas.Render(0, 0, 5, 5))

		expected = []string{
			"1  ",
			"   ",
			"   ",
		}
		assert.Equal(expected, canvas.Render(0, 0, 3, 3))
	})

	t.Run("offset", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(1, 1, []string{
			"1",
		})
		expected := []string{
			"     ",
			" 1   ",
			"     ",
			"     ",
			"     ",
		}
		assert.Equal(expected, canvas.Render(1, 1, 5, 5))
		assert.Equal(expected, canvas.Render(2, 2, 5, 5))
		assert.Equal(expected, canvas.Render(100, 100, 5, 5), "even if the offset is outside the canvas, it should still render")

		expected = []string{
			"1  ",
			"   ",
			"   ",
		}
		assert.Equal(expected, canvas.Render(1, 1, 3, 3))
	})

	t.Run("large window", func(t *testing.T) {
		assert := assert.New(t)
		canvas := NewBlankCanvas(5, 5)
		canvas.Draw(1, 1, []string{
			"1",
		})
		expected := []string{
			"          ",
			" 1        ",
			"          ",
			"          ",
			"          ",
			"          ",
			"          ",
			"          ",
			"          ",
			"          ",
		}
		result := canvas.Render(1, 1, 10, 10)
		assert.Equal(expected, result)
	})
}
