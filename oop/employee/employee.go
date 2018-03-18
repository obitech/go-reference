package employee

import (
    "fmt"
)

// We're not exporting the struct itself
type employee struct {
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

// Kind of a constructor
func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
    return employee{firstName, lastName, totalLeaves, leavesTaken}
}

func (e employee) LeavesRemaining() {
    fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
