package day08

import (
	"regexp"
	"strings"
)

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

func CountSteps(lines []string) int {
	directions := ParseDirections(lines[0])

	nodes := map[string]map[string]string{}

	for _, line := range lines[2:] {
		key := GetKey(line)
		left, right := GetLeftRight(line)

		nodes[key] = map[string]string{
			"key": key,
			"L":   left,
			"R":   right,
		}
	}

	return FollowDirections(directions, nodes)
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

func FollowDirections(directions DirectionRing, nodes map[string]map[string]string) int {
	currentDirection := directions.head
	currentNode := nodes["AAA"]
	count := 0

	for {
		if currentNode["key"] == "ZZZ" {
			break
		}

		currentNode = nodes[currentNode[currentDirection.value]]
		count++
		currentDirection = currentDirection.next
	}

	return count
}
