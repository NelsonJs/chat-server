package mq

import (
	"errors"
	"sync"
)

type Queue interface {
	 subscriber(topic string) (<- chan interface{},error)
	 unSubscriber(topic string,topicChan <- chan interface{}) error
	 publish(topic string,msg interface{}) error
	 setCapacity(capacity int64)
	 close() error
	 broadcast(msg interface{},topicChans []chan interface{}) error
}

type QueueImpl struct {
	topicMap map[string][]chan interface{}
	exit chan bool
	capacity int64
	sync.RWMutex
}

func NewQueueImpl() *QueueImpl {
	return &QueueImpl{
		topicMap: map[string][]chan interface{}{},
		exit: make(chan bool),
	}
}

func (queue *QueueImpl) subscriber(topic string) (<- chan interface{},error)  {
	select {
	case <-queue.exit:
		return nil,errors.New("queue is close")
	default:
	}
	if topic == "" {
		return nil,errors.New("topic cannot empty")
	}
	topicChan := make(chan interface{},queue.capacity)
	queue.Lock()
	queue.topicMap[topic] = append(queue.topicMap[topic], topicChan)
	queue.Unlock()
	return topicChan,nil
}

func (queue *QueueImpl) unSubscriber(topic string,topicChan <- chan interface{}) error  {
	select {
	case <-queue.exit:
		return errors.New("queue is close")
	default:
	}
	if topic == "" {
		return errors.New("topic cannot empty")
	}
	queue.RLock()
	chs,ok := queue.topicMap[topic]
	queue.RUnlock()
	if !ok {
		return nil
	}
	queue.Lock()
	topicChans := make([]chan interface{},0)
	for _,v := range chs {
		if v == topicChan {
			continue
		}
		topicChans = append(topicChans, v)
	}
	queue.topicMap[topic] = topicChans
	queue.Unlock()
	return nil
}

func (queue *QueueImpl) publish(topic string,msg interface{}) error  {
	select {
	case <- queue.exit:
		return errors.New("queue is close")
	default:
	}
	queue.RLock()
	topicChans,ok := queue.topicMap[topic]
	queue.RUnlock()
	if !ok {
		return errors.New("topic does not exists")
	}
	return queue.broadcast(msg,topicChans)
}

func (queue *QueueImpl) broadcast(msg interface{},topicChans []chan interface{}) error {
	select {
	case <- queue.exit:
		return errors.New("queue is close")
	default:
	}
	for _,v := range topicChans {
		v <- msg
	}
	return nil
}


