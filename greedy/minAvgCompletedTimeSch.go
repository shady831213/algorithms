package greedy

import (
	"container/heap"
)

const (
	wait   = 0
	run    = 1
	finish = 2
)

type task struct {
	r                     int
	p                     int
	clk                   *int
	startTime, finishTime int
	id                    int
}

func (t *task) init(r, p, id int) *task {
	t.r, t.p, t.id = r, p, id
	return t
}

func (t *task) state() int {
	if *(t.clk) < t.r+t.startTime {
		return wait
	} else if t.p > 0 {
		return run
	} else {
		return finish
	}
}

func (t *task) run() {
	if t.p > 0 {
		t.p--
	}
	if t.p == 0 {
		t.finishTime = *(t.clk)
	}
}

func newTask(r, p, id int) *task {
	return new(task).init(r, p, id)
}

type schTaskList struct {
	tasks []*task
	heap.Interface
}

func (h *schTaskList) init(self heap.Interface, tasks []*task) *schTaskList {
	h.tasks = tasks
	h.Interface = self
	//build heap O(n)
	for i := range h.tasks {
		heap.Fix(h, h.Len()-1-i)
	}
	return h
}

func (h *schTaskList) Swap(i, j int) {
	h.tasks[i], h.tasks[j] = h.tasks[j], h.tasks[i]
}

func (h *schTaskList) Len() int {
	return len(h.tasks)
}

func (h *schTaskList) Push(v interface{}) {
	h.tasks = append(h.tasks, v.(*task))
}

func (h *schTaskList) Pop() (v interface{}) {
	h.tasks, v = h.tasks[:h.Len()-1], h.tasks[h.Len()-1]
	return
}

type preSchTaskList struct {
	schTaskList
}

func (h *preSchTaskList) init(tasks []*task) *preSchTaskList {
	h.tasks = tasks
	h.schTaskList.init(h, tasks)
	return h
}

func (h *preSchTaskList) Less(i, j int) bool {
	return h.tasks[i].r < h.tasks[j].r
}

func newPreSchTaskList(tasks []*task) *preSchTaskList {
	return new(preSchTaskList).init(tasks)
}

type runningSchTaskList struct {
	schTaskList
}

func (h *runningSchTaskList) init(tasks []*task) *runningSchTaskList {
	h.tasks = tasks
	h.schTaskList.init(h, tasks)
	return h
}

func (h *runningSchTaskList) Less(i, j int) bool {
	return h.tasks[i].p < h.tasks[j].p
}

func newRunningSchTaskList(tasks []*task) *runningSchTaskList {
	return new(runningSchTaskList).init(tasks)
}

type scheduler struct {
	preTasks                                      *preSchTaskList
	runningTasks                                  *runningSchTaskList
	clk, timeout, totalCompletedTime, finishedCnt int
}

func (s *scheduler) init(tasks []*task, timeout int) *scheduler {
	maxTime := 0
	for i := range tasks {
		tasks[i].clk = &s.clk
		tasks[i].startTime = s.clk
		maxTime += tasks[i].r + tasks[i].p
	}
	if timeout > maxTime {
		s.timeout = timeout
	} else {
		s.timeout = maxTime
	}

	s.preTasks, s.runningTasks = newPreSchTaskList(tasks), newRunningSchTaskList([]*task{})
	return s
}

func (s *scheduler) newTask(t *task) {
	t.clk = &s.clk
	t.startTime = s.clk
	if t.r == 0 {
		heap.Push(s.runningTasks, t)
	} else {
		heap.Push(s.preTasks, t)
	}
}

func (s *scheduler) avgCompletedTime() float64 {
	return float64(s.totalCompletedTime) / float64(s.finishedCnt)
}

func (s *scheduler) run(input chan *task, output []chan *task) {
	for s.clk < s.timeout {
		//check input
		select {
		case t := <-input:
			s.newTask(t)
			inputLen := len(input)
			for i := 0; i < inputLen; i++ {
				t = <-input
				s.newTask(t)
			}
		default:
		}

		//check preTasks
		for s.preTasks.Len() > 0 && s.preTasks.tasks[0].state() != wait {
			temp := heap.Pop(s.preTasks).(*task)
			heap.Push(s.runningTasks, temp)
		}
		//check runningTasks
		if s.runningTasks.Len() > 0 {
			s.runningTasks.tasks[0].run()
			output[0] <- s.runningTasks.tasks[0]
			if s.runningTasks.tasks[0].state() == finish {
				s.finishedCnt++
				s.totalCompletedTime += s.runningTasks.tasks[0].finishTime - s.runningTasks.tasks[0].startTime
				heap.Pop(s.runningTasks)
			}
		} else {
			output[0] <- nil
		}
		<-output[1]
		s.clk++
	}
}

func newScheduler(tasks []*task, timeout int) *scheduler {
	return new(scheduler).init(tasks, timeout)
}
