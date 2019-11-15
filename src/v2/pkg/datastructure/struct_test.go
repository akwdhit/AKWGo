package datastructure

import (
	"testing"
)

type Parent struct {
	child1 Child
	Child2 Child
}

type Child struct {
	grandchild1 string
	Grandchild2 string
	grandchild3 int
	grandchild4 *GrandChild
	Grandchild5 *GrandChild
	Grandchild6 GrandChild
}

type GrandChild struct {
	grandChild1 string
}

func Test_ValidateStructValue(t *testing.T) {
	tests := []struct {
		name                  string
		parentName            string
		field                 interface{}
		shouldEmptyFieldEmpty bool
		wantErr               bool
	}{
		{
			"normal",
			"",
			&Parent{
				child1: Child{"AKW Keren", "AKW Keren_", 10, &GrandChild{"Test"}, &GrandChild{"Test_"}, GrandChild{"Test"}},
				Child2: Child{"AKW Keren2", "AKW Keren_", 100, &GrandChild{"Test"}, &GrandChild{"Test_"}, GrandChild{"Test"}},
			},
			true,
			false,
		},
		{
			"empty field",
			"",
			&Parent{},
			false,
			true,
		},
		{
			"empty field_2",
			"",
			&Parent{
				child1: Child{"AKW Keren", "AKW Keren_", 10, &GrandChild{"Test"}, &GrandChild{"Test_"}, GrandChild{"Test"}},
			},
			false,
			true,
		},
		{
			"empty field_3",
			"",
			&Parent{
				child1: Child{"AKW Keren", "AKW Keren_", 10, &GrandChild{"Test"}, &GrandChild{"Test_"}, GrandChild{"Test"}},
				Child2: Child{grandchild1: "AKW Keren2", grandchild3: 10},
			},
			false,
			true,
		},
		{
			"empty field_4",
			"",
			&Parent{
				child1: Child{"AKW Keren", "AKW Keren_", 10, &GrandChild{"Test"}, &GrandChild{"Test_"}, GrandChild{"Test"}},
				Child2: Child{grandchild1: "AKW Keren2", Grandchild2: "AKW Test Aja", grandchild3: 10},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyField, err := ValidateStructValue(tt.parentName, tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s - validateStructValue(), field = %#v, emptyField = %#v, error = %#v, wantErr %#v", tt.name, tt.field, emptyField, err, tt.wantErr)
				return
			}
			if emptyField == "" && !tt.shouldEmptyFieldEmpty {
				t.Errorf("%s - validateStructValue() Don't expect empty field has value but it does", tt.name)
			}
		})
	}
}
