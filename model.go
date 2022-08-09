package main

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var UsersData = []User{
	User{1, "Stephanie", "Turner"},
	User{2, "Anna", "Edmunds"},
	User{3, "Jan", "Vaughan"},
	User{4, "Grace", "North"},
	User{5, "Piers", "Morrison"},
}
