package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	TasksCreated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tasks_created_total",
			Help: "Total number of created tasks",
		},
	)
	TasksDeleted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tasks_deleted_total",
			Help: "Total number of deleted tasks",
		},
	)
)

func Init() {
	prometheus.MustRegister(TasksCreated)
	prometheus.MustRegister(TasksDeleted)
}
