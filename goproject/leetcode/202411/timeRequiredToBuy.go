package main

type Person struct {
	idx         int
	needTickets int
}

func timeRequiredToBuy(tickets []int, k int) int {
	times := 0
	newTickets := make([]*Person, 0)
	var kPerson *Person
	for idx, ticket := range tickets {
		p := &Person{idx, ticket}
		newTickets = append(newTickets, p)
		if idx == k {
			kPerson = p
		}
	}
	for kPerson.needTickets > 0 {
		tmp := newTickets[0]
		newTickets = newTickets[1:]
		times++
		if tmp == kPerson && tmp.needTickets == 1 {
			break
		} else if tmp.needTickets == 1 {
			continue
		} else {
			tmp.needTickets--
			newTickets = append(newTickets, tmp)
		}
	}
	return times
}
