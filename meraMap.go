package main

import "fmt"

func main() {
	//assigned values to meraMap 
	
	meraMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	//print statement of meraMap
	fmt.Println("mera map hai:", meraMap)

	//called the function khudkaMap in mainfunc,in the variable newMap and asigned arguments to it then printed the value of khudkaMap
	newMap := khudkaMap(meraMap)
	fmt.Println("khudka map hai:", newMap)
	deleteMap := delMap(meraMap)
	fmt.Println("deleted wala map hai:", deleteMap)
}

//created map using make function and int
//asigend value to 1 which is khudka example banao
func khudkaMap(meraMap map[int]string) map[int]string {
	meraMap[1] = "khudka exapmpla banao"
	fmt.Println(meraMap)
	return meraMap
}

//deleted the 2 element from mera Map
func delMap(meraMap map[int]string) map[int]string {
	delete(meraMap, 2)
	fmt.Println(meraMap)
	return meraMap
}
