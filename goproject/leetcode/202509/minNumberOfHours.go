package main

// https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition/?envType=problem-list-v2&envId=array

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	energyAdd := 0
	experienceAdd := 0
	n := len(energy)
	for i := 0; i < n; i++ {
		if initialEnergy <= energy[i] {
			energyAdd += energy[i] - initialEnergy + 1
			initialEnergy += energy[i] - initialEnergy + 1
		}
		initialEnergy -= energy[i]
		if initialExperience <= experience[i] {
			experienceAdd += experience[i] - initialExperience + 1
			initialExperience += experience[i] - initialExperience + 1
		}
		initialExperience += experience[i]
	}
	return energyAdd + experienceAdd
}
