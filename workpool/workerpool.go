package main

import (
	"fmt"
	"time"
)

const (
	MaxWorkers = 10    //os.Getenv("MAX_WORKERS")
	MaxQueue   = 10000 // os.Getenv("MAX_QUEUE")
)

type Payload struct {
	// [redacted]
}

// Job represents the job to be run
type Job struct {
	Id      int
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	id         int
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job, workerId int) Worker {
	fmt.Printf("\n Initializing worker Id	:%d", workerId)
	return Worker{
		WorkerPool: workerPool,
		id:         workerId,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			fmt.Printf("\n Declared worker	:%d	is free for job! and waiting for job!!", w.id)
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Done job Id	:%d", job.Id)
				// we have received a work request.
				// if err := job.Payload.UploadToS3(); err != nil {
				// log.Errorf("Error uploading to S3: %s", err.Error())
				// }

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		fmt.Println("\n Waiting for Job !!")
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

func main() {
	fmt.Println("Initializing worker pool ...")
	JobQueue = make(chan Job, MaxQueue)
	distpatcher := NewDispatcher(MaxWorkers)
	distpatcher.Run()
	i := 1
	for {
		i += 1
		fmt.Printf("\nPutting job	:%d  in Q", i)
		JobQueue <- Job{Id: i}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("All job done !!")
}
