package main

import (
	_ "embed"
	"strings"
)

type Device struct {
	outputs []string
}

func parse(input string) map[string]Device {
	devices := make(map[string]Device)
	for line := range strings.SplitSeq(input, "\n") {
		var device Device
		name := line[0:3]
		for output := range strings.FieldsSeq(line[5:]) {
			device.outputs = append(device.outputs, output)
		}
		devices[name] = device
	}
	return devices
}

func part1(input string) int {
	devices := parse(input)

	queue := []string{"you"}

	paths := 0
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]

		if name == "out" {
			paths++
			continue
		}

		device := devices[name]
		queue = append(queue, device.outputs...)
	}
	return paths
}

func visit(devices map[string]Device, visited map[string]int, from, to string) int {
	if from == to {
		// We found the goal device.
		return 1
	}

	if paths, ok := visited[from]; ok {
		// We already visited this device, just return the number of paths from here.
		return paths
	}

	paths := 0
	visited[from] = 0

	device := devices[from]
	for _, output := range device.outputs {
		paths += visit(devices, visited, output, to)
	}

	visited[from] = paths
	return paths
}

func part2(input string) int {
	devices := parse(input)

	// For these paths: svr ... dac ... fft ... out
	svrDac := visit(devices, make(map[string]int), "svr", "dac")
	dacFft := visit(devices, make(map[string]int), "dac", "fft")
	fftOut := visit(devices, make(map[string]int), "fft", "out")

	// For these paths: svr ... fft ... dac ... out
	svrFft := visit(devices, make(map[string]int), "svr", "fft")
	fftDac := visit(devices, make(map[string]int), "fft", "dac")
	dacOut := visit(devices, make(map[string]int), "dac", "out")

	return svrDac*dacFft*fftOut + svrFft*fftDac*dacOut
}

//go:embed puzzle.txt
var puzzle string
