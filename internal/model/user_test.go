package model

import (
	"strings"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name      string
		email     string
		expectErr error
	}{
		"正常系": {
			name:      "test",
			email:     "test@example.com",
			expectErr: nil,
		},
		"準異常系: nameが空": {
			name:      "",
			email:     "test@example.com",
			expectErr: ErrInvalidName,
		},
		"準異常系: nameが50文字を超える": {
			name:      strings.Repeat("a", 51),
			email:     "test@example.com",
			expectErr: ErrInvalidName,
		},
		"準異常系: emailに@が含まれない": {
			name:      "test",
			email:     "testexample.com",
			expectErr: ErrInvalidEmail,
		},
	}

	for name, tt := range testCases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if _, err := NewUser(tt.name, tt.email); err != tt.expectErr {
				t.Errorf("want %v, got %v", tt.expectErr, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name      string
		expectErr error
	}{
		"正常系": {
			name:      "test",
			expectErr: nil,
		},
		"準異常系: nameが空": {
			name:      "",
			expectErr: ErrInvalidName,
		},
		"準異常系: nameが50文字を超える": {
			name:      strings.Repeat("a", 51),
			expectErr: ErrInvalidName,
		},
	}

	for name, tt := range testCases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u := &User{"id", "name", "name@example.com", time.Now()}
			if err := u.UpdateName(tt.name); err != tt.expectErr {
				t.Errorf("want %v, got %v", tt.expectErr, err)
			}
			if u.Name != tt.name {
				t.Errorf("want %s, got %s", tt.name, u.Name)
			}
		})
	}
}

func TestUpdateEmail(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		email     string
		expectErr error
	}{
		"正常系": {
			email:     "test@exmple.com",
			expectErr: nil,
		},
		"準異常系: emailに@が含まれない": {
			email:     "testexample.com",
			expectErr: ErrInvalidEmail,
		},
	}

	for name, tt := range testCases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u := &User{"id", "name", "name@example.com", time.Now()}
			if err := u.UpdateEmail(tt.email); err != tt.expectErr {
				t.Errorf("want %v, got %v", tt.expectErr, err)
			}
			if u.Email != tt.email {
				t.Errorf("want %s, got %s", tt.email, u.Email)
			}
		})
	}
}
