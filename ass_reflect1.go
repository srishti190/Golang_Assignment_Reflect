package main

import (
	"fmt"
	"reflect"
)

type Events struct {
	Value string
}
type Contactinfo struct {
	Email  string
	Zip    int
	Events Events
}

var a int = 10

type Student struct {
	Firstname   string
	Lastname    string
	Contactinfo Contactinfo
	Category    map[string]string
	Shape       []string
	Aptr        *int
}

func main() {
	s := Student{
		Firstname: "Jim",
		Lastname:  "Anderson",
		Contactinfo: Contactinfo{
			"jim@gmail.com",
			8600,
			Events{
				"Music",
			},
		},
		Category: map[string]string{
			"red":   "apple",
			"green": "peacock",
			"white": "cloth",
		},
		Shape: []string{"Hi", "Hello"},
		Aptr:  &a,
	}
	bptr := &s
	tabs := " "
	display(bptr, tabs)
}
func display(s interface{}, tabs string) {

	reflectType := reflect.TypeOf(s)
	reflectValue := reflect.ValueOf(s)
	switch reflectType.Kind() {
	case reflect.String:
		fmt.Printf(tabs+"Type:%v Value:%v\n", reflectValue.Kind(), reflectValue.Interface())
	case reflect.Int:
		fmt.Printf(tabs+"Type:%v Value:%v\n", reflectValue.Kind(), reflectValue.Interface())
	case reflect.Struct:
		fmt.Printf("Type:%v\n", reflectValue.Kind())
		fmt.Println("values in struct are:")
		for i := 0; i < reflectValue.NumField(); i++ {
			s := reflectType.Field(i).Name
			if s[0] >= 97 && s[0] <= 123 {
				fmt.Println(tabs + "unexported field")
			} else {
				fmt.Printf(tabs+"Field: %s\n", s)
				display(reflectValue.Field(i).Interface(), tabs+"\t")
			}
		}
	case reflect.Map:
		fmt.Printf(tabs+"Type:%v\n", reflectValue.Kind())
		fmt.Println(tabs + "values in map are:")

		for _, j := range reflectValue.MapKeys() {
			fmt.Printf(tabs+"Key:[%v] ", j)
			display(reflectValue.MapIndex(j).Interface(), tabs+"\t")
		}
	case reflect.Slice, reflect.Array:
		fmt.Printf(tabs+"Type:%v\n", reflectValue.Kind())
		fmt.Println(tabs + "values in slice are:")

		for i := 0; i < reflectValue.Len(); i++ {
			//fmt.Printf(tabs+"value:[%v]\n", reflectValue.Index(i))
			display(reflectValue.Index(i).Interface(), tabs+"\t")
		}
	case reflect.Ptr:
		fmt.Printf(tabs+"Type:%v\n", reflectValue.Kind())
		display(reflectValue.Elem().Interface(), tabs+"\t")
	default:
		fmt.Println(tabs + "empty")
	}
}
