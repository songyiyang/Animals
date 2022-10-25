package worker

import "github.com/songyiyang/Animals/poly/Math"

type WorkerTyp string

const (
	WorkerTypStrWorker    = "strWorker"
	WorkerTypNumberWorker = "numberWorker"
)

type worker struct {
	typ  WorkerTyp
	Math Math.Math
}

func (s *worker) Typ() WorkerTyp {
	return s.typ
}

func NewWorker(typ WorkerTyp) *worker {
	switch typ {
	case WorkerTypStrWorker:
		return &worker{typ: typ, Math: &Math.StringMath{}}
	case WorkerTypNumberWorker:
		return &worker{typ: typ, Math: &Math.NumberMath{}}
	default:
		panic("do not support worker type")
	}
}
