package maybe

import "testing"

func TestMaybe(t *testing.T) {
	t.Run("testing maybe", func(t *testing.T) {
		var greeting string
		t.Run("User present", func(t *testing.T) {
			MaybeUser(getUser(1)).IfPresent(func(u *User) {
				greeting = "Hello " + u.name
			})
			if greeting == "" {
				t.Errorf("Expected: Hello [name], Actual: " + greeting)
			}
		})

		t.Run("User absent", func(t *testing.T) {
			MaybeUser(getUser(-1)).WhenAbsent(func() {
				greeting = ""
			})
			if greeting != "" {
				t.Errorf("Expected: Hello [name], Actual: " + greeting)
			}
		})
	})
}
