package main

import (
	"fmt"
	//"syscall/js"
)

var userRequest string

/*
func receiveInput(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		println("Нет аргументов")
		return nil
	}
	userRequest = args[0].String()
	fmt.Println("Получено из JS:", userRequest)
	return nil
}
*/

func main() {
	//js.Global().Set("receiveInput", js.FuncOf(receiveInput))

	strUserRequest := `POST /api/products HTTP/1.1
Host: example.com
Content-Type: application/json
Authorization: Bearer abc123
Content-Length: 31

{
  "name": "Box",
  "price": 100
}`

	slcUserRequest := []rune(strUserRequest)

	var start []rune
	var headers []rune
	var body []rune

	var method, path, version string
	mapHeaders := make(map[string]string)
	mapBody := make(map[string]string)

	// === ЭТАП 1: Разделение на части ===
	stage := "start"
	for i := 0; i < len(slcUserRequest); i++ {
		ch := slcUserRequest[i]

		if stage == "start" {
			if ch == '\n' {
				stage = "headers"
				continue
			}
			start = append(start, ch)
			continue
		}

		if stage == "headers" {
			if ch == '\n' && i > 0 && slcUserRequest[i-1] == '\n' {
				stage = "body"
				continue
			}
			headers = append(headers, ch)
			continue
		}

		if stage == "body" {
			body = append(body, ch)
		}
	}

	// === ЭТАП 2: Разбор стартовой строки ===
	part := 0
	var buf []rune
	for i := 0; i < len(start); i++ {
		ch := start[i]

		if ch == ' ' {
			switch part {
			case 0:
				method = string(buf)
			case 1:
				path = string(buf)
			}
			part++
			buf = nil
			continue
		}
		buf = append(buf, ch)
	}
	version = string(buf)

	// === ЭТАП 3: Разбор заголовков ===
	var currentLine []rune
	for i := 0; i < len(headers); i++ {
		ch := headers[i]
		if ch != '\n' {
			currentLine = append(currentLine, ch)
			continue
		}

		var keyRunes, valRunes []rune
		buildingKey := true
		for j := 0; j < len(currentLine); j++ {
			if buildingKey {
				if currentLine[j] == ':' {
					buildingKey = false
					continue
				}
				keyRunes = append(keyRunes, currentLine[j])
			} else {
				if len(valRunes) == 0 && currentLine[j] == ' ' {
					continue
				}
				valRunes = append(valRunes, currentLine[j])
			}
		}
		mapHeaders[string(keyRunes)] = string(valRunes)
		currentLine = nil
	}

	// === ВЫВОД ===
	fmt.Println("START LINE:")
	fmt.Printf("Method: %s\nPath: %s\nVersion: %s\n", method, path, version)

	fmt.Println("\nHEADERS:")
	for k, v := range mapHeaders {
		fmt.Printf("%s → %s\n", k, v)
	}

	fmt.Println("\nBODY:")
	bodyStr := string(body)
	fmt.Println(bodyStr)

	// === ЭТАП 4: Разбор тела (JSON без вложенности) ===
	var keyRunes []rune
	var valRunes []rune
	inKey := false
	inValue := false
	buildingKey := true
	reading := false
	mapBody = make(map[string]string)

	for i := 0; i < len(body); i++ {
		ch := body[i]

		// Игнорируем пробелы и переводы строк вне значений
		if !reading {
			if ch == '"' {
				reading = true
				if buildingKey {
					inKey = true
				} else {
					inValue = true
				}
				continue
			}
			continue
		}

		if reading {
			// Закрытие кавычек
			if ch == '"' {
				reading = false
				if inKey {
					inKey = false
					buildingKey = false
				} else if inValue {
					inValue = false
					buildingKey = true
					// Запись пары
					mapBody[string(keyRunes)] = string(valRunes)
					keyRunes = nil
					valRunes = nil
				}
				continue
			}

			// Заполнение ключа или значения
			if inKey {
				keyRunes = append(keyRunes, ch)
			}
			if inValue {
				valRunes = append(valRunes, ch)
			}
		}



		
	}

	//select {}
}
/*
ПРИМЕР ЗАПРОСА
POST /api/products HTTP/1.1\n
\n
Host: example.com
Content-Type: application/json
Authorization: Bearer abc123
Content-Length: 31

{
  "name": "Box",
  "price": 100
}
	
ПРИНАДЛЕЖНОСТЬ К ГРУППЕ ЗАГЛАВНЫХ БУКВ
for i := 65; i <= 90; i++ {
    fmt.Printf("%c ", i)
}

КОДИРОВКИ ВАЖНЫХ СИМВОЛОВ
КОД ПРОБЕЛА 32
КОД ДВОЕТОЧИЯ 58
КОД ФИГУРНЫХ СКОБОК { - 123 } - 125
*/