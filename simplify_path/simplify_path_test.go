package simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Case 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "Case 2",
			path: "//home//foo/",
			want: "/home/foo",
		},
		{
			name: "Case 3",
			path: "/a//b////c/d//././/..",
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
