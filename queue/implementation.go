package queue

type Queue struct{
	vals []string
}

func (q *Queue) Enqueue(name string){
	q.vals = append(q.vals, name);
}

func (q *Queue) Dequeue() string{
	val := q.vals[0]
	q.vals = q.vals[1:len(q.vals)];
	return val 
}