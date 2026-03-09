package main

import "fmt"

func main() {

	// Задание 1

	// arr1 := [...]string{"Разминка", "Легкий бег", "Спринт", "Заминка"}
	// arr2 := [...]string{"Медленная ходьба", "Растяжка", "Подготовка к бегу"}

	// fmt.Println("Длина первого массива: ", len(arr1))
	// fmt.Println("Длина второго массива: ", len(arr2))

	// Задание 2

	// subjectsList := [3]string{"Физика", "Химия", "География"}

	// fmt.Printf("Первый элемент массива: %s, последний элемент массива: %s \n", subjectsList[0], subjectsList[2])

	// subjectsList[1] = "Биология"

	// fmt.Println(subjectsList)

	// Задание 3

	// arr := [3]string{"Tom", "35", "New York"}

	// name := arr[0]
	// age := arr[1]
	// home := arr[2]

	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(home)

	// Задание 4

	// numbersList := [5]int{1, 2, 4, 4, 5}

	// isNumber := false

	// for _, number := range numbersList {
	// 	if number == 3 {
	// 		isNumber = true
	// 		break
	// 	}
	// }
	// if isNumber {
	// 	fmt.Println("Число 3 есть в массиве")
	// } else {
	// 	fmt.Println("Число 3 отсутствует в массиве")
	// }

	// Задание 5

	// friendsList := [5]string{"Arima", "Jhon", "Adam", "Bekbolat", "Tomas"}

	// isFriend := false

	// for _, friend := range friendsList {
	// 	if friend == "Bekbolat" {
	// 		isFriend = true
	// 		break
	// 	}
	// }
	// if isFriend {
	// 	fmt.Println("Мне очень повезло")
	// } else {
	// 	fmt.Println("Мне не повезло")
	// }

	// Задание 6

	// firstList := [3]int{1, 2, 3}
	// secondList := [3]int{1, 2, 4}

	// if firstList == secondList {
	// 	fmt.Println("Массивы равны")
	// } else {
	// 	fmt.Println("Массивы не равны")
	// }

	// Задание 7

	// myWishList := [2]string{"ПС", "Драгон хук"}
	// friendsWishList := [2]string{"Одежда", "Аркана"}

	// var registrationList []string

	// for _, v := range myWishList {
	// 	registrationList = append(registrationList, v)
	// }

	// for _, v := range friendsWishList {
	// 	registrationList = append(registrationList, v)
	// }

	// fmt.Println(registrationList)

	// Задание 8

	toyList := [3]string{"Car", "Doll", "Ball"}
	testToyList := toyList

	testToyList[1] = "Boat"

	fmt.Println(toyList)
	fmt.Println(testToyList)

}
