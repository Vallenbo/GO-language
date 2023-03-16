package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Sleep() string { //添加结构体方法，用于打印反射结构体方法
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}
func reflectType1() { //1-类型的基本反射
	var x float32 = 3.14
	v := reflect.TypeOf(x)                      //返回任意值的类型对象
	v1 := reflect.ValueOf(x)                    //返回任意值的类型对象的值
	fmt.Printf("TypeOf:%v,ValueOf:%v\n", v, v1) //TypeOf:float32,ValueOf:3.14
}

func reflectType2(x student) { //1-2 反射类型的底层
	v := reflect.TypeOf(x).Name()                      //返回type关键字构造自定义类型(如cat类型)
	v1 := reflect.TypeOf(x).Kind()                     //返回底层的（引用或基本）类型
	fmt.Printf("TypeOf.Name:%v,TypeOf.Kind:%v", v, v1) //TypeOf.Name:person,TypeOf.Kind:struct
}

func main() {
	var x = student{
		Name:  "沙河小王子",
		Score: 18,
	}
	reflectType1()
	reflectType2(x)

	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
	structReflect(x)
}

func reflectValue(x interface{}) { //2-1 通过反射获取值
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64: // v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32: // v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64: // v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", v.Float())
	}
}

func reflectSetValue2(x interface{}) { //3-1 reflect.ValueOf().Elem().SetInt()通过反射修改值
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 { // 反射中使用 Elem()方法获取指针对应的值
		v.Elem().SetInt(200)
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func structReflect(stu1 student) { //4- 1 结构体反射
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())     // student struct
	for i := 0; i < t.NumField(); i++ { // 通过for循环遍历结构体的所有字段信息
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
			field.Name, field.Index, field.Type, field.Tag.Get("json")) //name:Score index:[1] type:int json tag:score
	}

	if scoreField, ok := t.FieldByName("Score"); ok { // 通过字段名获取指定结构体字段信息
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
			scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json")) //name:Score index:[1] type:int json tag:score
	}
}

func printMethod(x interface{}) { //打印反射结构体方法
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)

		var args []reflect.Value // 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		v.Method(i).Call(args)
	}
}
