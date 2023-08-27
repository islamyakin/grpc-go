package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/islamyakin/grpc-go/model/model"
	"os"
)

var user1 = &model.User{
	Id:       "u01",
	Name:     "Islam Yakin",
	Password: "11002233",
	Gender:   model.UserGender_MALE,
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Rungkut",
	Coordinate: &model.GarageCoordinate{
		Latitude:  23.111111,
		Longitude: 25.00000,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

func main() {

	fmt.Printf("Original\n	%#v \n", user1)
	fmt.Printf("As String\n	%v	\n", user1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("As JSON String\n		%v \n", jsonString)

	protoObject := new(model.GarageList)
	err2 := jsonpb.UnmarshalString(jsonString, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Printf("As String\n %v \n", protoObject.String())
}
