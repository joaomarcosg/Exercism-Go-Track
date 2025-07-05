package lasagna

const OvenTime = 40

func RemaningOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElepsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}
