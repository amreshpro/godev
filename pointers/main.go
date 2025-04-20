package main

func updateNumber(num *int) {

	*num = *num + 10

}

func main() {
	num := 10

	updateNumber(&num)
	println(num)

	updateNumber(&num)
	println(num)
	updateNumber(&num)
	println(num)
	updateNumber(&num)
	println(num)

}
