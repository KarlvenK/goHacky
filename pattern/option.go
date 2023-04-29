package pattern

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	IP       string
	Port     string
	Name     string
	Location string
}

type ServerOption func(*Server) error

func NewServer(ip, port string, opts ...ServerOption) (*Server, error) {
	if add := net.ParseIP(ip); add == nil {
		return nil, errors.New("invalid ip address")
	}
	i, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("provided port is not an integer: %v", err)
	}
	if i < 0 || i > 65535 {
		return nil, fmt.Errorf("provided port is invalid")
	}
	s := &Server{
		IP:       ip,
		Port:     port,
		Name:     "Null",
		Location: "Null",
	}
	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithName(name string) ServerOption {
	return func(s *Server) error {
		s.Name = name
		return nil
	}
}

func WithLocation(loc string) ServerOption {
	return func(s *Server) error {
		s.Location = loc
		return nil
	}
}

func (s *Server) ServerDetail() string {
	return fmt.Sprintf(
		`{
	IP: %s
	Port: %s
	Name: %s
	Location: %s
}`, s.IP, s.Port, s.Name, s.Location,
	)
}
