package metrics

import (
	"errors"
	"net"
	"sync"
	"testing"
	"time"
)

type testSession struct {
	endpoint string
	conn     *net.UDPConn
	timer    *time.Timer
	wg       sync.WaitGroup
	err      error
}

func newSession() *testSession {
	sess := &testSession{
		endpoint: ":12345",
		timer:    time.NewTimer(time.Second),
	}
	sess.timer.Stop()
	return sess
}

func (s *testSession) start() error {
	addr, err := net.ResolveUDPAddr("udp", s.endpoint)
	if err != nil {
		return err
	}
	s.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	s.wg.Add(1)
	go s.run()
	return nil
}

func (s *testSession) stop() error {
	// allow up to 2 sec to complete session
	s.timer.Reset(2 * time.Second)
	s.wg.Wait()
	s.conn.Close()
	return s.err
}

func (s *testSession) run() {
	defer s.wg.Done()
	buffer := make([]byte, 1024)
	for {
		select {
		case <-s.timer.C:
			s.err = errors.New("No metrics message. Exiting by timeout")
			return
		default:

			n, err := s.conn.Read(buffer)
			if err != nil {
				s.err = err
				return
			}
			if n > 0 {
				s.timer.Stop()
				return
			}
		}
	}
}

func TestMetric(t *testing.T) {
	sess := newSession()
	if err := sess.start(); err != nil {
		t.Fatal(err)
	}

	dims := map[string]string{
		"test":     "myTest",
		"location": "myLocation",
		"error":    "myError",
		"errClass": "myErrorClass",
	}

	if err := AddMetric(sess.endpoint, "metricsNS", "metricName", 1, dims); err != nil {
		t.Fatal(err)
	}

	if err := sess.stop(); err != nil {
		t.Fatal(err)
	}
}
