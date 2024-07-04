package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	name string
	lastname string
	age	int
	id string
}

func (p Person) PrintPerson()  {
	fmt.Println("Normal Print")
	fmt.Printf("Name : %s\n",p.name)
	fmt.Printf("Surname : %s\n",p.lastname)
	fmt.Printf("Age : %d\n", p.age)
	fmt.Printf("Identity No : %s\n",p.id)
}

func (p Person) SavePersonToFile(){
	fmt.Println("Save to file")
	f,err := os.Create("personfile.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}

	defer f.Close()
	
	fmt.Fprintf(f,"Name : %s\n", p.name)
	fmt.Fprintf(f, "Surname : %s\n",p.lastname)
	fmt.Fprintf(f,"Age : %d\n", p.age)
	fmt.Fprintf(f,"Identity No: %s\n",p.id)
	fmt.Fprintln(f,"")
}

func (p Person) ReadFromFile()  {
	fmt.Println("Read from file")
	f, err := os.Open("personfile.txt")
	if err != nil {
		fmt.Print(err)
	}


	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}

		fmt.Print(line)
	}
	
}


func main()  {
	
	me := Person{
		"Lesego",
		"Mmatli",
		24,
		"8402224785886",
	}

	me.PrintPerson()
	me.SavePersonToFile()
	me.ReadFromFile()
}