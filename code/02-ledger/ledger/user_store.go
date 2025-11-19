package ledger

var userDB = make(map[string]*User)

func AddUser(u *User) bool {
	_, exists := userDB[u.ID]
	if !exists {
		userDB[u.ID] = u
		return true
	}
	return false
}

func GetUser(id string) *User {
	user, ok := userDB[id]
	if ok {
		return user
	}
	return nil
}

func ListUsers() []*User {
	list := make([]*User, 0, len(userDB))
	for _, val := range userDB {
		list = append(list, val)
	}
	return list
}

func DeleteUser(id string) {
	delete(userDB, id)
}
