package main

func average(salary []int) float64 {
	minSalary := salary[0]
	maxSalary := salary[1]
	if minSalary > maxSalary {
		minSalary, maxSalary = maxSalary, minSalary
	}
	sum := salary[0] + salary[1]
	for i := 2; i < len(salary); i++ {
		sum += salary[i]
		if salary[i] < minSalary {
			minSalary = salary[i]
		} else if salary[i] > maxSalary {
			maxSalary = salary[i]
		}
	}
	return float64((sum - minSalary - maxSalary)) / float64((len(salary) - 2))
}
