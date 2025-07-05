package lasagna

const ovenTime = 40

func RemaningOvenTime(actualMinutesInOven int) int {
	return ovenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElepsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}
