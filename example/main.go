package main

import (
	"github.com/ahmedkamals/colorify"
	"os"
)

func main() {
	colorize := colorify.NewColorable(os.Stdout)
	style := colorify.Style{
		Foreground: &colorify.ColorValue{
			Red:   88,
			Green: 188,
			Blue:  88,
		},
		Background: &colorify.ColorValue{
			Red:   188,
			Green: 88,
			Blue:  8,
		},
		Font: []colorify.FontEffect{
			colorify.Bold,
			colorify.Italic,
			colorify.Underline,
			colorify.CrossedOut,
		},
	}

	println(colorize.Wrap("Hello styled", style))
	println(colorize.Black("Text in black!"))
	println(colorize.Blue("Deep clue C!"))
	println(colorize.Cyan("Hello cyan!"))
	println(colorize.Gray("Gray logged text!"))
	println(colorize.Green("50 shades of Green!"))
	println(colorize.Magenta("Magenta!"))
	println(colorize.Red("The thin Red light!"))
	println(colorize.Orange("Orange is the new black!"))

	colorize.Set(colorify.Style{
		Foreground: &colorify.ColorValue{
			Red:   255,
			Green: 188,
			Blue:  88,
		},
		Font: []colorify.FontEffect{colorify.Bold},
	})
	print("Output will be styled\ntill next reset!")
	colorize.Reset()
	println("\n\nSample Colors\n==============\n")
	println(colorize.Sample())
}
