package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100] //в таком варианте строка justString занимает в памяти объем равный v
	justString = string([]rune(v[:100])) // здесь мы создаем срез и копируем его в новую строку
}
func main() {
	someFunc()
}
