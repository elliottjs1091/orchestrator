package task

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"
)

// universally unique identifiers

type State int

const (
	Pending State = iota // Task is queued but is waiting to be scheduled - iota auto-assigns a sequenced value
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	ContainerId   string
	Name          string
	State         State
	Image         string
	CPU           float64
	Memory        int64
	Disk          int64
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

type Config struct { // Holds configuration for tasks
	Name           string
	AttachStdin    bool
	AttachStdout   bool
	AttachStderror bool
	ExposedPorts   nat.PortSet
	Cmd            []string
	Image          string
	Cpu            float64
	Memory         int64
	Disk           int64
	Env            []string
	RestartPolicy  string
}

type Docker struct { // Uses Task struct from above - ContainerId field will be used to interact with containers
	Client *client.Client // Used to interact with Docker API
	Config Config
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(
		ctx, d.Config.Image, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	io.Copy(os.Stdout, reader)
}

func (cli *Client) ContainerCreate(
	ctx context.Context,
	config *container.Config,
	hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig,
	platform *specs.Platform,
	containerName string) (container.ContainerCreateCreatedBody, error)
