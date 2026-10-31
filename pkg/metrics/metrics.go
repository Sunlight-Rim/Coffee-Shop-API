package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	UserCreatedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "user_created_counter",
		Help: "How many users have been created",
	})
	CoffeeOrderCreatedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "coffee_order_created_counter",
		Help: "How many coffee orders have been created",
	})
)

func init() {
	prometheus.MustRegister(UserCreatedCounter)
	prometheus.MustRegister(CoffeeOrderCreatedCounter)
}
