package day20

import (
	"container/list"
	"fmt"
	"github.com/pivovarit/aoc/util"
	"math"
	"slices"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("pulsePropagationPart1", func() int {
		return pulsePropagationPart1(input)
	})
	util.Timed("pulsePropagationPart2", func() int {
		return pulsePropagationPart2(input)
	})
}

const (
	noSignal SignalType = iota
	low
	high
)

const (
	conjunction rune = '&'
	flipFlop    rune = '%'
	arrow            = "->"
)

const (
	conjunctionModule ModuleType = iota
	flipFlopModule
)

type Module interface {
	emit(signalType SignalType) SignalType
	receive(signalType SignalType, source string)
	targets() []string
	moduleType() ModuleType
	name() string
}

type (
	ModuleType     uint8
	SignalType     uint8
	FlipFlopModule struct {
		id          string
		destination []string
		on          bool
	}
	ConjunctionModule struct {
		id          string
		destination []string
		downstream  map[string]SignalType
	}
	QueueItem struct {
		sender, node string
		signal       SignalType
	}
)

func (m *FlipFlopModule) moduleType() ModuleType {
	return flipFlopModule
}

func (m *FlipFlopModule) receive(SignalType, string) {
}

func (m *FlipFlopModule) emit(signalType SignalType) SignalType {
	switch signalType {
	case low:
		if !m.on {
			m.on = !m.on
			return high
		} else {
			m.on = !m.on
			return low
		}
	case high:
		return noSignal
	default:
		panic("illegal signal type")
	}
}

func (m *FlipFlopModule) name() string {
	return m.id
}

func (m *FlipFlopModule) targets() []string {
	return m.destination
}

func (m *ConjunctionModule) name() string {
	return m.id
}

func (m *ConjunctionModule) moduleType() ModuleType {
	return conjunctionModule
}

func (m *ConjunctionModule) emit(SignalType) SignalType {
	for _, signal := range m.downstream {
		if signal == low {
			return high
		}
	}
	return low
}

func (m *FlipFlopModule) String() string {
	return fmt.Sprintf("name: %s, type: %s, targets: %s, on: %v", m.id, "flip-flop", m.destination, m.on)
}

func (m *ConjunctionModule) receive(signalType SignalType, source string) {
	m.downstream[source] = signalType
}

func (m *ConjunctionModule) targets() []string {
	return m.destination
}

func (m *ConjunctionModule) String() string {
	return fmt.Sprintf("name: %s, type: %s, targets: %s, downstream: %v", m.id, "conjunction", m.destination, m.downstream)
}

func pulsePropagationPart1(input []string) int {
	broadcaster, modules := parse(input)
	sumHigh := 0
	sumLow := 0
	for i := 0; i < 1000; i++ {
		lowPulses, highPulses := propagateCounting(broadcaster, modules)
		sumLow += lowPulses + 1
		sumHigh += highPulses

	}
	return sumHigh * sumLow
}

func propagateUntilHigh(broadcaster []string, modules map[string]Module, nodes []string) map[string]int {
	var cyclesToHigh = make(map[string]int)
	for _, node := range nodes {
		cyclesToHigh[node] = 0
	}

	for cycle := 1; cycle < math.MaxInt; cycle++ {
		processingQueue := list.New()
		for _, node := range broadcaster {
			processingQueue.PushBack(QueueItem{
				sender: "broadcaster",
				node:   node,
				signal: low,
			})
		}

		for processingQueue.Len() > 0 {
			item := processingQueue.Front()
			body := item.Value.(QueueItem)
			processingQueue.Remove(item)

			if slices.Contains(nodes, body.sender) && body.signal == high {
				if cyclesToHigh[body.sender] == 0 {
					cyclesToHigh[body.sender] = cycle
				}
				var found = true
				for _, v := range cyclesToHigh {
					if v == 0 {
						found = false
					}
				}

				if found {
					return cyclesToHigh
				}
			}

			module, exists := modules[body.node]
			if exists {
				module.receive(body.signal, body.sender)
				signal := module.emit(body.signal)
				if signal != noSignal {
					for _, target := range module.targets() {
						processingQueue.PushBack(QueueItem{body.node, target, signal})
					}
				}
			}
		}
	}
	return cyclesToHigh
}

func propagateCounting(broadcaster []string, modules map[string]Module) (int, int) {
	lowPulses := 0
	highPulses := 0
	processingQueue := list.New()
	for _, node := range broadcaster {
		processingQueue.PushBack(QueueItem{
			sender: "broadcaster",
			node:   node,
			signal: low,
		})
	}

	for processingQueue.Len() > 0 {
		item := processingQueue.Front()
		body := item.Value.(QueueItem)
		processingQueue.Remove(item)

		switch body.signal {
		case low:
			lowPulses++
		case high:
			highPulses++
		}

		module, exists := modules[body.node]
		if exists {
			module.receive(body.signal, body.sender)
			signal := module.emit(body.signal)
			if signal != noSignal {
				for _, target := range module.targets() {
					processingQueue.PushBack(QueueItem{body.node, target, signal})
				}
			}
		}
	}
	return lowPulses, highPulses
}

func pulsePropagationPart2(input []string) int {
	broadcaster, modules := parse(input)
	rxDependency := findDependencies("rx", modules)[0]
	dependencies := findDependencies(rxDependency, modules)
	var cycles []int
	for _, i := range propagateUntilHigh(broadcaster, modules, dependencies) {
		cycles = append(cycles, i)
	}
	return lcm(cycles)
}

func findDependencies(node string, modules map[string]Module) (result []string) {
	for _, module := range modules {
		if slices.Contains(module.targets(), node) {
			result = append(result, module.name())
		}
	}
	return
}

func parse(input []string) ([]string, map[string]Module) {
	var modules = make(map[string]Module)
	var conjunctionModules []string
	var broadcaster []string

	for _, line := range input {
		var targets []string
		for _, target := range strings.Split(line[strings.Index(line, arrow)+len(arrow):], ",") {
			targets = append(targets, strings.TrimSpace(target))
		}

		module := strings.TrimSpace(line[:strings.Index(line, arrow)])

		if rune(module[0]) == flipFlop {
			m := FlipFlopModule{module[1:], targets, false}
			modules[module[1:]] = &m
		} else if rune(module[0]) == conjunction {
			m := ConjunctionModule{module[1:], targets, map[string]SignalType{}}
			modules[module[1:]] = &m
			conjunctionModules = append(conjunctionModules, m.name())
		} else {
			broadcaster = targets
		}
	}

	for _, moduleName := range conjunctionModules {
		cm := modules[moduleName]
		var downstream = make(map[string]SignalType)
		for _, dependency := range modules {
			if slices.Contains(dependency.targets(), cm.name()) {
				downstream[dependency.name()] = low
			}
		}
		modules[moduleName] = &ConjunctionModule{
			id:          moduleName,
			destination: cm.targets(),
			downstream:  downstream,
		}
	}

	return broadcaster, modules
}

func lcm(numbers []int) int {
	result := 1
	for _, x := range numbers {
		gcd := result
		b := x
		for b != 0 {
			t := b
			b = gcd % b
			gcd = t
		}
		result = result / gcd * x
	}
	return result
}
