package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// a := "hello"
	u, _ := json.Marshal(DeleteUserActivitiesRequest{UserId: "1", Activity: "aa"})
	fmt.Println(string(u), "###")
	var v DeleteUserActivitiesRequest
	json.Unmarshal(u, &v)

	fmt.Println("number Of Empty Field --", numberOfEmptyField(v))
	// fmt.Println(v.UserId)
	// fmt.Println(v.Activity == "")
	// fmt.Println(v.Description)

	// f

	//  u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	//     if err != nil {
	//         panic(err)
	//     }
	//     fmt.Println(string(u)) // {"Name":"Bob","Age":10,"Active":true}

	// a, _ := json.Marshal(28192)
	// fmt.Println(string(a)) // 20192
	// json.Marshal(`{"UserId":"123"}`,DeleteUserActivitiesRequest)
}

type DeleteUserActivitiesRequest struct {
	UserId      string            `json:"userId,omitempty"`
	Activity    string            `json:"activity"`
	Description map[string]string `json:"description,omitempty"`
	Timestamp   string            `json:"timestamp,omitempty"`
}
type User struct {
	Name        string
	Age         int
	Active      bool
	lastLoginAt string
}

// delete it
// mt.Println(DeleteUserActivitiesRequest{UserId: a})

// uu := DeleteUserActivitiesRequest{UserId: a}

// keys := reflect.ValueOf(uu).FieldByIndex([1])

// fmt.Println(keys) // [a b c]

// for k := range  u{
// 	fmt.Println(u[k])
// }

func numberOfField(u DeleteUserActivitiesRequest) (int,int) {
	totalNumOfFeild := reflect.ValueOf(u).NumField()
	emptyField := 0

	for i := 0; i < totalNumOfFeild; i++ {
		value := fmt.Sprint(reflect.ValueOf(u).Field(i))
		if value == "" {
			emptyField = emptyField + 1
		}
	}

	return totalNumOfFeild,emptyField
}
