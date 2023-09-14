package models

type Profile struct {
	Name string `json:"name" bson:"name"`
}

type ProfileResponse struct {
	StringMessage string `json:"stringmessage" bson:"stringmessage"`
}
