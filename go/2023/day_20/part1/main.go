package main

import (
	"fmt"
	"os"
	"strings"
)

type Module struct {
	name       string
	typeModule string
	memory     any
	outputs    []string
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_20/input.txt")

	modules := map[string]Module{}
	broadcastTargets := []string{}

	for _, line := range strings.Split(string(f), "\n") {
		values := strings.Split(line, " -> ")
		outputs := strings.Split(values[1], ", ")

		if values[0] == "broadcaster" {
			broadcastTargets = append(broadcastTargets, outputs...)
		} else {
			typeModule := string(values[0][0])
			name := string(values[0][1:])
			if typeModule == "%" {
				modules[name] = Module{name: name, typeModule: typeModule, outputs: outputs, memory: "off"}
			} else {
				modules[name] = Module{name: name, typeModule: typeModule, outputs: outputs}
			}
		}
	}

	for name, module := range modules {
		for _, output := range module.outputs {
			outputModule := modules[output]

			if modules[output].typeModule != "&" {
				continue
			}

			if outputModule.memory == nil {
				outputModule.memory = map[string]string{}
			}

			outputModule.memory.(map[string]string)[name] = "low"
			modules[output] = outputModule
		}
	}

	low, high := 0, 0

	for i := 0; i < 1000; i++ {
		low++

		queue := [][3]string{}
		for _, target := range broadcastTargets {
			queue = append(queue, [3]string{"broadcaster", target, "low"})
		}

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]

			origin, target, pulse := item[0], item[1], item[2]

			if pulse == "low" {
				low++
			} else {
				high++
			}

			module, ok := modules[target]

			if !ok {
				continue
			}

			if module.typeModule == "%" {
				if pulse == "high" {
					continue
				}

				var outgoing string

				if module.memory == "on" {
					module.memory = "off"
					outgoing = "low"
				} else {
					module.memory = "on"
					outgoing = "high"
				}

				for _, x := range module.outputs {
					queue = append(queue, [3]string{module.name, x, outgoing})
				}
			} else {
				module.memory.(map[string]string)[origin] = pulse

				outgoing := "low"

				for _, memory := range module.memory.(map[string]string) {
					if memory == "low" {
						outgoing = "high"
					}
				}

				for _, x := range module.outputs {
					queue = append(queue, [3]string{module.name, x, outgoing})
				}
			}
			modules[target] = module
		}
	}

	fmt.Println(low * high)
}
