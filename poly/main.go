package main

import (
	"fmt"
	"github.com/songyiyang/Animals/poly/worker"
	"reflect"
)

var strW = worker.NewWorker(worker.WorkerTypStrWorker)
var numberW = worker.NewWorker(worker.WorkerTypNumberWorker)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add("1", "2"))
}

func add(a interface{}, b interface{}) interface{} {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.String && tb.Kind() == reflect.String {
		return strW.Math.Add(a, b)
	} else if ta.Kind() == reflect.Int && tb.Kind() == reflect.Int {
		return numberW.Math.Add(a, b)
	}

	return nil
}
