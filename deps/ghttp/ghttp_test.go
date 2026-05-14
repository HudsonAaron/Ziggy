package ghttp

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHttpConfOK(t *testing.T) {
	addr, err := GetHttpConf(map[string]any{
		"ip":   "127.0.0.1",
		"port": 8080,
	})
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}
	if addr != "127.0.0.1:8080" {
		t.Fatalf("unexpected addr: %s", addr)
	}
}

func TestGetHttpConfDefaultIPAndFloatPort(t *testing.T) {
	addr, err := GetHttpConf(map[string]any{
		"port": float64(8081),
	})
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}
	if addr != "127.0.0.1:8081" {
		t.Fatalf("unexpected addr: %s", addr)
	}
}

func TestGetHttpConfInvalidPort(t *testing.T) {
	_, err := GetHttpConf(map[string]any{
		"port": 70000,
	})
	if err == nil {
		t.Fatalf("expected err for invalid port")
	}
}

func TestGetHttpConfNonIntegerFloatPort(t *testing.T) {
	_, err := GetHttpConf(map[string]any{
		"port": 8081.5,
	})
	if err == nil {
		t.Fatalf("expected err for non-integer float port")
	}
}

func TestRunMuxRouter(t *testing.T) {
	handler := RunMuxRouter([]HRouter{
		{
			Path: "/status",
			Handle: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("OK"))
			},
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: %d", rec.Code)
	}
	if rec.Body.String() != "OK" {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestStartAndStop(t *testing.T) {
	_ = Stop()
	t.Cleanup(func() {
		_ = Stop()
	})

	port := getFreePort(t)
	err := Start(map[string]any{
		"ip":   "127.0.0.1",
		"port": port,
	}, nil)
	if err != nil {
		t.Fatalf("expected start success, got %v", err)
	}

	if err = Stop(); err != nil {
		t.Fatalf("expected stop success, got %v", err)
	}
}

func TestStartWhenAlreadyRunning(t *testing.T) {
	_ = Stop()
	t.Cleanup(func() {
		_ = Stop()
	})

	port := getFreePort(t)
	err := Start(map[string]any{
		"ip":   "127.0.0.1",
		"port": port,
	}, nil)
	if err != nil {
		t.Fatalf("expected first start success, got %v", err)
	}

	err = Start(map[string]any{
		"ip":   "127.0.0.1",
		"port": port + 1,
	}, nil)
	if err == nil {
		t.Fatalf("expected second start failure when server already running")
	}
}

func getFreePort(t *testing.T) int {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to get free port: %v", err)
	}
	defer listener.Close()

	addr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		t.Fatalf("unexpected listener addr type: %T", listener.Addr())
	}
	if addr.Port <= 0 {
		t.Fatalf("invalid free port: %v", fmt.Sprint(addr.Port))
	}
	return addr.Port
}
