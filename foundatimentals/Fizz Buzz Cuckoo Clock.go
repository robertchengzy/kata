package foundatimentals

import (
	"strings"
	"strconv"
)

/*Your story
You've always loved both Fizz Buzz katas and cuckoo clocks, and when you walked by a garage sale and saw an ornate
cuckoo clock with a missing pendulum, and a "Beyond-Ultimate Raspberry Pi Starter Kit" filled with all sorts of sensors
and motors and other components, it's like you were suddenly hit by a beam of light and knew that it was your mission
to combine the two to create a computerized Fizz Buzz cuckoo clock!

You took them home and set up shop on the kitchen table, getting more and more excited as you got everything working
together just perfectly. Soon the only task remaining was to write a function to select from the sounds you had
recorded depending on what time it was:

Your plan
When a minute is evenly divisible by three, the clock will say the word "Fizz".
When a minute is evenly divisible by five, the clock will say the word "Buzz".
When a minute is evenly divisible by both, the clock will say "Fizz Buzz", with two exceptions:
On the hour, instead of "Fizz Buzz", the clock door will open, and the cuckoo bird will come out and "Cuckoo"
between one and twelve times depending on the hour.
On the half hour, instead of "Fizz Buzz", the clock door will open, and the cuckoo will come out and "Cuckoo" just once.
With minutes that are not evenly divisible by either three or five, at first you had intended to have the clock just
say the numbers ala Fizz Buzz, but then you decided at least for version 1.0 to just have the clock make a quiet,
subtle "tick" sound for a little more clock nature and a little less noise.
Your input will be a string containing hour and minute values in 24-hour time, separated by a colon, and with leading
zeros. For example, 1:34 pm would be "13:34".

Your return value will be a string containing the combination of Fizz, Buzz, Cuckoo, and/or tick sounds that the clock
needs to make at that time, separated by spaces. Note that although the input is in 24-hour time, cuckoo clocks'
cuckoos are in 12-hour time.

Some examples
"13:34"       "tick"
"21:00"       "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
"11:15"       "Fizz Buzz"
"03:03"       "Fizz"
"14:30"       "Cuckoo"
"08:55"       "Buzz"
"00:00"       "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
"12:00"       "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
Have fun!*/

func FizzBuzzCuckooClock(time string) string {
	h, _ := strconv.Atoi(time[0:2])
	m, _ := strconv.Atoi(time[3:5])
	switch {
	case m==0:
		return strings.Repeat("Cuckoo ", (h+11)%12) + "Cuckoo"
	case m==30:
		return "Cuckoo"
	case m%15 == 0:
		return "Fizz Buzz"
	case m%5 == 0:
		return "Buzz"
	case m%3 == 0:
		return "Fizz"
	default:
		return "tick"
	}
}

func FizzBuzzCuckooClock2(time string) string {
	timeArr := strings.Split(time, ":")
	clock, _ := strconv.ParseInt(timeArr[0], 10, 64)
	value, _ := strconv.ParseInt(timeArr[1], 10, 64)
	if value == 30 {
		return "Cuckoo"
	}

	if value == 0 {
		if clock == 0 {
			clock = 12
		} else if clock > 12 {
			clock -= 12
		}

		var message string
		for i := int64(0); i < clock; i++ {
			if i == 0 {
				message += "Cuckoo"
				continue
			}
			message += " Cuckoo"
		}
		return message
	}

	if value%3 == 0 && value%5 == 0 {
		return "Fizz Buzz"
	}

	if value%3 == 0 {
		return "Fizz"
	}

	if value%5 == 0 {
		return "Buzz"
	}

	return "tick"
}
