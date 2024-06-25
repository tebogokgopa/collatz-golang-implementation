package main

import(
	"fmt"
	"container/list"
)


func collatzIterative(start_num int) *list.List {

	l := list.New()
	l.PushBack(start_num)

	for start_num > 1{

		if start_num == 1 {
			l.PushBack(start_num)
			break
		}

		if start_num % 2 == 0 {
			start_num = start_num / 2
			l.PushBack(start_num)
		}else{
			start_num  = start_num * 3 + 1
			l.PushBack(start_num)
		}

	}

	return  l
}

func collatzRecursive(start_num int) (int) {

	if start_num == 1 {
		fmt.Print(start_num)
		return start_num
	}

	if start_num % 2 == 0 {
		fmt.Printf("%v ",start_num)
		start_num = collatzRecursive(start_num / 2)
		
	}else{
		fmt.Printf("%v ",start_num)
		start_num  =collatzRecursive( start_num * 3 + 1)
	}

	return start_num
}

func main()  {
	fmt.Println("\nCollatz Sequence iterative")
	
	p := collatzIterative(13)
	for i := p.Front(); i != nil; i = i.Next(){
		fmt.Printf("%v ",i.Value)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Collatz Sequence recursive")
	s := collatzRecursive(13)

	fmt.Println("\n\n ")
	
	fmt.Printf("Sequence stops at %v", s)
	fmt.Println("\n ")
}

