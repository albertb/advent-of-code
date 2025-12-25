package main

import (
	"fmt"
	"strings"
)

type Reindeer struct{ speed, duration, rest int }

func parse(input string) map[string]Reindeer {
	reindeers := make(map[string]Reindeer)
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) < 1 {
			continue
		}
		fields := strings.Fields(line)

		var r Reindeer
		fmt.Sscanf(fields[3], "%d", &r.speed)
		fmt.Sscanf(fields[6], "%d", &r.duration)
		fmt.Sscanf(fields[13], "%d", &r.rest)
		reindeers[fields[0]] = r
	}
	return reindeers
}

func part1(input string, time int) int {
	reindeers := parse(input)

	var furtherst int
	for _, r := range reindeers {
		var distance int
		left := time
		for left > 0 {
			distance += min(left, r.duration) * r.speed
			left -= r.duration + r.rest
		}
		furtherst = max(furtherst, distance)
	}
	return furtherst
}

type Racer struct {
	Reindeer

	position int  // How far this reindeer has flown.
	flying   bool // Whether this reindeer is resting (vs. flying).
	left     int  // How many seconds are left before mode change.
}

func Start(r Reindeer) Racer {
	return Racer{r, 0, true, r.duration}
}

func (r *Racer) race() int {
	if r.flying {
		if r.left > 0 {
			r.left--
			r.position += r.speed
		} else {
			r.left = r.rest - 1
			r.flying = false
		}
	} else {
		if r.left > 0 {
			r.left--
		} else {
			r.left = r.duration - 1
			r.position += r.speed
			r.flying = true
		}
	}
	return r.position
}

func part2(input string, time int) int {
	reindeers := parse(input)

	racers := make(map[string]*Racer)
	for name, r := range reindeers {
		racer := Start(r)
		racers[name] = &racer
	}

	points := make(map[string]int)
	for range time {
		var leaders []string // The name of the leaders.
		var ahead int        // How far they have flown.
		for name := range racers {
			pos := racers[name].race()
			if pos > ahead {
				ahead = pos
				leaders = []string{name}
			} else if pos == ahead {
				leaders = append(leaders, name)
			}
		}
		for _, leader := range leaders {
			points[leader]++
		}
	}

	var winner int
	for _, p := range points {
		winner = max(winner, p)
	}
	return winner
}
