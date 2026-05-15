package gactor

import (
	"testing"
	"time"
)

func TestStartInitState(t *testing.T) {
	actor, err := Start("test-start-init-state", func(actor *GActor[any], v ...interface{}) (any, error) {
		return "initialized", nil
	})
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	defer actor.StopTimeout(nil, 200*time.Millisecond)

	state := GetState(actor)
	if state != "initialized" {
		t.Fatalf("unexpected state: %v", state)
	}
}

func TestCallTimeout(t *testing.T) {
	actor, err := Start[any]("test-call-timeout", nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	defer func() {
		actor.StopTimeout(nil, 200*time.Millisecond)
	}()

	_, err = actor.CallTimeout("hello", func(actor *GActor[any], v ...interface{}) (interface{}, any, error) {
		time.Sleep(80 * time.Millisecond)
		return "ok", v[len(v)-1], nil
	}, 10*time.Millisecond)
	if err == nil {
		t.Fatal("expect timeout error")
	}
}

func TestStopRemovesActor(t *testing.T) {
	actor, err := Start[any]("test-stop-removes-actor", nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	actor.StopTimeout(nil, 200*time.Millisecond)

	deadline := time.Now().Add(500 * time.Millisecond)
	for time.Now().Before(deadline) {
		if _, ok := getAnyActorStatus(actor.key); !ok {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("actor should be removed after stop")
}

// Test typed actor Stop/GetState — this was the broken path before the refactor
type testState struct {
	name  string
	count int
}

func TestTypedActorStopAndGetState(t *testing.T) {
	actor, err := Start(&testState{name: "typed", count: 42}, func(actor *GActor[*testState], v ...interface{}) (*testState, error) {
		return &testState{name: "initialized", count: 100}, nil
	})
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}

	state := GetState(actor)
	if state.name != "initialized" || state.count != 100 {
		t.Fatalf("unexpected state: %+v", state)
	}

	status := GetStatus(actor)
	if status == "actor not found" {
		t.Fatal("GetStatus should find the actor")
	}

	// Stop via typed API
	actor.Stop(nil)

	// Verify removed from actorMap
	deadline := time.Now().Add(500 * time.Millisecond)
	for time.Now().Before(deadline) {
		if _, ok := getAnyActorStatus(actor.key); !ok {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("typed actor should be removed after stop")
}

func TestTypedActorCall(t *testing.T) {
	actor, err := Start(&testState{name: "call-test", count: 0}, func(actor *GActor[*testState], v ...interface{}) (*testState, error) {
		return &testState{name: "call-test", count: 1}, nil
	})
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	defer actor.StopTimeout(nil, 200*time.Millisecond)

	reply, err := actor.Call("increment", func(actor *GActor[*testState], v ...interface{}) (interface{}, *testState, error) {
		s := v[len(v)-1].(*testState)
		s.count++
		return "done", s, nil
	})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if reply != "done" {
		t.Fatalf("unexpected reply: %v", reply)
	}

	state := GetState(actor)
	if state.count != 2 {
		t.Fatalf("expected count 2, got %d", state.count)
	}
}

func TestTypedActorInfo(t *testing.T) {
	actor, err := Start(&testState{name: "info-test", count: 0}, nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	defer actor.StopTimeout(nil, 200*time.Millisecond)

	err = actor.Info("set count", func(actor *GActor[*testState], v ...interface{}) (interface{}, *testState, error) {
		s := v[len(v)-1].(*testState)
		s.count = 99
		return nil, s, nil
	})
	if err != nil {
		t.Fatalf("info failed: %v", err)
	}

	// Give time for async processing
	time.Sleep(50 * time.Millisecond)
	state := GetState(actor)
	if state.count != 99 {
		t.Fatalf("expected count 99, got %d", state.count)
	}
}

func TestTypedActorSendAfter(t *testing.T) {
	actor, err := Start(&testState{name: "timer-test", count: 0}, nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	defer actor.StopTimeout(nil, 500*time.Millisecond)

	timerID, err := actor.SendAfter(50*time.Millisecond, "tick", func(actor *GActor[*testState], v ...interface{}) (interface{}, *testState, error) {
		s := v[len(v)-1].(*testState)
		s.count++
		return nil, s, nil
	})
	if err != nil {
		t.Fatalf("sendAfter failed: %v", err)
	}

	// Wait for timer to be added and check TTL
	time.Sleep(10 * time.Millisecond)
	ttl := actor.GetTimerTTL(timerID)
	if ttl <= 0 {
		t.Fatalf("expected positive TTL, got %v", ttl)
	}

	// Wait for timer to fire
	time.Sleep(150 * time.Millisecond)
	state := GetState(actor)
	if state.count != 1 {
		t.Fatalf("expected count 1 after timer fire, got %d", state.count)
	}
}

func TestStopPackageLevel(t *testing.T) {
	actor, err := Start(&testState{name: "pkg-stop", count: 0}, nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}

	// Stop via package-level function
	Stop(actor)

	// Verify removed
	deadline := time.Now().Add(500 * time.Millisecond)
	for time.Now().Before(deadline) {
		if _, ok := getAnyActorStatus(actor.key); !ok {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("actor should be removed after package-level Stop")
}
