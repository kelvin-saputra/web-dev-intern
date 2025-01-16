package model

/*
DATA SAMPLE:
	{
		"userId": 1,
		"id": 1,
		"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  	},
*/
type Comment struct {
	UserId int `json:"userId" gorm:"not null"` //This is a foreign key
	Id int `json:"id" gorm:"primaryKey"` //This is a primary key
	Title string `json:"title" gorm:"not null"`
	Body string `json:"body" gorm:"not null"`	
}