package model

import (
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	sessionToken := SessionToken{ExpiresAt: time.Now().AddDate(0, 1, 0)}

	if sessionToken.IsValid() != true {
		t.Error("New session token returned but not valid", sessionToken.ExpiresAt)
	}

	sessionToken.ExpiresAt = time.Now().AddDate(0, -1, 0)
	if sessionToken.IsValid() != false {
		t.Error("Expired token returns true", sessionToken.ExpiresAt)
	}
}
