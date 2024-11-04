package data

import (
	"encoding/json"
	"fmt"
	"os"
)

func storeSubjectsToJsonFile() {

	var Subjects = make([][]string, 0)

	Subjects = append(Subjects, []string{"Calculas0", "Esp0"})

	Subjects = append(Subjects, []string{"calculasi",
		"information system",
		"introduction to computing",
		"espi",
		"fundmentals of business",
		"physics"})
	Subjects = append(Subjects, []string{"calculasii",
		"espii",
		"problem solving",
		"advanced physics",
		"discrete mathimatics",
		"enterperner ship",
		"communication skills"})
	Subjects = append(Subjects, []string{"linear algebra",
		"data base",
		"digital",
		"oop",
		"probapility",
		"networks"})
	Subjects = append(Subjects, []string{"advanced application programming",
		"computer architecture",
		"cyper security",
		"data structure and algorithms",
		"software engineering",
		"Web"})
	Subjects = append(Subjects, []string{"differential equations",
		"intro to AI",
		"operating systems",
		"system programming",
		"theory of computations",
		"Flutter"})
	Subjects = append(Subjects, []string{"advanced statistics",
		"computer graphics",
		"numerical method",
		"system modeling and simulation",
		"Professional training I", "Computing algorithms"})

	Subjects = append(Subjects, []string{"structure of programming language",
		"computer and society",
		"project1",
		"digital image processing",
		"embedded system programming", "professional training II"})

	Subjects = append(Subjects, []string{"computer System srcurity",
		"Human computer interaction",
		"project2",
		"robotics application",
		"deep learning", "professional trainingIII"})

	json, err := json.MarshalIndent(Subjects, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("json/subjects.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(json)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Subjects written to subjects.json succesfully !")
}

func SubjectsFromJsonToSlice() ([][]string, error) {
	data, err := os.ReadFile("json/subjects.json")
	// file , err := os.Open("subjects.json")
	if err != nil {
		return nil, err
	}
	//  defer file.Close()
	var subjects [][]string
	//  if err := json.NewDecoder(file).Decode(&subjects) ; err != nil
	if err := json.Unmarshal(data, &subjects); err != nil {
		return nil, err
	}
	return subjects, nil

}
