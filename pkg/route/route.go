package route

func ReconstructRoute(tickets []Ticket) []string {
	if len(tickets) == 0 {
		return []string{}
	}

	ticketMap := make(map[string]string)
	reverseMap := make(map[string]string)
	for _, ticket := range tickets {
		ticketMap[ticket.Source] = ticket.Destination
		reverseMap[ticket.Destination] = ticket.Source
	}

	starts := []string{}
	for src := range ticketMap {
		if _, found := reverseMap[src]; !found {
			starts = append(starts, src)
		}
	}

	if len(starts) == 0 {
		starts = append(starts, tickets[0].Source)
	}

	var itinerary []string
	var fstart string
	for _, start := range starts {
		part := []string{start}
		fstart = start
		for {
			next, found := ticketMap[start]
			if !found {
				break
			}

			part = append(part, next)

			if next == fstart {
				break
			}

			start = next
		}
		itinerary = append(itinerary, part...)
	}

	return itinerary
}
