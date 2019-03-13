package main

import (
	"reflect"
	"fmt"
	"encoding/json"
)
type User struct {
	Name  string `name`
	Age   int `age`
	class string `class`
}

func (user User) GetName() string {
	return user.Name
}

func (u User) Print(prfix string){
	fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}


func main() {
	u := User{"张三", 20, "high"}
	t := reflect.TypeOf(u)
	fmt.Println(t)
	v := reflect.ValueOf(u)
	fmt.Println(v)
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)
	u1 := v.Interface().(User)
	fmt.Println(u1)
	t1 := v.Type()
	fmt.Println(t1)
	t2 := reflect.TypeOf(u1.Age)
	fmt.Println(t2)
	fmt.Println(v.Kind())
	fmt.Println(t.Kind())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	mPrint:=v.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))

	mGetName:=v.MethodByName("GetName")
	args =[]reflect.Value{}
	fmt.Println(mGetName.Call(args))

	p := reflect.ValueOf(&u)
	newValue := p.Elem()
	fmt.Println("type of pointer:", newValue.Field(0))
	fmt.Println("settability of pointer:", newValue.Field(0).CanSet() )
	newValue.Field(0).Set(reflect.ValueOf("李四"))
	fmt.Println(u.Name)
	newValue.Field(0).SetString("王五")
	fmt.Println(u.Name)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name )
	}

	h:=`{"name":"洪六","age":36}`
	err:=json.Unmarshal([]byte(h),&u)
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(u)
	}

}

//https://www.flysnow.org/2017/06/13/go-in-action-go-reflect.html
//https://juejin.im/post/5a75a4fb5188257a82110544
//https://riot.im/app