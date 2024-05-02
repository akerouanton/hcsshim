//go:build windows

package hcsshim

import (
	"context"
	"io"
	"time"

	"github.com/Microsoft/hcsshim/internal/hcs/schema1"
)

// ProcessConfig is used as both the input of Container.CreateProcess
// and to convert the parameters to JSON for passing onto the HCS
type ProcessConfig = schema1.ProcessConfig

type Layer = schema1.Layer
type MappedDir = schema1.MappedDir
type MappedPipe = schema1.MappedPipe
type HvRuntime = schema1.HvRuntime
type MappedVirtualDisk = schema1.MappedVirtualDisk

// AssignedDevice represents a device that has been directly assigned to a container
//
// NOTE: Support added in RS5
type AssignedDevice = schema1.AssignedDevice

// ContainerConfig is used as both the input of CreateContainer
// and to convert the parameters to JSON for passing onto the HCS
type ContainerConfig = schema1.ContainerConfig

type ComputeSystemQuery = schema1.ComputeSystemQuery

// Container represents a created (but not necessarily running) container.
type Container interface {
	// Start synchronously starts the container.
	Start(context.Context) error

	// Shutdown requests a container shutdown, but it may not actually be shutdown until Wait() succeeds.
	Shutdown(context.Context) error

	// Terminate requests a container terminate, but it may not actually be terminated until Wait() succeeds.
	Terminate(context.Context) error

	// Waits synchronously waits for the container to shutdown or terminate.
	Wait() error

	// WaitTimeout synchronously waits for the container to terminate or the duration to elapse. It
	// returns false if timeout occurs.
	WaitTimeout(time.Duration) error

	// Pause pauses the execution of a container.
	Pause(context.Context) error

	// Resume resumes the execution of a container.
	Resume(context.Context) error

	// HasPendingUpdates returns true if the container has updates pending to install.
	HasPendingUpdates() (bool, error)

	// Statistics returns statistics for a container.
	Statistics(context.Context) (Statistics, error)

	// ProcessList returns details for the processes in a container.
	ProcessList(context.Context) ([]ProcessListItem, error)

	// MappedVirtualDisks returns virtual disks mapped to a utility VM, indexed by controller
	MappedVirtualDisks(context.Context) (map[int]MappedVirtualDiskController, error)

	// CreateProcess launches a new process within the container.
	CreateProcess(ctx context.Context, c *ProcessConfig) (Process, error)

	// OpenProcess gets an interface to an existing process within the container.
	OpenProcess(ctx context.Context, pid int) (Process, error)

	// Close cleans up any state associated with the container but does not terminate or wait for it.
	Close() error

	// Modify the System
	Modify(ctx context.Context, config *ResourceModificationRequestResponse) error
}

// Process represents a running or exited process.
type Process interface {
	// Pid returns the process ID of the process within the container.
	Pid() int

	// Kill signals the process to terminate but does not wait for it to finish terminating.
	Kill() error

	// Wait waits for the process to exit.
	Wait() error

	// WaitTimeout waits for the process to exit or the duration to elapse. It returns
	// false if timeout occurs.
	WaitTimeout(time.Duration) error

	// ExitCode returns the exit code of the process. The process must have
	// already terminated.
	ExitCode() (int, error)

	// ResizeConsole resizes the console of the process.
	ResizeConsole(width, height uint16) error

	// Stdio returns the stdin, stdout, and stderr pipes, respectively. Closing
	// these pipes does not close the underlying pipes; it should be possible to
	// call this multiple times to get multiple interfaces.
	Stdio() (io.WriteCloser, io.ReadCloser, io.ReadCloser, error)

	// CloseStdin closes the write side of the stdin pipe so that the process is
	// notified on the read side that there is no more data in stdin.
	CloseStdin() error

	// Close cleans up any state associated with the process but does not kill
	// or wait on it.
	Close() error
}
