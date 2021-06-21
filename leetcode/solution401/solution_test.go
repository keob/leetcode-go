package solution401

import (
	"reflect"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	tests := []struct {
		name     string
		turnedOn int
		want     []string
	}{
		{"test 1", 1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := readBinaryWatch(tt.turnedOn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
