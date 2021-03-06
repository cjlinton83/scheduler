package process

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
)

// Process represents a collection of process data.
type Process struct {
	ID           int
	Arrival      int
	Burst        int
	Priority     int
	WorkingBurst int
	Start        int
	Finished     int
}

// List is a collection of Process pointers.
type List []*Process

// NewList returns an empty list of processes.
func NewList() List {
	return make(List, 0)
}

// NewListFromFile parses an input file and returns a List
// upon success.
func NewListFromFile(path string) (List, error) {
	list := make(List, 0)

	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(bytes.NewReader(fileBytes))

	sc.Scan() // this call to Scan() removes the header from the input file.
	for sc.Scan() {
		var id, arrival, burst, priority int

		_, err := fmt.Sscanf(sc.Text(), "%d %d %d %d",
			&id, &arrival, &burst, &priority)
		if err != nil {
			return nil, err
		}

		list = append(list, &Process{id, arrival, burst, priority, 0, 0, 0})
	}

	return list, nil
}

// ShowList shows information of processes in List.
func (l List) ShowList() {
	fmt.Println("ID\tArrival\tBurst\tPriority")

	for i := 0; i < len(l); i++ {
		fmt.Printf("%d\t%d\t%d\t%d\n",
			l[i].ID, l[i].Arrival, l[i].Burst, l[i].Priority)
	}
	fmt.Printf("Number of jobs: %d\n\n", len(l))
}

//ComputeAndShowStats shows process statistics for processes in List.
func (l List) ComputeAndShowStats(time int, name string) {
	throughput, turnaround, response := 0.0, 0.0, 0.0
	numJobs := len(l)

	fmt.Printf("Terminated jobs (%s)\n", name)
	fmt.Println("ID\tArrival\tCompletion")

	for i := 0; i < numJobs; i++ {
		fmt.Printf("%d\t%d\t%d\n",
			l[i].ID, l[i].Arrival, l[i].Finished)
		turnaround += float64(l[i].Finished - l[i].Arrival)
		response += float64(l[i].Start - l[i].Arrival)
	}

	throughput = float64(numJobs) / float64(time)
	turnaround /= float64(numJobs)
	response /= float64(numJobs)

	fmt.Println("Run Stats:")
	fmt.Printf("Throughput: %.2f\n", throughput)
	fmt.Printf("Average turnaround time: %.2f\n", turnaround)
	fmt.Printf("Average response time: %.2f\n", response)
	fmt.Println()
}

// ClearStats resets WorkingBurst, Start and Finished.
func (l List) ClearStats() {
	for i := 0; i < len(l); i++ {
		l[i].WorkingBurst = l[i].Burst
		l[i].Start = -1
		l[i].Finished = -1
	}
}

// IsEmpty returns true if the length of the List equals zero.
func (l List) IsEmpty() bool {
	return len(l) == 0
}

// Front returns the process at the front of the list to the caller.
func (l List) Front() *Process {
	return l[0]
}

// PopFront removes the process from the front of the list and returns it
// to the caller.
func (l *List) PopFront() *Process {
	p := (*l)[0]
	*l = (*l)[1:]
	return p
}

// PushBack appends a process to the back of the list.
func (l *List) PushBack(p *Process) {
	*l = append(*l, p)
}
