package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	start := 1
	end := n
	for start < end {
		m := start + (end-start)/2
		if isBadVersion(m) {
			end = m
		} else {
			start = m + 1
		}
	}
	return start
}
