package bootstrap

import (
	"log"
	"os"

	"github.com/Arauco8/go_web/internal/domain"
	"github.com/Arauco8/go_web/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "user: ", log.Lshortfile|log.LstdFlags)
}

func NewDB() user.DB {

	return user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john_doe@gmail.com",
		}, {
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane_doe@gmail.com",
		}, {
			ID:        3,
			FirstName: "Alice",
			LastName:  "Smith",
			Email:     "alice_smith@gmail.com",
		}, {
			ID:        4,
			FirstName: "Bob",
			LastName:  "Johnson",
			Email:     "bob_johnson@gmail.com",
		}, {
			ID:        5,
			FirstName: "Charlie",
			LastName:  "Brown",
			Email:     "charlie_brown@gmail.com",
		}, {
			ID:        6,
			FirstName: "Diana",
			LastName:  "Prince",
			Email:     "diana_prince@gmail.com",
		}, {
			ID:        7,
			FirstName: "Eve",
			LastName:  "Adams",
			Email:     "eve_adams@gmail.com",
		}, {
			ID:        8,
			FirstName: "Frank",
			LastName:  "Miller",
			Email:     "frank_miller@gmail.com",
		}, {
			ID:        9,
			FirstName: "Grace",
			LastName:  "Hopper",
			Email:     "grace_hopper@gmail.com",
		}, {
			ID:        10,
			FirstName: "Hank",
			LastName:  "Pym",
			Email:     "hank_pym@gmail.com",
		}},
		MaxUserID: 3,
	}
}
