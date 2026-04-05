package storage

import (
	"time"

	"gitee.com/wennmu/gopkg.git/dologger"
)

// SyncScheduler handles periodic sync to COS
type SyncScheduler struct {
	sync       *COSSync
	interval   time.Duration
	stopCh     chan struct{}
	isRunning  bool
}

// NewSyncScheduler creates a new sync scheduler
func NewSyncScheduler(sync *COSSync, intervalSeconds int) *SyncScheduler {
	return &SyncScheduler{
		sync:     sync,
		interval: time.Duration(intervalSeconds) * time.Second,
		stopCh:   make(chan struct{}),
	}
}

// Start begins the periodic sync loop
func (s *SyncScheduler) Start() {
	if s.isRunning {
		dologger.Warn("sync scheduler is already running")
		return
	}
	s.isRunning = true
	dologger.Info("sync scheduler started, interval:", s.interval.String())

	go s.run()
}

// Stop stops the periodic sync loop
func (s *SyncScheduler) Stop() {
	if !s.isRunning {
		return
	}
	close(s.stopCh)
	s.isRunning = false
	dologger.Info("sync scheduler stopped")
}

func (s *SyncScheduler) run() {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			dologger.Info("scheduled sync to COS started")
			result, err := s.sync.PushToCOS()
			if err != nil {
				dologger.Error("scheduled sync to COS failed:", err.Error())
			} else {
				dologger.Info("scheduled sync to COS completed:", result.Message)
			}
		case <-s.stopCh:
			return
		}
	}
}
