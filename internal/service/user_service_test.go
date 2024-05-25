package service

import (
	"encoding/base64"
	"testing"

	"github.com/mystpen/Pet-API/internal/model"
)

func TestCreateToken(t *testing.T) {
	userService := NewUserService(nil)
	
	tests := []struct {
		name string
		user *model.User
		want string
	}{
		{
			name: "Sample",
			user: &model.User{
				UserName:      "dana",
				PlainPassword: "123456789",
			},
			want: base64.StdEncoding.EncodeToString([]byte("dana:123456789")),
		},
		{
			name: "Empty",
			user: &model.User{},
			want: base64.StdEncoding.EncodeToString([]byte(":")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := userService.CreateToken(tt.user)

			if res != tt.want {
				t.Errorf("want %q; got %q", tt.want, res)
			}
		})
	}
}
