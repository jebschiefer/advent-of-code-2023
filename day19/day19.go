package day19

import (
	"aoc2023/utilities"
	"regexp"
	"strconv"
	"strings"
)

type Workflow struct {
	label string
	rules Rules
}

type Workflows map[string]Workflow

type Rule struct {
	category    string
	comparator  string
	valueLimit  int
	destination string
}

type Rules []Rule

type Rating struct {
	category string
	value    int
}

type RatedPart map[string]Rating
type RatedParts []RatedPart

func sumParts(ratedParts RatedParts) int {
	sum := 0

	for _, ratedPart := range ratedParts {
		for _, rating := range ratedPart {
			sum += rating.value
		}
	}

	return sum
}

func getAccepted(input string) RatedParts {
	accepted := RatedParts{}
	workflows, ratedParts := parseInput(input)

	for _, ratedPart := range ratedParts {
		destination := sortPart("in", workflows, ratedPart)

		if destination == "A" {
			accepted = append(accepted, ratedPart)
		}
	}

	return accepted
}

func parseInput(input string) (Workflows, RatedParts) {
	re := regexp.MustCompile(`\n\n`)
	parts := re.Split(input, -1)

	return parseWorkflows(parts[0]), parseRatedParts(parts[1])
}

func parseWorkflows(input string) Workflows {
	workflows := Workflows{}
	lines := utilities.GetLines(input)

	re := regexp.MustCompile(`(\w+){(.+)}`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		label := matches[1]
		ruleString := matches[2]

		workflows[label] = Workflow{
			label: label,
			rules: parseRules(ruleString),
		}
	}

	return workflows
}

func parseRules(ruleString string) Rules {
	ruleList := Rules{}
	rules := strings.Split(ruleString, ",")

	re := regexp.MustCompile(`^([x,m,a,s])(.)(\d+):(.+)`)

	for _, rule := range rules {
		newRule := Rule{
			category:    "default",
			destination: rule,
		}

		if strings.Contains(rule, ":") {
			matches := re.FindStringSubmatch(rule)
			valueLimit, _ := strconv.Atoi(matches[3])

			newRule.category = matches[1]
			newRule.comparator = matches[2]
			newRule.valueLimit = valueLimit
			newRule.destination = matches[4]
		}

		ruleList = append(ruleList, newRule)
	}

	return ruleList
}

func parseRatedParts(input string) RatedParts {
	ratedParts := RatedParts{}
	lines := utilities.GetLines(input)

	re := regexp.MustCompile(`(.)=(\d+)`)

	for _, line := range lines {
		ratedPart := RatedPart{}
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			category := match[1]
			value, _ := strconv.Atoi(match[2])

			ratedPart[category] = Rating{
				category: category,
				value:    value,
			}
		}

		ratedParts = append(ratedParts, ratedPart)
	}

	return ratedParts
}

func sortPart(start string, workflows Workflows, ratedPart RatedPart) string {
	workflow := workflows[start]

	for _, rule := range workflow.rules {
		if rule.category == "default" {
			if rule.destination == "R" || rule.destination == "A" {
				return rule.destination
			}

			return sortPart(rule.destination, workflows, ratedPart)
		}

		destination := ""
		partRating := ratedPart[rule.category]

		if rule.comparator == ">" && partRating.value > rule.valueLimit {
			destination = rule.destination
		} else if rule.comparator == "<" && partRating.value < rule.valueLimit {
			destination = rule.destination
		}

		if destination != "" {
			return sortPart(destination, workflows, ratedPart)
		}
	}

	return start
}

func getPossibleRatingsCombinations(label string, workflows Workflows, _valueRanges map[string]map[string]int, buckets map[string]int) map[string]int {
	oldValueRanges, valueRanges := map[string]map[string]int{}, map[string]map[string]int{}

	for category, valueRange := range _valueRanges {
		oldValueRanges[category] = map[string]int{
			"min": valueRange["min"],
			"max": valueRange["max"],
		}
		valueRanges[category] = map[string]int{
			"min": valueRange["min"],
			"max": valueRange["max"],
		}
	}

	subTotal := getCominationsForRange(valueRanges)

	if label == "A" || label == "R" {
		buckets[label] += subTotal
	}

	workflow := workflows[label]

	for _, rule := range workflow.rules {
		if rule.category == "default" {
			return getPossibleRatingsCombinations(rule.destination, workflows, oldValueRanges, buckets)
		}

		currentRange := valueRanges[rule.category]

		destination := ""

		if rule.comparator == ">" && currentRange["max"] > rule.valueLimit {
			currentRange["min"] = rule.valueLimit + 1
			oldValueRanges[rule.category]["max"] = currentRange["min"]
			destination = rule.destination
		} else if rule.comparator == "<" && currentRange["max"] > rule.valueLimit {
			currentRange["max"] = rule.valueLimit - 1
			oldValueRanges[rule.category]["min"] = currentRange["max"]
			destination = rule.destination
		}

		if destination != "" {
			valueRanges[rule.category] = currentRange
			getPossibleRatingsCombinations(rule.destination, workflows, valueRanges, buckets)
		}
	}

	return buckets
}

func part2(input string) int {
	workflows, _ := parseInput(input)

	valueRanges := map[string]map[string]int{
		"x": {
			"min": 0,
			"max": 4000,
		},
		"m": {
			"min": 0,
			"max": 4000,
		},
		"a": {
			"min": 0,
			"max": 4000,
		},
		"s": {
			"min": 0,
			"max": 4000,
		},
	}

	buckets := map[string]int{
		"A": 1,
		"R": 1,
	}

	buckets = getPossibleRatingsCombinations("in", workflows, valueRanges, buckets)

	return buckets["A"]
}

func getCominationsForRange(valueRanges map[string]map[string]int) int {
	total := 1

	for _, valueRange := range valueRanges {
		valuesInRange := valueRange["max"] - valueRange["min"]
		total *= valuesInRange
	}

	return total
}
