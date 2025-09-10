package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var CalculateWorkingCarsPerHour float64 = float64(productionRate) * successRate/100
    return (CalculateWorkingCarsPerHour)
    //panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var CalculateWorkingCarsPerMinute int = int((float64(productionRate) * successRate)/60/100)
    return (CalculateWorkingCarsPerMinute)
	//panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var CalculateCost uint = uint(carsCount)%10 * 10000 + (uint(carsCount)-uint(carsCount)%10)/10*95000
    return (CalculateCost)
	//panic("CalculateCost not implemented")
}
