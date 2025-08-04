package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100] //в таком варианте строка justString занимает в памяти объем равный v что приводит к утечке памяти
	justString = string([]rune(v[:100])) //исправил: cоздаю срез от изначальной строки и копирую его в новую строку, утечки нет
}
func main() {
	someFunc()
}
