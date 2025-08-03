package main //	Основной пакет.

import "fmt"

type FlightNumber string //Пользовательские типы
type Dollar string       //для читабельности.
type Km int
type People int

//Интерфейс, который ассоциирует себя с той функцией,
//которая имеет указанный в ней метод.(СДЕЛАЛ С ПОМОЩЬЮ)
type CostPrinter interface {
	CostPrint() string
}

type Plane struct { // Создали тип Plane с параметриами.
	ID              FlightNumber
	Cost            Dollar
	CountryOfOrigin string
	EngineType      string
	FullModelName   string
	Color           string
	FlightRange     Km
	Capacity        People
}

//====МЕТОД СТОИМОСТИ ЭКСПЛУАТАЦИИ=====//
func (p Plane) CostPrint() string {
	msg := fmt.Sprintf("МОДЕЛЬ - %s;\nБОРТОВОЙ НОМЕР - %s;\n ==================== \n СТОИМОСТЬ ЭКСПЛУАТАЦИИ - %s/час.", p.FullModelName, p.ID, p.Cost)
	return msg
}

//Главная функция - горутина в которой выполняется весь код.
func main() {
	//Создали экземпляры структуры.
	var Plane_1 = Plane{
		ID:              "PL-100",
		Cost:            "$200",
		CountryOfOrigin: "Germany",
		EngineType:      "Jumo 210D",
		FullModelName:   "Messerschmitt Bf-109",
		Color:           "black",
		FlightRange:     500,
		Capacity:        50,
	}
	var Plane_2 = Plane{
		ID:              "PL-101",
		Cost:            "$200",
		CountryOfOrigin: "Germany",
		EngineType:      "Jumo 210D",
		FullModelName:   "Messerschmitt Bf-109",
		Color:           "black",
		FlightRange:     500,
		Capacity:        50,
	}
	var Plane_3 = Plane{
		ID:              "PL-102",
		Cost:            "$200",
		CountryOfOrigin: "Germany",
		EngineType:      "Jumo 210D",
		FullModelName:   "Messerschmitt Bf-109",
		Color:           "black",
		FlightRange:     500,
		Capacity:        50,
	}
	
	var arrPlane [3]Plane = [3]Plane{Plane_1, Plane_2, Plane_3}       //Перенесли созданные экземпляры в массив.
	var slcPlane []Plane = arrPlane[:]                                //Создаем слайс на основе массива выше.
	var mpPlane map[FlightNumber]Plane = make(map[FlightNumber]Plane) //Создаем карту для создания коллекции экземпляров.

	var request FlightNumber = "PL-102" //Заранее прописаный запрос.

	var interfaceSls []CostPrinter
	for _, p := range slcPlane { //Заполнение слайса интерфейсов (СДЕЛАЛ С ПОМОЩЬЮ)
		interfaceSls = append(interfaceSls, p)
	}

	var slcPlaneReserv []Plane
	copy(slcPlaneReserv,slcPlane)

	MapInit(slcPlane, mpPlane)//Реализуем инициализацию карты через функцию
	CostInfoPrint(Plane_2)//ИНФОРМАЦИЯ О СТОИМОСТИ ЭКСПЛУАТАЦИИ
	CostInfoPrintAll(interfaceSls)//ИНФОРМАЦИЯ О СТОИМОСТИ ЭКСПЛУАТАЦИИ ПО СРЕЗУ
	MapSearch(request, mpPlane)//ПОИСК ПО КАРТЕ
}

//==========ИНИЦАЛИЗИРУЕМ КАРТУ=========//
func MapInit(slcPlane []Plane, mpPlane map[FlightNumber]Plane) {
	var InitVar Plane

	for index := range slcPlane {
		InitVar = slcPlane[index]
		mpPlane[InitVar.ID] = slcPlane[index]
	}

	fmt.Println("\nКАРТА УСПЕШНО ПРОИНИЦИАЛИЗИРОВАННА!")
}
//=============ПЕЧАТЬ КАРТЫ=============//
func TestPrint(slcPlane []Plane, mpPlane map[FlightNumber]Plane) {
	for key := range mpPlane {
		fmt.Println(mpPlane[key])
	}
}
//============ИНФОРМАЦИЯ О СТОИМОСТИ ЭКСПЛУАТАЦИИ===========//(СДЕЛАЛ С ПОМОЩЬЮ)
func CostInfoPrint(c CostPrinter) {
	fmt.Println(c.CostPrint())
}
//============ИНФОРМАЦИЯ О СТОИМОСТИ ЭКСПЛУАТАЦИИ ПО СРЕЗУ===========//(СДЕЛАЛ С ПОМОЩЬЮ)
func CostInfoPrintAll(c []CostPrinter) {
	for index := range c {
		fmt.Println(c[index].CostPrint())
	}
}
//============ПОИСК ПО КАРТЕ===========//
func MapSearch(request FlightNumber,mpPlane map[FlightNumber]Plane){
	str, exs := mpPlane[request]
	if exs {
		fmt.Println("Воздушное судно найдено:\n",str)
	} else {
		fmt.Println("Воздушное судно отсутствует:\n",str)
	}
}

