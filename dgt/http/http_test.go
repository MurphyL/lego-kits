package http

import (
	"testing"
)

func TestRest(t *testing.T) {
	agent := RequestAgent{}
	agent.Apply("")
	t.Log(agent)
}
