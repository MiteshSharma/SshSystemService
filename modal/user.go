package modal

type User struct {
	Id             string `json:"id" bson:"_id"`
	userId         string `json:"userId" bson:"userId"`
	name           string `json:"name" bson:"name"`
	email          string `json:"email" bson:"email"`
	password       string `json:"password" bson:"password"`
}

func (u User) IsValid() bool {
	if (u.email != "" && u.userId != "") {
		return true
	}
	return false
}