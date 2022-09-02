package urlencoder

import (
	"testing"
)

func TestGetEncodedUrlPath(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "test simple path",
			url:  "https://test.com/a",
			want: "a",
		},
		{
			name: "test multiple path",
			url:  "https://test.com/a/b/c",
			want: "a",
		},
		{
			name: "test query string",
			url:  "https://test.com/abc/b/c?q=1",
			want: "abc",
		},
		{
			name: "empty path",
			url:  "https://test.com",
			want: "",
		},
		{
			name: "slash only path",
			url:  "https://test.com/",
			want: "",
		},
		{
			name: "invalid url",
			url:  "daacaa2",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEncodedUrlPath(tt.url); got != tt.want {
				t.Fatalf("GetEncodedUrlPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeUrlPath(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		wantEmpty bool
	}{
		{
			name:      "test simple path",
			url:       "https://test.com/a",
			wantEmpty: false,
		},
		{
			name:      "test multiple path",
			url:       "https://test.com/a/b/c",
			wantEmpty: false,
		},
		{
			name:      "empty path",
			url:       "",
			wantEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeUrlPath(tt.url); (len(got) == 0) != tt.wantEmpty {
				t.Errorf("EncodeUrlPath() = url path length %v, want empty? %v", len(got), tt.wantEmpty)
			}
		})
	}
}
