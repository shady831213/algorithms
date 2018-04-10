package greedy

import (
	"reflect"
	"testing"
)

func collectResult(input []chan *task, result *[]*task) {
	for {
		t := <-input[0]
		if t == nil {
			*result = append(*result, t)
		} else {
			temp := *t
			*result = append(*result, &temp)
		}
		input[1] <- nil
	}
}

func TestBasicScheduler(t *testing.T) {
	s := newScheduler([]*task{newTask(1, 3, 0), newTask(0, 5, 1), newTask(5, 1, 2)}, 0)
	output := []chan *task{make(chan *task), make(chan *task)}
	result := make([]*task, 0, 0)
	expect := []*task{newTask(0, 4, 1), newTask(1, 2, 0), newTask(1, 1, 0), newTask(1, 0, 0),
		newTask(0, 3, 1), newTask(5, 0, 2), newTask(0, 2, 1), newTask(0, 1, 1), newTask(0, 0, 1)}
	for i := range expect {
		expect[i].clk = &s.clk
		if expect[i].p == 0 {
			expect[i].finishTime = i
		}
	}
	expect = append(expect, make([]*task, 6, 6)...)

	go collectResult(output, &result)
	s.run(nil, output)

	if s.finishedCnt != 3 {
		t.Logf("expect finishedCnt = %0d, but get %0d", 3, s.finishedCnt)
		t.Fail()
	}

	if s.totalCompletedTime != 16 {
		t.Logf("expect totalCompletedTime = %0d, but get %0d", 16, s.totalCompletedTime)
		t.Fail()
	}

	if s.avgCompletedTime() != 16.0/3.0 {
		t.Logf("expect avgCompletedTime = %0f, but get %0f", 16.0/3.0, s.avgCompletedTime())
		t.Fail()
	}

	if !reflect.DeepEqual(result, expect) {
		t.Log("Scheduled Seq wrong!")
		t.Log("expect is:")
		for i := range expect {
			t.Logf("%+v\n", expect[i])
		}
		t.Log("result is:")
		for i := range result {
			t.Logf("%+v\n", result[i])
		}
		t.Fail()
	}
}

func TestBasicScheduler_Insert(t *testing.T) {
	s := newScheduler([]*task{newTask(1, 3, 0), newTask(0, 5, 1), newTask(5, 1, 2)}, 30)
	output := []chan *task{make(chan *task), make(chan *task)}
	input := make(chan *task, 1)
	result := make([]*task, 0, 0)

	expect := []*task{newTask(0, 4, 1), newTask(0, 2, 3), newTask(0, 1, 3), newTask(0, 0, 3),
		newTask(1, 2, 0), newTask(3, 0, 6), newTask(5, 0, 2), newTask(4, 1, 4), newTask(4, 0, 4),
		newTask(1, 1, 0), newTask(1, 0, 0), newTask(0, 3, 1), newTask(0, 2, 1), newTask(0, 1, 1), newTask(0, 0, 1),
		newTask(2, 6, 5), newTask(2, 5, 5), newTask(2, 4, 5), newTask(2, 3, 5), newTask(2, 2, 5), newTask(2, 1, 5),
		newTask(2, 0, 5)}

	for i := range expect {
		expect[i].clk = &s.clk
		if expect[i].id == 3 || expect[i].id == 4 {
			expect[i].startTime = 1
		} else if expect[i].id == 5 || expect[i].id == 6 {
			expect[i].startTime = 2
		}
		if expect[i].p == 0 {
			expect[i].finishTime = i
		}

	}
	expect = append(expect, make([]*task, 8, 8)...)

	go collectResult(output, &result)
	go func() {
		input <- newTask(0, 3, 3)
		input <- newTask(4, 2, 4)
		input <- newTask(2, 7, 5)
		input <- newTask(3, 1, 6)
	}()
	s.run(input, output)

	if s.finishedCnt != 7 {
		t.Logf("expect finishedCnt = %0d, but get %0d", 7, s.finishedCnt)
		t.Fail()
	}
	//expect the running duration:finishTime - startTime
	if s.totalCompletedTime != 61 {
		t.Logf("expect totalCompletedTime = %0d, but get %0d", 61, s.totalCompletedTime)
		t.Fail()
	}

	if s.avgCompletedTime() != 61.0/7.0 {
		t.Logf("expect avgCompletedTime = %0f, but get %0f", 61.0/7.0, s.avgCompletedTime())
		t.Fail()
	}

	if !reflect.DeepEqual(result, expect) {
		t.Log("Scheduled Seq wrong!")
		t.Log("expect is:")
		for i := range expect {
			t.Logf("%+v\n", expect[i])
		}
		t.Log("result is:")
		for i := range result {
			t.Logf("%+v\n", result[i])
		}
		t.Fail()
	}
}
