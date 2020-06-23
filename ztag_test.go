package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_fileExists(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "True",
			arg:  "ztag.go",
			want: true,
		},
		{
			name: "False",
			arg:  "Really fake file name",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExists(tt.arg); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dirExists(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "True",
			arg:  "../ztag",
			want: true,
		},
		{
			name: "False",
			arg:  "Really fake folder name",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirExists(tt.arg); got != tt.want {
				t.Errorf("dirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
