package coverage

import (
	"os"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW



func TestLen(t *testing.T) {
	var people People

	if people.Len() != 0 {
		t.Errorf("Wrong People Len empty")
	}

	people = append(people, Person{})
	people = append(people, Person{})
	people = append(people, Person{})

	if people.Len() != 3 {
		t.Errorf("Wrong People Len 3")
	}

	people = people[0:1]

	if people.Len() != 1 {
		t.Errorf("Wrong People Len 1")
	}
}

func TestLess(t *testing.T) {
	var people People
	now := time.Time{}

	people = append(people, Person{"BBB", "BBB", now})
	people = append(people, Person{"BBB", "BBB", now.Add(5 * time.Minute)})
	people = append(people, Person{"AAA", "BBB", now})
	people = append(people, Person{"AAA", "AAA", now})
	people = append(people, Person{"AAA", "AAA", now})

	if people.Less(0, 1) {
		t.Errorf("Wrong People Less by Birthday")
	}

	if people.Less(0, 2) {
		t.Errorf("Wrong People Less by FirstName")
	}

	if people.Less(2, 3) {
		t.Errorf("Wrong People Less by LastName")
	}

	if people.Less(3, 4) {
		t.Errorf("Wrong People Less Eq")
	}
}

func TestSwap(t *testing.T) {
	var people People
	now := time.Time{}

	people1 := Person{"first", "first", now}
	people2 := Person{"first", "first", now}

	people = append(people, people1)
	people = append(people, people2)

	people.Swap(0, 1)

	if people[0] != people2 || people[1] != people1 {
		t.Errorf("Wrong People Swap")
	}
}

func TestNew(t *testing.T) {

}

func TestRows(t *testing.T) {

}

func TestCols(t *testing.T) {

}

func TestSet(t *testing.T) {

	
}
