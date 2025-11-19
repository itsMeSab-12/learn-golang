package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Ledger")

	/* -------------------------------- User Test ------------------------------- */
	// user1 := ledger.NewUser("Alice")
	// fmt.Println("user 1:", user1)
	// user2 := ledger.NewUser("Bob")
	// fmt.Println("user 2:", user2)
	// user3 := ledger.NewUser("Charlie")
	// fmt.Println("user 3:", user3)

	// // List all users
	// fmt.Println("All Users:")
	// for _, user := range ledger.ListUsers() {
	// 	fmt.Printf("ID: %s, Name: %s\n", user.Id, user.Name)
	// }

	// // Get a specific user
	// fmt.Println("\nGet User by ID:")
	// foundUser := ledger.GetUser(user2.Id)
	// if foundUser != nil {
	// 	fmt.Printf("Found User: ID: %s, Name: %s\n", foundUser.Id, foundUser.Name)
	// } else {
	// 	fmt.Println("User not found.")
	// }

	// // Delete a user
	// fmt.Println("\nDeleting user:", user1.Name)
	// ledger.DeleteUser(user1.Id)

	// // List users again
	// fmt.Println("\nUsers after deletion:")
	// for _, user := range ledger.ListUsers() {
	// 	fmt.Printf("ID: %s, Name: %s\n", user.Id, user.Name)
	// }

}
