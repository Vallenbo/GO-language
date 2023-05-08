package main

import "time"

func SingleFieldValidate() {
	v := valiate
	var err error

	var b bool
	err = v.Var(b, "boolean")
	outRes("boolean", &err)

	var i = "100"
	err = v.Var(i, "number")
	outRes("number", &err)

	var f = "100.123"
	err = v.Var(f, "numeric")
	outRes("numeric", &err)

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	err = v.Var(slice, "max=15,min=2")
	outRes("slice", &err)

	var mp = make(map[int]int)
	mp[1] = 1
	mp[2] = 2
	mp[3] = 3
	mp[4] = 4
	err = v.Var(mp, "max=15,min=2")
	outRes("mp", &err)

	var timeStr = time.Now().Format("2006-01-02 15:04:05")
	err = v.Var(timeStr, "datetime=2006-01-02 15:04:05")
	outRes("datetime", &err)

	s1 := "abc"
	s2 := "abc"
	err = v.VarWithValue(s1, s2, "eqfield")
	outRes("eqfield", &err)
}