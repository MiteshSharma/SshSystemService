package modal

type Employee struct {
	Id             string `json:"id" bson:"_id"`
	name           string `json:"name" bson:"name"`
	email          string `json:"email" bson:"email"`
	key            string `json:"key" bson:"key"`
}

func (e Employee) IsValid() bool {
	if (e.email != "" && e.key != "") {
		return true
	}
	return false
}