package switcher

import "testing"

func TestPerson_String(t *testing.T) {
	type fields struct {
		name string
		age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				name: tt.fields.name,
				age:  tt.fields.age,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	Run()
}
