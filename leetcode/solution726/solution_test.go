package solution726

import "testing"

func TestCountOfAtoms(t *testing.T) {
	tests := []struct {
		name    string
		formula string
		want    string
	}{
		{"test 1", "H2O", "H2O"},
		{"test 2", "Mg(OH)2", "H2MgO2"},
		{"test 3", "K4(ON(SO3)2)2", "K4N2O14S4"},
		{"test 4", "Be32", "Be32"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %s, want %s", got, tt.want)
			}
		})
	}
}
