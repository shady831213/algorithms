package greedy

import (
	"testing"
	"reflect"
)

func TestBasicScheduler(t *testing.T) {
	s := newScheduler([]*task{newTask(1, 3, 0), newTask(0, 5, 1), newTask(5, 1, 2)}, 0)

	output := []chan *task{make(chan *task), make(chan *task)}

	expect := []*task{newTask(0,4,1), newTask(1,2,0), newTask(1,1,0),newTask(1,0,0),
	newTask(0,3,1), newTask(5,0,2),newTask(0,2,1),newTask(0,1,1), newTask(0,0,1)}
	for i := range expect {
		expect[i].clk = &s.clk
		if expect[i].p == 0 {
			expect[i].finishTime = i
		}
	}
	expect = append(expect, make([]*task,6,6)...)

	result := make([]*task, 0, 0)

	go func() {
		for {
			t := <-output[0]
			if t == nil {
				result = append(result, t)
			} else {
				temp := *t
				result = append(result, &temp)
			}
			output[1] <- nil
		}
	}()

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
		t.Logf("expect avgCompletedTime = %0f, but get %0f", 16.0/3.0 , s.avgCompletedTime())
		t.Fail()
	}

	if ! reflect.DeepEqual(result,expect) {
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
