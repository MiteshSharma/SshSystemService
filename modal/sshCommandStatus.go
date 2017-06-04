package modal

type SSHCommandStatus struct {
	Id          string 	 `json:"id" bson:"_id"`
	CommandId   string  	 `json:"commandId" bson:"commandId"`
	UserId      string  	 `json:"userId" bson:"userId"`
	Status      string  	 `json:"status" bson:"status"`
	Message     string	 `json:"message" bson:"message"`
}
