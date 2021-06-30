package entity

type User struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"min=1,max=10,required"`
	Age  int    `json:"age" binding:"gte=1,lte=60"`
}
