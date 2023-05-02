package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

var employeeInstance *employee

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

//constructure
func Init(firstName string, lastName string, totalName int, leavesTaken int) *employee {
	if employeeInstance == nil {
		employeeInstance = &employee{
			FirstName:   firstName,
			LastName:    lastName,
			TotalLeaves: totalName,
			LeavesTaken: leavesTaken,
		}
	}
	return employeeInstance
}
