package offer05

import "testing"

var tests = []struct {
	name string
	s    string
	want string
}{
	{"case 1", "We are happy.", "We%20are%20happy."},
}

func TestReplaceSpace1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace1(tt.s); got != tt.want {
				t.Errorf("replaceSpace1 () = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceSpace2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace2(tt.s); got != tt.want {
				t.Errorf("replaceSpace1 () = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceSpace1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceSpace1("We are happy.")
	}
}

func BenchmarkReplaceSpace2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceSpace2("We are happy.")
	}
}
