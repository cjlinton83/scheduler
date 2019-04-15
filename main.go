package main

import (
	"log"
	"sort"

	"github.com/cjlinton83/scheduler/process"
)

func main() {
	newQ, err := process.NewListFromFile("input.txt")
	if err != nil {
		log.Fatal("error:", err)
	}
	newQ.ShowList()

	finishedQ, time := fcfs(newQ)
	finishedQ.ShowStats(time)

	// finishedQ, time := sjf(newQ)
	// finishedQ.ShowStats(time)
}

func fcfs(newQ process.List) (process.List, int) {
	totalTime := 0

	arrivalQ := process.NewList()
	runQ := process.NewList()
	finishedQ := process.NewList()

	newQ.ClearStats()

	// initialize arrivalQ and sort by arrival time.
	arrivalQ = append(arrivalQ, newQ...)
	sort.SliceStable(arrivalQ, func(i, j int) bool {
		return arrivalQ[i].Arrival < arrivalQ[j].Arrival
	})

	for !arrivalQ.IsEmpty() || !runQ.IsEmpty() {
		if !arrivalQ.IsEmpty() {
			if arrivalQ.Front().Arrival <= totalTime {
				p := arrivalQ.PopFront()
				runQ.PushBack(p)
				continue
			}
		}

		if !runQ.IsEmpty() {
			p := runQ.PopFront()
			p.Start = totalTime
			totalTime += p.Burst
			p.Finished = totalTime
			finishedQ.PushBack(p)
		}
	}

	return finishedQ, totalTime
}

// func sjf(newQ process.List) (process.List, int) {

// }
