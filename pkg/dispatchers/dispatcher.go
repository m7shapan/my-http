package dispatchers

type Job interface {
	Execute() interface{}
}

type dispatcher struct {
	maxWorker     int
	resultChannel chan interface{}
	jobChannel    chan Job
	workers       []worker
}

func NewDispatcher(resultChannel chan interface{}, maxWorker int) dispatcher {
	jobChannel := make(chan Job, maxWorker)

	return dispatcher{
		maxWorker:     maxWorker,
		resultChannel: resultChannel,
		jobChannel:    jobChannel,
	}
}

func (d *dispatcher) Start() {
	for i := 0; i < d.maxWorker; i++ {
		worker := NewWorker(d.jobChannel, d.resultChannel)
		worker.Start()
		d.workers = append(d.workers, worker)
	}
}

func (d *dispatcher) Push(job Job) {
	d.jobChannel <- job
}

func (d *dispatcher) Close() {
	for i := 0; i < len(d.workers); i++ {
		d.workers[i].Close()
	}

	close(d.resultChannel)
}

type worker struct {
	jobChannel    chan Job
	resultChannel chan interface{}
	quit          chan chan bool
}

func NewWorker(jobChannel chan Job, resultChannel chan interface{}) worker {
	return worker{
		jobChannel:    jobChannel,
		resultChannel: resultChannel,
		quit:          make(chan chan bool),
	}
}

func (w *worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.jobChannel:
				w.resultChannel <- job.Execute()
			case q := <-w.quit:
				q <- true
				return
			}
		}
	}()
}

func (w *worker) Close() {
	quit := make(chan bool, 1)
	w.quit <- quit
	<-quit
}
