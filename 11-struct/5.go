package main

import (
	"encoding/json"
	"fmt"
)

//Use Field Tags in the Definition of Struct Type
type Company struct {
	CompanyName    string `json:"name"`
	OfficeLocation string `json:"location"`
	AdminPhone     string `json:"phone"`
}

func main() {
	//Use Field Tags in the Definition of Struct Type
	//raw string literals vs interpreted string literal "\n\t"
	jsonString := `
    {
        "name": "Cybage",
        "location": "Kalyaninagar",
        "phone": "02066041700"
    }`

	var c1 Company
	json.Unmarshal([]byte(jsonString), &c1)
	fmt.Printf("c1 = %+v\n", c1)

	c2 := Company{"google", "California", "16502530000"}
	jsonStr, err := json.Marshal(c2)
	if err == nil {
		fmt.Printf("c2 to json = %s\n", jsonStr)
	}

}
