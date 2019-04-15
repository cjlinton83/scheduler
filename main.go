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

	finishedQ, time = sjf(newQ)
	finishedQ.ShowStats(time)
}

func fcfs(newQ process.List) (process.List, int) {
	totalTime := 0

	readyQ := process.NewList()
	runQ := process.NewList()
	finishedQ := process.NewList()

	newQ.ClearStats()

	// initialize readyQ and sort by arrival time.
	readyQ = append(readyQ, newQ...)
	sort.SliceStable(readyQ, func(i, j int) bool {
		return readyQ[i].Arrival < readyQ[j].Arrival
	})

	for !readyQ.IsEmpty() || !runQ.IsEmpty() {
		if !readyQ.IsEmpty() {
			if readyQ.Front().Arrival <= totalTime {
				p := readyQ.PopFront()
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

func sjf(newQ process.List) (process.List, int) {
	totalTime := 0

	readyQ := process.NewList()
	runQ := process.NewList()
	finishedQ := process.NewList()

	newQ.ClearStats()

	// initialize readyQ and sort by arrival time.
	readyQ = append(readyQ, newQ...)
	sort.SliceStable(readyQ, func(i, j int) bool {
		return readyQ[i].Arrival < readyQ[j].Arrival
	})

	for !readyQ.IsEmpty() || !runQ.IsEmpty() {
		if !readyQ.IsEmpty() {
			if readyQ.Front().Arrival <= totalTime {
				p := readyQ.PopFront()
				runQ.PushBack(p)
				continue
			}
		}

		if !runQ.IsEmpty() {
			// sort runQ by burst time (SJF) *Not Efficient*
			sort.SliceStable(runQ, func(i, j int) bool {
				return runQ[i].Burst < runQ[j].Burst
			})

			p := runQ.PopFront()
			p.Start = totalTime
			totalTime += p.Burst
			p.Finished = totalTime
			finishedQ.PushBack(p)
		}
	}

	return finishedQ, totalTime
}
