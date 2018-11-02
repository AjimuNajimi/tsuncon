package main

import (
  "strings"
	"testing"
)

func TestTsuncon(t *testing.T) {
  msg, err := getContributions("diplozoon")
  if err != nil {
    t.Fatal(err)
  }
  if !strings.Contains(msg, "ξ ﾟ⊿ ﾟ)ξ <") {
    t.Errorf("expected ξ ﾟ⊿ ﾟ)ξ < something, got=%s", msg)
  }
}
