package task
type State int
const (
	PendingState State = iota // Task is queued but is waiting to be scheduled
	Scheduled
	Running
	Completed
	Failed
)

import (
	"github.com/google/uuid" // universally unique identifiers
	"github.com/docker/go-connections/nat"
	"time"
)

type Task struct {
	ID uuid.uuid
	Name string
	State State
	Image string
	Memory int
	Disk int
	ExposedPorts nat.PortSet
	PortBindings map[string]string
	RestartPolicy string
	StartTime time.time
	FinishTime time.Time
}
