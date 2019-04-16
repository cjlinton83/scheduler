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
	finishedQ.ComputeAndShowStats(time, "First Come, First Serve")

	newQ.ShowList()
	finishedQ, time = sjf(newQ)
	finishedQ.ComputeAndShowStats(time, "Shortest Job First")

	newQ.ShowList()
	finishedQ, time = rr(newQ, 15)
	finishedQ.ComputeAndShowStats(time, "Round Robin")
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

func rr(newQ process.List, quantum int) (process.List, int) {
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

			if p.Start < 0 {
				p.Start = totalTime
			}

			if p.WorkingBurst <= quantum {
				totalTime += p.WorkingBurst
				p.Finished = totalTime
				finishedQ.PushBack(p)
			} else {
				totalTime += quantum
				p.WorkingBurst -= quantum
				runQ.PushBack(p)
			}
		}
	}

	return finishedQ, totalTime
}
