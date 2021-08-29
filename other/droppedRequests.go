package main

func droppedRequests(requestTime []int32) int32 {
	droppedRequests := 0

	for i := range requestTime {
		if i > 2 && requestTime[i] == requestTime[i-3] {
			droppedRequests += 1
			continue
		}

		if i > 19 && requestTime[i]-requestTime[i-20] < 10 {
			droppedRequests += 1
			continue
		}

		if i > 59 && requestTime[i]-requestTime[i-60] < 60 {
			droppedRequests += 1
		}
	}

	return int32(droppedRequests)

}
