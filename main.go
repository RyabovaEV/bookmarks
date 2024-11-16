package main

import (
	"fmt"
)

func main() {
	bookmarks := map[string]string{
		"main":  "Главная",
		"about": "О нас",
	}
	actions := map[int]string{
		1: "output",
		2: "append",
		3: "delete",
		4: "exit",
	}
	//var choiceAction int
	fmt.Println("__Утилита для работы с закладками__")
	for {
		choiceAction := actionSelect(actions)
		switch choiceAction {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			key, value := appendBookmarks()
			bookmarks[key] = value
		case 3:
			deleteKey := deleteBookmarks(bookmarks)
			delete(bookmarks, deleteKey)
		case 4:
			fmt.Println("Выход из программы...")
			return
		}
	}
}

func actionSelect(actions map[int]string) int {
	var choiceAction int
	fmt.Println("Выберите действие:")
	for key, value := range actions {
		fmt.Println(key, "-", value)
	}
	for {
		fmt.Println("Введите № действия: ")
		fmt.Scan(&choiceAction)
		for key := range actions {
			if key == choiceAction {
				return choiceAction
			}
		}
		fmt.Println("!!! Введены не корректные данные, повторите ввод !!!")
		continue
	}
}

func printBookmarks(bookmarks map[string]string) {
	fmt.Println("Список закладок: ")
	for key, value := range bookmarks {
		fmt.Println(key, value)
	}
}

func appendBookmarks() (key string, value string) {
	fmt.Println("Введите ключ: ")
	fmt.Scan(&key)
	fmt.Println("Введите значение: ")
	fmt.Scan(&value)
	return key, value
}

func deleteBookmarks(bookmarks map[string]string) string {
	var deleteKey string
	fmt.Println("Какой элемент хотите удалть?")
	for key, value := range bookmarks {
		fmt.Println(key, value)
	}
	for {
		fmt.Print("Введите ключ: ")
		fmt.Scan(&deleteKey)
		for key := range bookmarks {
			if key == deleteKey {
				return deleteKey
				//break
			}
		}
		fmt.Println("Повторите ввод, введен не существующий ключ!")
		continue
	}
}
