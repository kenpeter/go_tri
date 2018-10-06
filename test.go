package main
import "fmt"

func main() { 
 	arr := []int{1, 2, 3} 
	change(&arr)
}

func change(arr *[]int) {
	for i:=0; i<len(*arr); i++ {
		item := (*arr)[i]	
		fmt.Println(item)
	}
			
}
