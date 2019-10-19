colorify [![CircleCI](https://circleci.com/gh/ahmedkamals/colorify.svg?style=svg)](https://circleci.com/gh/ahmedkamals/colorify "Build Status")
==========

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE  "License")
[![release](https://img.shields.io/github/release/ahmedkamals/colorify.svg?style=flat-square)](https://github.com/ahmedkamals/colorify/releases/latest "Release")
[![Travis CI](https://travis-ci.org/ahmedkamals/colorify.svg)](https://travis-ci.org/ahmedkamals/colorify "Cross Build Status [Linux, OSx]") 
[![codecov](https://codecov.io/gh/ahmedkamals/colorify/branch/master/graph/badge.svg)](https://codecov.io/gh/ahmedkamals/colorify "Code Coverage")
[![Coverage Status](https://coveralls.io/repos/github/ahmedkamals/colorify/badge.svg?branch=master)](https://coveralls.io/github/ahmedkamals/colorify?branch=master  "Code Coverage")
[![GolangCI](https://golangci.com/badges/github.com/ahmedkamals/colorify.svg?style=flat-square)](https://golangci.com/r/github.com/ahmedkamals/colorify "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/ahmedkamals/colorify)](https://goreportcard.com/report/github.com/ahmedkamals/colorify  "Go Report Card")
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c282df1ff33c43ddb5da1d7fe4e85674)](https://www.codacy.com/app/ahmedkamals/colorify?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ahmedkamals/colorify&amp;utm_campaign=Badge_Grade "Code Quality")
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/ahmedkamals/colorify "Documentation")
[![GoDoc](https://godoc.org/github.com/ahmedkamals/colorify?status.svg)](https://godoc.org/github.com/ahmedkamals/colorify "API Documentation")
[![Docker Pulls](https://img.shields.io/docker/pulls/ahmedkamal/colorify.svg?maxAge=604800)](https://hub.docker.com/r/ahmedkamal/colorify/ "Docker Pulls")
[![DepShield Badge](https://depshield.sonatype.org/badges/ahmedkamals/colorify/depshield.svg)](https://depshield.github.io "DepShield")
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fahmedkamals%2Fcolorify.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fahmedkamals%2Fcolorify?ref=badge_shield "Dependencies")
[![Join the chat at https://gitter.im/ahmedkamals/colorify](https://badges.gitter.im/ahmedkamals/colorify.svg)](https://gitter.im/ahmedkamals/colorify?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge "Let's discuss")

```bash

 _____       _            _  __       
/  __ \     | |          (_)/ _|      
| /  \/ ___ | | ___  _ __ _| |_ _   _ 
| |    / _ \| |/ _ \| '__| |  _| | | |
| \__/\ (_) | | (_) | |  | | | | |_| |
 \____/\___/|_|\___/|_|  |_|_|  \__, |
                                 __/ |
                                |___/ 

```

is a library that helps applying RGB colors based on [24 bit - ANSI escape sequences][1] for console output.

Table of Contents
-----------------

*   [🏎️ Getting Started](#getting-started)

    *   [Prerequisites](#prerequisites)
    *   [Example](#example)

*   [🕸️ Tests](#tests)

    *   [Benchmarks](#benchmarks)

*   [🤝 Contribution](#-contribution)

*   [👨‍💻 Author](#-author)

*   [🆓 License](#-license)

🏎️ Getting Started
------------------

### Prerequisites

*   [Golang 1.13 or later][2].

### Example

```go
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

	colorize.Set(style)
	print("Styled output!")
	colorize.Reset()
	println("\n\nSample Colors\n==============\n")
	println(colorize.Sample())
```

![Sample output](https://github.com/ahmedkamals/colorify/raw/master/assets/img/sample.png)

🕸️ Tests
--------
    
```bash
$ make test
```

### Benchmarks

![Benchmarks](https://github.com/ahmedkamals/colorify/raw/master/assets/img/bench.png)

🤝 Contribution
---------------

Please refer to [CONTRIBUTING.md](https://github.com/ahmedkamals/colorify/blob/master/CONTRIBUTING.md) file.

👨‍💻 Author
-----------

[ahmedkamals][3]

🆓 LICENSE
----------

Colorify is released under MIT license, refer [LICENSE](https://github.com/ahmedkamals/colorify/blob/master/LICENSE) file.  
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fahmedkamals%2Fcolorify.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fahmedkamals%2Fcolorify?ref=badge_large)


Enjoy!
[![Analytics](http://www.google-analytics.com/__utm.gif?utmwv=4&utmn=869876874&utmac=UA-136526477-1&utmcs=ISO-8859-1&utmhn=github.com&utmdt=colorify&utmcn=1&utmr=0&utmp=/ahmedkamals/colorify?utm_source=www.github.com&utm_campaign=colorify&utm_term=colorify&utm_content=colorify&utm_medium=repository&utmac=UA-136526477-1)]()

[1]: https://en.wikipedia.org/wiki/ANSI_escape_code#24-bit "Ansi Escape Sequenece"
[2]: https://golang.org/dl/ "Download Golang"
[3]: https://github.com/ahmedkamals "Author"
