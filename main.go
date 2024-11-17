package main

import (
	"fmt"
)

type bokmarkMap = map[string]string

func main() {

	bookmarks := bokmarkMap{}
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
			appendBookmarks(bookmarks)
		case 3:
			deleteBookmarks(bookmarks)
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

func printBookmarks(bookmarks bokmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	} else {
		fmt.Println("Список закладок: ")
		for key, value := range bookmarks {
			fmt.Println(key, value)
		}
	}
}

func appendBookmarks(bookmarks bokmarkMap) {
	var key string
	var value string
	fmt.Println("Введите ключ: ")
	fmt.Scan(&key)
	fmt.Println("Введите значение: ")
	fmt.Scan(&value)

	bookmarks[key] = value
	//return bookmarks
}

func deleteBookmarks(bookmarks bokmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
		//return bookmarks
	} else {
		var deleteKey string
		fmt.Println("Какой элемент хотите удалть?")
		for key, value := range bookmarks {
			fmt.Println(key, value)
		}
	MainFor:
		for {
			fmt.Print("Введите ключ: ")
			fmt.Scan(&deleteKey)
			for key := range bookmarks {
				if key == deleteKey {
					delete(bookmarks, deleteKey)
					fmt.Println("Элемент успешно удалён.")
					break MainFor //return bookmarks
				}
			}
			fmt.Println("Повторите ввод, введен не существующий ключ!")
		}
	}
}
