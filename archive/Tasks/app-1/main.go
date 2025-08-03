package main

import "fmt"

type Soldaten struct {
	Name       string
	Alter      int
	Dienstgrad string
	Abteilung  string
}

func (s Soldaten) Informationen() (str string) {
	str = fmt.Sprintf(" Name:\t\t%s,\n Alter:\t\t%d,\n Dienstgrad:\t%s,\n Abteilung:\t%s\n\n", s.Name, s.Alter, s.Dienstgrad, s.Abteilung)
	return
}

func PrintInfo(SSlc []Soldaten) {
	fmt.Printf("\n|========================|==|========================|")
	fmt.Printf("\nInformationen über die Personalien der zur Verleihung \n      der Ehrenmedaille vorangegangenen Personen:\n")
	fmt.Printf("|====================================================|\n")
	for index, soldat := range SSlc {
			fmt.Printf("\t|============|Canididate No.[%d]|============|\n", index+1)
			fmt.Print(soldat.Informationen())
		}
	}

func AlterInfo(SSlc []Soldaten) {
	fmt.Print("\n\n|================|Altersabstufung|================|\n")
	var minID, maxID, sum int
	for i, s := range SSlc {
	switch true{
	case s.Alter < SSlc[minID].Alter:
		minID = i
	case s.Alter > SSlc[maxID].Alter:
		maxID = i
	}
	sum += s.Alter
	}
	fmt.Printf(" Min Alter:%d\n MaxAlter:%d\n MittelwertAlter:%3.f", minID, maxID, float64(sum) / float64(len(SSlc)))
	fmt.Print("\n|=================================================|\n")
}

func AbteilungSum(SSlc []Soldaten) {
	AbteilungMap := make(map[string]int)
	for _, value := range SSlc {
		AbteilungMap[value.Abteilung]++  // ← вот здесь работает магия
	}

	fmt.Print("\n\n|=================|Abteilung Zählung|=================|\n")
	for abt, count := range AbteilungMap {
		fmt.Printf(" %s: %d Soldat(en)\n", abt, count)
	}
	fmt.Print("|=====================================================|\n")
}

func main() {
	SSlc := []Soldaten{
		//Heer
		{"Karl Heinemann", 38, "Leutnant", "Grenadier-Regiment 578"},
		{"Walter Zimmermann", 27, "Hauptmann", "Grenadier-Regiment 578"},
		{"Ernst Richter", 24, "Unteroffizier", "Grenadier-Regiment 578"},
		//Luftwaffe
		{"Klaus Neubauer", 31, "Oberleutnant", "Jagdgeschwader 54 \"Grünherz\""},
		{"Hans Keller", 22, "Feldwebel", "Jagdgeschwader 54 \"Grünherz\""},
		{"Karl Heinemann", 28, "Leutnant", "Jagdgeschwader 54 \"Grünherz\""},
		//Waffen-SS
		{"Rudolf Drexler", 35, "Sturmbannführer", "SS-Division \"Totenkopf\""},
		{"Karl Heinemann", 28, "Leutnant", "SS-Division \"Totenkopf\""},
		//Kriegsmarine
		{"Heinrich Vogt", 50, "Kapitänleutnant", "U-Boot-Flottille 7"},
		{"Karl-Heinz Möller", 19, "Matrose", "U-Boot-Flottille 7"},
	}

	PrintInfo(SSlc)
	AlterInfo(SSlc)
	AbteilungSum(SSlc)
}
