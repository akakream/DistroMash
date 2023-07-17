package controllers

import (
	"strings"
	"testing"
)

const seed = 12

const p0 = ""
const p20 = "QmKhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9,QmBhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9"

var peers = []string{"QmAhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmBhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmChb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmDhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmEhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmFhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmGhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmHhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmJhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                  "QmKhb9o5jV4SzB31gUgBGkgptwWCTqn5AQaT5MzB7bnLV9",
                 }

func TestRandom0PercentOfPeers(t *testing.T) {
    peers0, err := randomNPercentOfPeers(peers, 0, seed)
    if err != nil {
        t.Error(err)
    }
    if arraytoStringTestHelper(peers0) != p0 {
        t.Errorf("peers0: expected %s got: %s", p0, peers0)
    }
}

func TestRandom20PercentOfPeers(t *testing.T) {
    peers20, err := randomNPercentOfPeers(peers, 20, seed)
    if err != nil {
        t.Error(err)
    }
    if arraytoStringTestHelper(peers20) != p20 {
        t.Errorf("peers20: expected %s got: %s", p20, peers20)
    }
}

func arraytoStringTestHelper(p []string) string {
    return strings.Join(p, ",")
}
