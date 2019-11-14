package datastructure

import "testing"

type Parent struct {
	child1 Child
	Child2 Child
}

type Child struct {
	grandchild1 string
	grandchild2 int
	grandchild3 GrandChild
}

type GrandChild struct {
	grandChild1 string
}

// Test_ValidateStructValueNormal : Normal case
func Test_ValidateStructValueNormal(t *testing.T) {
	emptyField, err := ValidateStructValue("", &Parent{
		child1: Child{"AKW Keren", 10, GrandChild{"Test"}},
		Child2: Child{"AKW Keren2", 100, GrandChild{"Test"}},
	})
	if err != nil {
		t.Logf("Error happened: %#v", err)
	}

	if emptyField != "" {
		t.Errorf("Don't expect empty field has value but it does")
	}
}

// Test_ValidateStructValueNormal : Normal case
func Test_ValidateStructValueEmptyField(t *testing.T) {
	emptyField, err := ValidateStructValue("", &Parent{})
	if err != nil {
		t.Logf("Error happened: %#v", err)
	}

	if emptyField == "" {
		t.Errorf("Expecting empty field has value but it does not")
	}
}

// Test_ValidateStructValueNormal : Normal case
func Test_ValidateStructValueEmptyField2(t *testing.T) {
	emptyField, err := ValidateStructValue("", &Parent{
		child1: Child{"AKW Keren", 10, GrandChild{"Test"}},
	})
	if err != nil {
		t.Logf("Error happened: %#v", err)
	}

	if emptyField == "" {
		t.Errorf("Expecting empty field has value but it does not")
	}
}

// Test_ValidateStructValueNormal : Normal case
func Test_ValidateStructValueEmptyField3(t *testing.T) {
	emptyField, err := ValidateStructValue("", &Parent{
		child1: Child{"AKW Keren", 10, GrandChild{"Test"}},
		Child2: Child{grandchild1: "AKW Keren", grandchild2: 10},
	})
	if err != nil {
		t.Logf("Error happened: %#v", err)
	}

	if emptyField == "" {
		t.Errorf("Expecting empty field has value but it does not")
	}
}
