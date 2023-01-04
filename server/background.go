package main

import "time"

type BackgroundServer struct {
	Storer Storer
}

func NewBackgroundServer(userStorer Storer) *BackgroundServer {
	return &BackgroundServer{
		Storer: userStorer,
	}
}

func (s *BackgroundServer) StartbackgroundTasks() {
	reportTicker := time.NewTicker(3 * time.Second)

	go func() {
		for range reportTicker.C {
			s.GenerateReport()
		}
	}()
}
