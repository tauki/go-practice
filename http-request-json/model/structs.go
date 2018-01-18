package model

type Category struct {
	ID       int    	`json:"id"`
	Name     string 	`json:"name"`
	Active   bool   	`json:"active"`
	Status   int    	`json:"status"`
	BangName string 	`json:"bang_name"`
	Type     int    	`json:"type"`
	Sub 	 []Category	`json:"sub_categories"`

	Detail	 string		`json:"detail"`
}