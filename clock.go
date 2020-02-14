package main

import (
	"fmt"
	"time"
)

func main() {

	zero := [...]string{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := [...]string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := [...]string{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := [...]string{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	
	four := [...]string{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	
	five := [...]string{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	
	six := [...]string{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	
	
	seven := [...]string{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := [...]string{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	
	nine := [...]string{
		"███",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	
	
	colon := [...]string{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}
	
	numbers := [...][5]string{zero, one, two, three, four, five, six, seven, eight, nine,}
		
		now := time.Now()

		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...][5]string{
		numbers[hour/10], numbers[hour%10],
		colon,
		numbers[min/10], numbers[min%10],
		colon,
		numbers[sec/10], numbers[sec%10],
		}

		for line := range clock[0] {
			for digot := range clock {
				fmt.Print(clock[digot][line], " ")
			}
			fmt.Println()
		}
}
