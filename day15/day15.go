package day15

import (
	"strconv"
	"strings"
)

type Lens struct {
	label       string
	focalLength int
}

func hash(input string) int {
	currentValue := 0

	for _, char := range input {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}

func sumHash(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	sequence := strings.Split(input, ",")

	sum := 0

	for _, part := range sequence {
		sum += hash(part)
	}

	return sum
}

func placeLenses(input string) [][]Lens {
	input = strings.ReplaceAll(input, "\n", "")
	sequence := strings.Split(input, ",")

	boxes := make([][]Lens, 256)

	for _, step := range sequence {
		if strings.Contains(step, "-") {
			label := strings.Replace(step, "-", "", 1)
			lens := Lens{label: label}

			boxes = removeLens(boxes, lens)
		} else {
			parts := strings.Split(step, "=")
			focalLength, _ := strconv.Atoi(parts[1])

			lens := Lens{
				label:       parts[0],
				focalLength: focalLength,
			}

			boxes = addOrUpdateLens(boxes, lens)
		}
	}

	return boxes
}

func removeLens(boxes [][]Lens, oldLens Lens) [][]Lens {
	box := hash(oldLens.label)
	boxOfLenses := boxes[box]

	if len(boxOfLenses) > 0 {
		for i, lens := range boxOfLenses {
			if lens.label == oldLens.label {
				boxOfLenses = append(boxOfLenses[:i], boxOfLenses[i+1:]...)
				boxes[box] = boxOfLenses
			}
		}
	}

	return boxes
}

func addOrUpdateLens(boxes [][]Lens, newLens Lens) [][]Lens {
	box := hash(newLens.label)
	boxOfLenses := boxes[box]

	if len(boxOfLenses) == 0 {
		boxOfLenses = []Lens{}
	}

	found := false

	for i, lens := range boxOfLenses {
		if lens.label == newLens.label {
			boxOfLenses[i].focalLength = newLens.focalLength
			found = true
		}
	}

	if !found {
		boxOfLenses = append(boxOfLenses, newLens)
		boxes[box] = boxOfLenses
	}

	return boxes
}

func sumFocusingPower(boxes [][]Lens) int {
	sum := 0

	for i, box := range boxes {
		for j, lens := range box {
			slotNumber := j + 1
			focusingPower := (i + 1) * slotNumber * lens.focalLength
			sum += focusingPower
		}
	}

	return sum
}
