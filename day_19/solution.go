package day19

import (
	"github.com/pivovarit/aoc/util"
	"strconv"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("aplentyPart1", func() int {
		return aplentyPart1(input)
	})
	util.Timed("aplentyPart2", func() int {
		return aplentyPart2(input)
	})
}

const (
	greaterThan ConditionType = iota
	lessThan
)

type (
	Rating rune
	Range  struct{ from, to int }
	Part   struct {
		ratings map[Rating]int
	}
	PartRange map[Rating]Range
	Rule      struct {
		name       string
		conditions []Condition
		fallback   string
	}
	Condition struct {
		conditionType ConditionType
		threshold     int
		rating        Rating
		result        string
	}
	ConditionType int
)

func (r Range) length() int {
	return r.to - r.from + 1
}

const firstRule = "in"
const (
	resultAccepted = "A"
	resultRejected = "R"
	resultEmpty    = ""
)
const (
	ratingX Rating = 'x'
	ratingM Rating = 'm'
	ratingA Rating = 'a'
	ratingS Rating = 's'
)

func aplentyPart1(input []string) int {
	t := Tokenizer{input: input}
	workflows := parseRules(t.next())
	rawParts := t.next()

	var sum int
	for _, part := range parseParts(rawParts) {
		if evaluate(part, workflows) == resultAccepted {
			for _, v := range part.ratings {
				sum += v
			}
		}
	}

	return sum
}

func aplentyPart2(input []string) int {
	t := Tokenizer{input: input}
	return evaluateRange(map[Rating]Range{
		ratingX: {1, 4000},
		ratingM: {1, 4000},
		ratingA: {1, 4000},
		ratingS: {1, 4000}},
		firstRule, parseRules(t.next()))
}

func process(condition Condition, parts map[Rating]Range, workflows map[string]Rule) (sum int) {
	switch condition.result {
	case resultAccepted:
		switch condition.conditionType {
		case greaterThan:
			sum += combinations(withAdjustedRatingRange(condition.rating, parts, Range{condition.threshold + 1, parts[condition.rating].to}))
			parts[condition.rating] = Range{parts[condition.rating].from, condition.threshold}
		case lessThan:
			sum += combinations(withAdjustedRatingRange(condition.rating, parts, Range{parts[condition.rating].from, condition.threshold - 1}))
			parts[condition.rating] = Range{condition.threshold, parts[condition.rating].to}
		}
	case resultRejected:
		switch condition.conditionType {
		case greaterThan:
			parts[condition.rating] = Range{parts[condition.rating].from, condition.threshold}
		case lessThan:
			parts[condition.rating] = Range{condition.threshold, parts[condition.rating].to}
		}
	default:
		switch condition.conditionType {
		case greaterThan:
			adjustedRange := Range{condition.threshold + 1, parts[condition.rating].to}
			parts[condition.rating] = Range{parts[condition.rating].from, condition.threshold}
			sum += evaluateRange(withAdjustedRatingRange(condition.rating, parts, adjustedRange), condition.result, workflows)
		case lessThan:
			adjustedRange := Range{parts[condition.rating].from, condition.threshold - 1}
			parts[condition.rating] = Range{condition.threshold, parts[condition.rating].to}
			sum += evaluateRange(withAdjustedRatingRange(condition.rating, parts, adjustedRange), condition.result, workflows)
		}
	}
	return sum
}

func evaluateRange(parts map[Rating]Range, workflow string, workflows map[string]Rule) (sum int) {
	rule := workflows[workflow]

	for _, condition := range rule.conditions {
		sum += process(condition, parts, workflows)
	}

	switch rule.fallback {
	case resultAccepted:
		sum += combinations(parts)
	case resultRejected:
	default:
		if parts[ratingX].length() > 0 || parts[ratingM].length() > 0 || parts[ratingA].length() > 0 || parts[ratingS].length() > 0 {
			sum += evaluateRange(parts, rule.fallback, workflows)
		}
	}
	return sum
}

func withAdjustedRatingRange(rating Rating, parts map[Rating]Range, r Range) map[Rating]Range {
	var adjustedParts = make(map[Rating]Range)
	for k, v := range parts {
		adjustedParts[k] = v
	}

	adjustedParts[rating] = r
	return adjustedParts
}

func combinations(entry map[Rating]Range) int {
	last := -1
	var sum = 0
	for _, r := range entry {
		if last == -1 {
			last = r.length()
			sum = r.length()
		} else {
			sum *= r.length()
		}
	}
	return sum
}

func evaluate(part Part, workflows map[string]Rule) string {
	var outcome = firstRule
	for outcome != resultRejected && outcome != resultAccepted {
		rule := workflows[outcome]
		outcome = resultEmpty
		for _, condition := range rule.conditions {
			var result = resultEmpty
			if condition.conditionType == greaterThan && part.ratings[condition.rating] > condition.threshold {
				result = condition.result
			} else if condition.conditionType == lessThan && part.ratings[condition.rating] < condition.threshold {
				result = condition.result
			}
			if result != resultEmpty {
				outcome = result
				break
			}
		}
		if outcome == resultEmpty {
			outcome = rule.fallback
		}
	}
	return outcome
}

func parseParts(rawParts []string) []Part {
	var parts []Part
	for _, part := range rawParts {
		var ratings = make(map[Rating]int)
		for _, ratingLine := range strings.Split(part[1:len(part)-1], ",") {
			rating := Rating(ratingLine[0])
			value := strings.Split(ratingLine, "=")[1]
			valueNumeric, _ := strconv.Atoi(value)
			ratings[rating] = valueNumeric
		}
		parts = append(parts, Part{ratings})
	}

	return parts
}

func parseRules(workflows []string) map[string]Rule {
	var mappedRules = make(map[string]Rule)

	for _, line := range workflows {
		rule := parseRule(line)
		mappedRules[rule.name] = rule
	}

	return mappedRules
}

func parseRule(line string) Rule {
	split := strings.Split(line, "{")
	name := split[0]
	rules := split[1][:len(split[1])-1]
	rulesByComma := strings.Split(rules, ",")
	fallback := rulesByComma[len(rulesByComma)-1]
	rulesByComma = rulesByComma[:len(rulesByComma)-1]
	return Rule{name, parseConditions(rulesByComma), fallback}
}

func parseConditions(rulesByComma []string) []Condition {
	var parsedConditions []Condition
	for _, rule := range rulesByComma {
		ruleSplit := strings.Split(rule, ":")
		ruleString := ruleSplit[0]
		result := ruleSplit[1]
		if strings.ContainsRune(ruleString, '>') {
			rating := Rating(int32(ruleString[0]))
			value := strings.Split(ruleString, ">")[1]
			valueNumeric, _ := strconv.Atoi(value)
			parsedConditions = append(parsedConditions, Condition{
				conditionType: greaterThan,
				threshold:     valueNumeric,
				rating:        rating,
				result:        result,
			})
		} else {
			rating := Rating(int32(ruleString[0]))
			value := strings.Split(ruleString, "<")[1]
			valueNumeric, _ := strconv.Atoi(value)
			parsedConditions = append(parsedConditions, Condition{
				conditionType: lessThan,
				threshold:     valueNumeric,
				rating:        rating,
				result:        result,
			})
		}
	}
	return parsedConditions
}

type Tokenizer struct {
	input   []string
	lastIdx int
}

func (t *Tokenizer) next() []string {
	var next []string
	for i := t.lastIdx; i < len(t.input)-1; i++ {
		if len(t.input[i]) == 0 {
			next = t.input[t.lastIdx:i]
			t.lastIdx = i + 1
			return next
		}
	}

	result := t.input[t.lastIdx:]
	t.lastIdx = len(t.input) - 1
	return result
}
