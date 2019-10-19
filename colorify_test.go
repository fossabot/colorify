package colorify

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestWrap(t *testing.T) {
	testCases := []struct {
		id           string
		input        string
		appliedStyle Style
		expected     string
	}{
		{
			id:    "Should output in black color.",
			input: "Output this in black",
			appliedStyle: Style{
				Foreground: &ColorValue{
					Red:   0,
					Green: 0,
					Blue:  0,
				},
			},
			expected: "\x1b[38;2;0;0;0mOutput this in black\x1b[0m",
		},
		{
			id:    "Should output in bold Red color.",
			input: "Output this in bold Red",
			appliedStyle: Style{
				Foreground: &ColorValue{
					Red:   255,
					Green: 0,
					Blue:  0,
				},
				Font: []FontEffect{Bold},
			},
			expected: "\x1b[38;2;255;0;0;1mOutput this in bold Red\x1b[0m",
		},
		{
			id:    "Should output in bold italic Green color.",
			input: "Output this in bold italic Green",
			appliedStyle: Style{
				Foreground: &ColorValue{
					Red:   0,
					Green: 255,
					Blue:  0,
				},
				Font: []FontEffect{Bold, Italic},
			},
			expected: "\x1b[38;2;0;255;0;1;3mOutput this in bold italic Green\x1b[0m",
		},
	}

	colorize := NewColorable(os.Stdout)
	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			assert.Equal(t, fmt.Sprintf("%q", testCase.expected), fmt.Sprintf("%q", colorize.Wrap(testCase.input, testCase.appliedStyle)))
		})
	}
}

func TestSet(t *testing.T) {
	testCases := []struct {
		id           string
		input        string
		appliedStyle Style
		expected     string
	}{
		{
			id:    "Should output in black color.",
			input: "Output this in black",
			appliedStyle: Style{
				Foreground: &ColorValue{
					Red:   0,
					Green: 0,
					Blue:  0,
				},
			},
			expected: "\x1b[38;2;0;0;0mOutput this in black\x1b[0m",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			output := captureOutput(func(output *os.File) {
				colorize := NewColorable(output)
				colorize.Set(testCase.appliedStyle)
				fmt.Print(testCase.input)
				colorize.Reset()
			})

			assert.Equal(t, fmt.Sprintf("%q", testCase.expected), fmt.Sprintf("%q", output))
		})
	}
}

func TestReset(t *testing.T) {
	testCases := []struct {
		id           string
		input        string
		appliedStyle Style
		expected     string
	}{
		{
			id:    "Should output in normal color.",
			input: "Output this normally.",
			appliedStyle: Style{
				Foreground: &ColorValue{
					Red:   0,
					Green: 0,
					Blue:  0,
				},
			},
			expected: "\x1b[38;2;0;0;0m\x1b[0mOutput this normally.",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			output := captureOutput(func(output *os.File) {
				colorize := NewColorable(output)
				colorize.Set(testCase.appliedStyle)
				colorize.Reset()
				fmt.Print(testCase.input)
			})

			assert.Equal(t, fmt.Sprintf("%q", testCase.expected), fmt.Sprintf("%q", output))
		})
	}
}

func TestDirectColors(t *testing.T) {
	testCases := []struct {
		id           string
		input        string
		appliedStyle func(*Colorable, string) string
		expected     string
	}{
		{
			id:    "Should output in black color.",
			input: "Output this in black color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Black(str)
			},
			expected: "\x1b[38;2;0;0;0mOutput this in black color.\x1b[0m",
		},
		{
			id:    "Should output in blue color.",
			input: "Output this in blue color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Blue(str)
			},
			expected: "\x1b[38;2;0;0;255mOutput this in blue color.\x1b[0m",
		},
		{
			id:    "Should output in cyan color.",
			input: "Output this in cyan color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Cyan(str)
			},
			expected: "\x1b[38;2;0;255;255mOutput this in cyan color.\x1b[0m",
		},
		{
			id:    "Should output in gray color.",
			input: "Output this in gray color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Gray(str)
			},
			expected: "\x1b[38;2;128;128;128mOutput this in gray color.\x1b[0m",
		},
		{
			id:    "Should output in green color.",
			input: "Output this in green color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Green(str)
			},
			expected: "\x1b[38;2;0;255;0mOutput this in green color.\x1b[0m",
		},
		{
			id:    "Should output in magenta color.",
			input: "Output this in magenta color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Magenta(str)
			},
			expected: "\x1b[38;2;255;0;255mOutput this in magenta color.\x1b[0m",
		},
		{
			id:    "Should output in orange color.",
			input: "Output this in orange color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Orange(str)
			},
			expected: "\x1b[38;2;255;165;0mOutput this in orange color.\x1b[0m",
		},
		{
			id:    "Should output in red color.",
			input: "Output this in red color.",
			appliedStyle: func(colorable *Colorable, str string) string {
				return colorable.Red(str)
			},
			expected: "\x1b[38;2;255;0;0mOutput this in red color.\x1b[0m",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			colorize := NewColorable(os.Stdout)
			output := testCase.appliedStyle(colorize, testCase.input)

			assert.Equal(t, fmt.Sprintf("%q", testCase.expected), fmt.Sprintf("%q", output))
		})
	}
}

func captureOutput(f func(output *os.File)) string {
	rescueStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	f(writer)

	writer.Close()
	out, _ := ioutil.ReadAll(reader)
	os.Stdout = rescueStdout

	return string(out)
}
