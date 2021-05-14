package maybe

type Maybe interface {
	IfPresent(f func(c *User))
	WhenAbsent(f func())
}

func MaybeUser(user *User) Maybe {
	if user == nil {
		return &Absent{}
	}
	return &Present{u: user}
}

type Present struct {
	u *User
}

func (m *Present) IfPresent(f func(c *User)) {
	f(m.u)
}

func (m *Present) WhenAbsent(f func()) {

}

type Absent struct{}

func (n *Absent) IfPresent(f func(c *User)) {

}

func (n *Absent) WhenAbsent(f func()) {
	f()
}
