package tui

import (
	"fmt"
	"math/rand/v2"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

var demoFirstNames = []string{
	"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank",
	"Iris", "Jack", "Kate", "Leo", "Maya", "Nick", "Olivia", "Pete",
	"Quinn", "Rose", "Sam", "Tina", "Uma", "Vic", "Wendy", "Xander",
	"Yuki", "Zara", "Axel", "Bella", "Cruz", "Dex", "Ember", "Flynn",
	"Gemma", "Hugo", "Ivy", "Jade", "Knox", "Luna", "Milo", "Nova",
	"Opal", "Piper", "Reed", "Sage", "Theo", "Vale", "Wren", "Zoe",
	"Ash", "Bryn",
}

var demoLastNames = []string{
	"Wonder", "Smith", "Rose", "King", "Park", "Moon", "Fox", "Stone",
	"Blake", "Cruz", "Drake", "Frost", "Gray", "Hart", "Lake", "Mars",
	"Noble", "Quinn", "Rain", "Snow", "Vale", "West", "York", "Zen",
	"Storm",
}

var demoGroupNames = []string{
	"Builders United", "Scripters Hub", "Fashion Circle", "SL Explorers",
	"Music Lovers", "Racing League", "Art Gallery Collective", "Virtual Architects",
	"Photography Society", "Dance Club International", "Sailing Community",
	"Aviation Enthusiasts", "Roleplay Alliance", "Trading Post", "Newcomers Welcome",
}

func generateDemoFriends(count int) []sl.Friend {
	friends := make([]sl.Friend, count)
	for i := range count {
		first := demoFirstNames[rand.IntN(len(demoFirstNames))]
		last := demoLastNames[rand.IntN(len(demoLastNames))]
		online := rand.IntN(100) < 40

		friends[i] = sl.Friend{
			DisplayName:  fmt.Sprintf("%s %s", first, last),
			InternalName: fmt.Sprintf("%s.%s", toLower(first), toLower(last)),
			Online:       online,
		}
	}
	return friends
}

func generateDemoGroups(count int) []sl.Group {
	groups := make([]sl.Group, count)
	for i := range count {
		name := demoGroupNames[i%len(demoGroupNames)]
		if i >= len(demoGroupNames) {
			name = fmt.Sprintf("%s %d", name, i/len(demoGroupNames)+1)
		}
		groups[i] = sl.Group{
			Name:        name,
			MemberCount: fmt.Sprintf("%d", rand.IntN(500)+5),
		}
	}
	return groups
}

func toLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	return string(b)
}
