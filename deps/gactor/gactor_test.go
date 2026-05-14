package gactor

import (
	"testing"
	"time"
)

func TestStartInitState(t *testing.T) {
	actor, err := Start("test-start-init-state", func(v ...interface{}) (GActorState, error) {
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
	actor, err := Start("test-call-timeout", nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	defer func() {
		actor.StopTimeout(nil, 200*time.Millisecond)
	}()

	_, err = actor.CallTimeout("hello", func(v ...interface{}) (interface{}, GActorState, error) {
		time.Sleep(80 * time.Millisecond)
		return "ok", v[len(v)-1], nil
	}, 10*time.Millisecond)
	if err == nil {
		t.Fatal("expect timeout error")
	}
}

func TestStopRemovesActor(t *testing.T) {
	actor, err := Start("test-stop-removes-actor", nil)
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}
	actor.StopTimeout(nil, 200*time.Millisecond)

	deadline := time.Now().Add(500 * time.Millisecond)
	for time.Now().Before(deadline) {
		if _, ok := getActorStatus(actor.key); !ok {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("actor should be removed after stop")
}
