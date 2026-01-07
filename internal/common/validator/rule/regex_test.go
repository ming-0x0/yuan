package rule

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "test@example.com", false},
		{"valid email with dots", "user.name@domain.co.uk", false},
		{"valid email with plus", "user+mailbox@sub.domain.com", false},
		{"invalid: no @", "test", true},
		{"invalid: no domain", "test@", true},
		{"invalid: starting with @", "@domain.com", true},
		{"invalid: no TLD", "test@domain", true},
		{"invalid: short TLD", "test@domain.c", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			rule := IsEmail(tt.email)
			err := rule.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmail(%v).Validate() error = %v, wantErr %v", tt.email, err, tt.wantErr)
			}
		})
	}
}
