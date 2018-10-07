package main
import "fmt"

func main() {
	gen(4)
}

func gen(n int) {
	// append arr to arr
	all := [][]int{}
	row := []int{}
	for i:=0; i<n; i++ {
		// to start, row reuse
		// append element 1 by 1
		row = append([]int{1}, row...)
		// skip 0, and end
		for j:=1; j<len(row)-1; j++	 {
			// in place
			row[j] = row[j] + row[j+1]			
		} // end

		// val not ref 
		all = append(all, row)
	} // end

	fmt.Println(all)
}
