package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 编写代码利用反射实现一个ini文件的解析器程序。
// 根据ini文件的key，输出key对应的value

type MysqlConfig struct {
	Ip       string `ini:"ip"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
}

type Config struct {
	MysqlConfig `ini:"mysqld"`
}

func loadIni(filename string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct { // 1. 参数校验，参数必须是指针类型和结构体类型
		err = errors.New("type error should be struct ptr ")
		return err
	}

	file, err := os.ReadFile(filename) // 2. 打开文件
	if err != nil {
		err = fmt.Errorf("open file error %v\n", err)
		return err
	}

	lineSlice := strings.Split(string(file), "\r\n") // 3. 一行一行读取文件
	var section string                               // 定义section切片
	var structName string

	for idx, line := range lineSlice { // 遍历slice
		line = strings.TrimSpace(line)                                                      // 去除首尾空格
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 { // 3.1 如果是注释，则跳过 \\如果是空行，则继续
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] == '[' && line[len(line)-1] == ']' && len(line) > 2 { // 3.2 如果开头是[，而且是以]结尾的，且[]不为空，则被认定是section

				section = line[1 : len(line)-1] // 将section 加入到section切片变量里
			}
			for i := 0; i < t.Elem().NumField(); i++ { //根据字符串section去data里面根据反射找到对应的结构体
				field := t.Elem().Field(i)
				if section == field.Tag.Get("ini") {
					structName = field.Name
				}
			}
		} else { // 4. 拆分键值对
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") { // 4.1 以等号分割，左边为key，右边为value
				err = fmt.Errorf("line:%d syntax error ", idx+1)
				return
			}
			index := strings.Index(line, "=")          // 获取到line的key和value
			key := strings.TrimSpace(line[:index])     // user
			value := strings.TrimSpace(line[index+1:]) //mysql

			// 4.2 将ini文件中的section和结构体的名字对应
			v := reflect.ValueOf(data)                 // interface的值,此时为空
			sValue := v.Elem().FieldByName(structName) // 结构体的值 { 10.0.0.1 mysql abc123 0}
			sType := sValue.Type()                     // 结构体的类型 main.MysqlConfig

			if sType.Kind() != reflect.Struct { // 判断结构体的字段是否是struct类型
				err = fmt.Errorf("should be struct")
				return err
			}
			// 根据structName 去data里面把嵌套的结构体取出
			// 声明字段值
			var fname string                  // 结构体的字段名
			var fieldType reflect.StructField // 结构体字段类型
			for i := 0; i < sValue.NumField(); i++ {
				fieldName := sType.Field(i)
				fieldType = fieldName
				if fieldName.Tag.Get("ini") == key { //  遍历结构体的每一个字段，判断这个tag是不是等于key
					fname = fieldName.Name // 找到对应的字段
					break
				}
			}

			fieldObj := sValue.FieldByName(fname) // 4.3如果key == tag，and value的类型等于结构体定义的类型，则赋值
			switch fieldType.Type.Kind() {        // 判断ini文件中value的种类，字串/数字/布尔
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				fieldObj.SetInt(valueInt)

			case reflect.Bool:
				var valueBool bool
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	var config Config
	err := loadIni("./config.ini", &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(
		config.MysqlConfig.Ip,
		config.MysqlConfig.Password,
		config.MysqlConfig.Port,
		config.MysqlConfig.User,
	)
}
