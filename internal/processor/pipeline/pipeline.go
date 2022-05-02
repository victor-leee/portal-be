package pipeline

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"sync"
	"sync/atomic"
)

/**
Status change
INIT -> CLONING -> BUILDING -> PUSHING -> DEPLOYING -> SUCCESS
any failure would cause the status set to FAILURE
any other status is UNKNOWN
*/

type SimpleFuncWithErr func() error

type ProgressStatus int

const (
	Init ProgressStatus = iota
	Cloning
	Building
	Pushing
	Deploying
	Success
	Failure
	Unknown
)

var progress = sync.Map{}

type Manager interface {
	// Clone sets the clone function
	Clone(f SimpleFuncWithErr) Manager
	// Build sets the build function
	Build(f SimpleFuncWithErr) Manager
	// Push sets the push registry function
	Push(f SimpleFuncWithErr) Manager
	// Deploy sets the deploying to cluster function
	Deploy(f SimpleFuncWithErr) Manager
	// Do start the execution of the pipeline
	Do(ctx context.Context) (*string, error)
}

type defaultPipelineProgressManager struct {
	fs []SimpleFuncWithErr
}

func New() Manager {
	return &defaultPipelineProgressManager{
		fs: make([]SimpleFuncWithErr, 4),
	}
}

func (d *defaultPipelineProgressManager) Clone(f SimpleFuncWithErr) Manager {
	d.fs[0] = f
	return d
}

func (d *defaultPipelineProgressManager) Build(f SimpleFuncWithErr) Manager {
	d.fs[1] = f
	return d
}

func (d *defaultPipelineProgressManager) Push(f SimpleFuncWithErr) Manager {
	d.fs[2] = f
	return d
}

func (d *defaultPipelineProgressManager) Deploy(f SimpleFuncWithErr) Manager {
	d.fs[3] = f
	return d
}

func (d *defaultPipelineProgressManager) Do(ctx context.Context) (*string, error) {
	for _, f := range d.fs {
		if f == nil {
			return nil, errors.New("must provide all 4 functions to start pipeline")
		}
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	progress.Store(uid.String(), Init)
	go d.exec(ctx, uid)

	return proto.String(uid.String()), nil
}

func (d *defaultPipelineProgressManager) exec(ctx context.Context, uid uuid.UUID) {
	done := make(chan struct{})
	errStore := atomic.Value{}
	go func() {
		progress.Store(uid.String(), Cloning)
		for _, f := range d.fs {
			if e := f(); e != nil {
				errStore.Store(e)
				progress.Store(uid.String(), Failure)
				break
			}
			status, _ := progress.Load(uid.String())
			if intStatus, ok := status.(ProgressStatus); ok {
				progress.Store(uid.String(), intStatus+1)
			}
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		if errStore.Load() != nil {
			progress.Store(uid.String(), Failure)
			logrus.Errorf("ERROR:%v", errStore.Load())
		}
	case <-ctx.Done():
		progress.Store(uid.String(), Failure)
		errStore.Store(ctx.Done())
	}
}

func Query(uid string) ProgressStatus {
	if s, ok := progress.Load(uid); ok {
		return s.(ProgressStatus)
	}
	return Unknown
}
