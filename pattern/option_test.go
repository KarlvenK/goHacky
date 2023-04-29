package pattern

import "testing"

func TestOption(t *testing.T) {
	server, err := NewServer(
		"127.0.0.1",
		"7890",
		WithLocation("Foo"),
		WithName("Proxy"),
	)
	if err == nil {
		t.Log("server info:\n", server.ServerDetail())
	} else {
		t.Errorf("nooooooo: %v", err)
	}
}
