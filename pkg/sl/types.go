package sl

// Friend represents a Second Life user in the friends list.
type Friend struct {
	DisplayName  string
	InternalName string
	Online       bool
}

// Group represents a Second Life group.
type Group struct {
	Name        string
	MemberCount string
}

// Lindens represents the L$ balance.
type Lindens struct {
	Balance string
}
