package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestLoadConfig(t *testing.T) {
  is := is.New(t)
  s := newServer()
  empty := config{}
  cfg := s.loadConfig()
  if cfg == empty {
    is.Fail()
  }
}