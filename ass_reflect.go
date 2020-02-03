package main

import (
	"fmt"
	"reflect"
)

// type Student struct {
// 	Fname  string
// 	Lname  string
// 	City   string
// 	Mobile int64
// }

type Contactinfo struct {
	Email string
	Zip   int
}

/*type Events struct {
    cnt   int
    event map[string]Event
}

type Event struct {
    value int64
}*/
/*type Events struct {
	Value string
}*/
type Student struct {
	Firstname   string
	Lastname    string
	Contactinfo Contactinfo
	//Category map[string]string
	Category map[string]string
	Shape    []string
}

func main() {
	//s := Student{"Chetan", "Kumar", "Bangalore", 7777777777}
	s := Student{
		Firstname: "Jim",
		Lastname:  "Anderson",
		Contactinfo: Contactinfo{
			"jim@gmail.com",
			8600,
		},
		Category: map[string]string{
			"red":   "kfdjg",
			"green": "kfdjg",
			"white": "jgh",
		},
		Shape: []string{"Hi", "Hello"},
	}
	display2(s)

	/*v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			fmt.Printf("Type: %v \t Field: %s \t Value: %v \n", v.Field(i).Type(), t.Field(i).Name, v.Field(i).Interface())
			fmt.Println("-------------")
		case reflect.Map:
			fmt.Printf("Type: %v \t Field: %s \n", v.Field(i).Type(), t.Field(i).Name)
			for i, j := range s.Category {
				fmt.Printf("     key:[%s] value:[%s]\n", i, j)
			}
			fmt.Println("-------------")
		case reflect.Slice:
			fmt.Printf("Type: %v \t Field: %s \n", v.Field(i).Type(), t.Field(i).Name)
			for i, j := range s.Shape {
				fmt.Printf("     Index:[%v] value:[%v]\n", i, j)
			}
			fmt.Println("-------------")
		case reflect.Struct:
			vi := reflect.ValueOf(s.Contactinfo)
			ti := reflect.TypeOf(s.Contactinfo)
			for i := 0; i < vi.NumField(); i++ {
				fmt.Printf("Type: %v \t Field: %s \t Value: %v \n", vi.Field(i).Type(), ti.Field(i).Name, vi.Field(i).Interface())
			}
			fmt.Println("-------------")
		}
		//fmt.Printf("Type: %v \t Field: %s \t Value: %v \n", v.Field(i).Type(), t.Field(i).Name, v.Field(i).Interface())
	}*/
}
func display(s interface{}) {

	reflectType := reflect.TypeOf(s)
	reflectValue := reflect.ValueOf(s)
	//typeName := reflectType.Name
	//valueType := reflect.ValueOf(s).Type()
	//switch reflectValue.Kind() {
	switch reflectType.Kind() {
	case reflect.String:
		fmt.Printf("%s : %v(%v)\n", reflectType.Name, reflectValue.Kind(), reflectValue.Interface())
		fmt.Println("in string")
	case reflect.Int:
		fmt.Printf("%s : %v(%v)\n", reflectType.Name, reflectValue.Kind(), reflectValue.Interface())
	case reflect.Struct:
		//fmt.Printf("%s : %v(%v)\n", typeName, valueType, reflectValue.Field(i).Interface())
		display2(reflectValue.Interface())
		//fmt.Println("in struct")
	case reflect.Map:
		fmt.Printf("%s : %v(%v)\n", reflectType.Name, reflectValue.Type(), reflectValue.Interface())
	case reflect.Slice:
		fmt.Printf("%s : %v(%v)\n", reflectType.Name, reflectValue.Type(), reflectValue.Interface())
	default:
		fmt.Println("empty")
	}
}
func display2(s interface{}) {

	reflectType := reflect.TypeOf(s)
	reflectValue := reflect.ValueOf(s)
	for i := 0; i < reflectValue.NumField(); i++ {
		typeName := reflectType.Field(i).Name

		valueType := reflect.ValueOf(s).Field(i).Type()
		//valueValue := reflectValue.Field(i).Interface()
		//v := valueValue
		switch reflectValue.Field(i).Kind() {
		case reflect.String:
			fmt.Printf("%s : %v(%v)\n", typeName, valueType, reflectValue.Field(i).Interface())
		case reflect.Int:
			fmt.Printf("%s : %v(%v)\n", typeName, valueType, reflectValue.Field(i).Interface())
		case reflect.Struct:
			//fmt.Printf("%s : %v(%v)\n", typeName, valueType, reflectValue.Field(i).Interface())
			display(reflectValue.Field(i).Interface())
			//fmt.Println("in struct")
		case reflect.Map:
			fmt.Printf("%s :%v(%v)\n", typeName, valueType, reflectValue.Field(i).Interface())
		case reflect.Slice:
			fmt.Printf("%s :%v(%v)\n", typeName, valueType, reflectValue.Field(i).Interface())
		default:
			fmt.Println("empty")
		}
	}
}
