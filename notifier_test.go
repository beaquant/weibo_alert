package main

import (
	"net/http"
	"testing"
	"time"
)

func TestClient_PushToWechat(t *testing.T) {
	c := NewClient(http.DefaultClient)
	ScKey = "SCU28704T55ea4ee1b39512b35eb63b36a24caafe5b35de6d732e0"
	c.PushToWechat("123", "456", "789", "hello", time.Now())
	time.Sleep(time.Second)
}
