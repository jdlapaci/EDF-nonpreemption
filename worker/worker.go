package worker

import (
	"EDF/constant"
	"log"
	"sync"
	"EDF/task"
	"time"
)

type TaskQueue struct {
	sync.Mutex
	Queue []*task.Task
}

type WorkerPool struct {
	sync.Mutex
	Pool []*Worker
}

// Worker is the agent to process tasks
type Worker struct {
	WorkerID 	int
	TaskChan 	chan *task.Task
	StopChan 	chan interface{}
	FlagChan	chan interface{}
}


func NewWorker(workerid int) *Worker{
	return &Worker{
		WorkerID:	workerid,
		TaskChan:	make(chan *task.Task,constant.WORKER_NR),
		StopChan:	make(chan interface{}),
		FlagChan:	make(chan interface{},1),
	}
}

func (wp *WorkerPool) CreateWorkerPool() {
	for i:=0;i<constant.WORKER_NR;i++ {
		wp.Pool=append(wp.Pool,NewWorker(i))
	}
}



// TaskProcessLoop processes tasks without preemption
func (w *Worker) TaskProcessLoop() {
	log.Printf("Worker<%d>: Task processor starts\n", w.WorkerID)
loop:
	for {
		select {
		case t := <-w.TaskChan:
			// This worker receives a new task to run
			// To be implemented
			w.Process(t)
			w.FlagChan <- 0 	// this will alert the scheduler that this worker is ready to be put back into its FreeWorkerBuf
			log.Printf("worker<%d> sends flag\n",w.WorkerID)
		case <-w.StopChan:
			// Receive signal to stop
			// To be implemented
			w.FlagChan <- 0
			break loop
		}
	}
	log.Printf("Worker<%d>: Task processor ends\n", w.WorkerID)
}

// Process runs a task on a worker without preemption
func (w *Worker) Process(t *task.Task) {
	log.Printf("Worker <%d>: App<%s>/Task<%d> starts (ddl %v)\n", w.WorkerID, t.AppID, t.TaskID, t.Deadline)
	// Process the task
	// To be implemented
	time.Sleep(t.TotalRunTime) // This simulates the time it takes for a task to run to completion
	log.Printf("Worker <%d>: App<%s>/Task<%d> ends\n", w.WorkerID, t.AppID, t.TaskID)
}
