package main

import (
	"fmt"
	"sync"
)

type airportManager struct {
	isBandFree    bool
	lock          *sync.Mutex
	airplaneQueue []airplane
}

func newAirportManager() *airportManager {
	return &airportManager{
		isBandFree: true,
		lock:       &sync.Mutex{},
	}
}

func (s *airportManager) canLand(t airplane) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	fmt.Println("airportManager: we got your msg " + t.getName() + " please wait")
	if s.isBandFree {
		s.isBandFree = false
		fmt.Println("airportManager: you can land now " + t.getName())
		return true
	} else {
		fmt.Println("airportManager: you can't land now " + t.getName() + " please wait...")

	}
	s.airplaneQueue = append(s.airplaneQueue, t)
	return false
}

func (s *airportManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isBandFree {
		s.isBandFree = true
	}
	if len(s.airplaneQueue) > 0 {
		firstAirplane := s.airplaneQueue[0]
		s.airplaneQueue = s.airplaneQueue[1:]
		fmt.Println("airportManager: you can land now " + firstAirplane.getName())
		firstAirplane.permitLanding()
	}
}
