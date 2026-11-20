package main

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false 
	}

	front := q.data[0]
	q.data = q.data[1:]
	return front, true
}

func (q *Queue) Front() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}

	return q.data[0], true
}


func (q *Queue) Len() int {
	return len(q.data)
}