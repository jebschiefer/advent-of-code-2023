package day08

import (
	"regexp"
	"strings"
)

type Node map[string]string

type Direction struct {
	value string
	next  *Direction
}

type DirectionRing struct {
	head *Direction
	tail *Direction
}

func (ring *DirectionRing) Append(value string) {
	direction := &Direction{
		value: value,
	}

	if ring.head == nil {
		ring.head = direction
	} else {
		last := ring.tail
		last.next = direction
	}

	ring.tail = direction
	// Point last value to head to complete a ring.
	ring.tail.next = ring.head
}

func CountSteps(lines []string, startingPattern *regexp.Regexp, endingPattern *regexp.Regexp) int {
	directions := ParseDirections(lines[0])

	nodes := map[string]Node{}

	for _, line := range lines[2:] {
		key := GetKey(line)
		left, right := GetLeftRight(line)

		nodes[key] = Node{
			"key": key,
			"L":   left,
			"R":   right,
		}
	}

	startingNodes := GetStartingNodes(nodes, startingPattern)

	minSteps := []int{}

	for _, node := range startingNodes {
		currentMinSteps := FollowDirections(directions, nodes, node, endingPattern)
		minSteps = append(minSteps, currentMinSteps)
	}

	return LeastCommonMultiple(minSteps...)
}

func ParseDirections(line string) DirectionRing {
	directions := DirectionRing{}
	values := strings.Split(line, "")

	for _, value := range values {
		directions.Append(value)
	}

	return directions
}

func GetKey(line string) string {
	re := regexp.MustCompile(`^\w{3}`)
	return re.FindString(line)
}

func GetLeftRight(line string) (string, string) {
	re := regexp.MustCompile(`(\w{3}), (\w{3})`)
	matches := re.FindAllStringSubmatch(line, -1)
	return matches[0][1], matches[0][2]
}

func FollowDirections(directions DirectionRing, nodes map[string]Node, startingNode Node, endingPattern *regexp.Regexp) int {
	currentDirection := directions.head
	currentNode := startingNode
	count := 0

	for {
		if endingPattern.FindString(currentNode["key"]) != "" {
			break
		}

		currentNode = nodes[currentNode[currentDirection.value]]
		count++
		currentDirection = currentDirection.next
	}

	return count
}

func GetStartingNodes(nodes map[string]Node, startPattern *regexp.Regexp) []Node {
	startingNodes := []Node{}

	for key, node := range nodes {
		if startPattern.FindString(key) != "" {
			startingNodes = append(startingNodes, node)
		}
	}

	return startingNodes
}

// Least common multiple using greatest common divisor
// https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
func LeastCommonMultiple(numbers ...int) int {
	numbersCount := len(numbers)

	if numbersCount == 1 {
		return numbers[0]
	}

	a := numbers[0]
	b := numbers[1]

	result := a * b / GreatestCommonDivisor(a, b)

	if numbersCount > 2 {
		numbers = numbers[2:]

		for _, number := range numbers {
			result = LeastCommonMultiple(result, number)
		}
	}

	return result
}

// Greatest common divisor using the Euclidean Algorithm
// https://en.wikipedia.org/wiki/Greatest_common_divisor#Euclidean_algorithm
func GreatestCommonDivisor(a int, b int) int {
	if b == 0 {
		return a
	}

	return GreatestCommonDivisor(b, a%b)
}
