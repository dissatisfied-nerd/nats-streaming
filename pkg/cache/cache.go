package cache

import (
	"log"
	"sync"

	"github.com/dissatisfied-nerd/ns-service/pkg/model"
)

// simpliest cache on map with mutex ue to asynchronousness of subscriber

type MemCache struct {
	data  map[string]model.Order
	mutex sync.Mutex
}

func NewMemCache() *MemCache {
	var mutex sync.Mutex

	return &MemCache{
		make(map[string]model.Order),
		mutex,
	}
}

func (m *MemCache) Add(order model.Order) {
	m.mutex.Lock()
	m.data[order.Order_uid] = order
	m.mutex.Unlock()
}

func (m *MemCache) Get(id string) model.Order {
	defer m.mutex.Unlock()
	m.mutex.Lock()

	result, status := m.data[id]

	if !status {
		log.Println("No such order")
		return model.Order{}
	}

	return result
}
