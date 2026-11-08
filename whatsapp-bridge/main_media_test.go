package main

import "testing"

func TestExtractDirectPathFromURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{"legacy", "https://mmg.whatsapp.net/v/t62/file.enc?ccb=11-4&oh=signature", "/v/t62/file.enc?ccb=11-4&oh=signature"},
		{"modern", "https://mmg.whatsapp.net/d/f/file.enc?auth=signature", "/d/f/file.enc?auth=signature"},
		{"escaped", "https://mmg.whatsapp.net/v/a%2Fb%20c.enc?oh=a%2fb%2Bc&x=1&x=2&empty=&plus=a+b", "/v/a%2Fb%20c.enc?oh=a%2fb%2Bc&x=1&x=2&empty=&plus=a+b"},
		{"https port", "https://mmg.whatsapp.net:443/v/file.enc?oh=signature", "/v/file.enc?oh=signature"},
		{"malformed", "https://mmg.whatsapp.net/%zz?oh=secret", ""},
		{"http", "http://mmg.whatsapp.net/v/file.enc?oh=secret", ""},
		{"wrong host", "https://example.com/v/file.enc?oh=secret", ""},
		{"host suffix", "https://mmg.whatsapp.net.example.com/v/file.enc?oh=secret", ""},
		{"userinfo", "https://user:secret@mmg.whatsapp.net/v/file.enc?oh=secret", ""},
		{"wrong port", "https://mmg.whatsapp.net:8443/v/file.enc?oh=secret", ""},
		{"invalid port", "https://mmg.whatsapp.net:bad/v/file.enc?oh=secret", ""},
		{"fragment", "https://mmg.whatsapp.net/v/file.enc?oh=secret#fragment", ""},
		{"empty fragment", "https://mmg.whatsapp.net/v/file.enc?oh=secret#", ""},
		{"no path", "https://mmg.whatsapp.net?oh=secret", ""},
		{"root path", "https://mmg.whatsapp.net/?oh=secret", ""},
		{"no query", "https://mmg.whatsapp.net/v/file.enc", ""},
		{"empty query", "https://mmg.whatsapp.net/v/file.enc?", ""},
		{"relative", "/v/file.enc?oh=secret", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractDirectPathFromURL(tt.url)
			if tt.want == "" {
				if got != "" || err == nil || err.Error() != "invalid WhatsApp media URL" {
					t.Fatal("expected an empty path and sanitized error")
				}
				return
			}
			if err != nil || got != tt.want {
				t.Fatal("signed request URI was not preserved")
			}
		})
	}
}
