package gosayhello

import "strconv"

func SayHello(name string, age int) string  {
	return "helo " + name + " saya berumur " + strconv.Itoa(age) + " tahun"
}