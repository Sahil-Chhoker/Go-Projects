package main


// DO NOT RUN(ONLY FOR EXPLANATION)
func main() {
	//defer keyword delays the function written after it
	//until the functions returns something

	//EX:
	func logAndDelete(users map[string]user, name string) (log string){
		defer delete(users, name) // this defer func get called -->
		users, ok := users[name]

		if !ok{
			// --> here
			return logNotFound
		}

		if user.admin{
			// --> and here
			return logAdmin
		}

		// --> and here
		return logDeleted
	}


}
