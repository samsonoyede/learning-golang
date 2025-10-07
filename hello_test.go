package hello

// tests for hello package

//import testing package
import "testing"

// t pointer is for raising errors
func TestSpeak(t *testing.T) {
	//creating a slice of structs
	// Interesting logic
	subtests := []struct {
		items  []string
		result string
	}{

		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Gideon"},
			result: "Hello, Gideon!",
		},
		{
			items:  []string{"Wisdom", "Cherry"},
			result: "Hello, Wisdom, Cherry!",
		},
	}

	// For loop, looping through subtests
	for _, st := range subtests {
		if s := SpeakNames(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
