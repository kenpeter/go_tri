package main
import "fmt"
func main() {
	gen(5)	
}

func gen(numRows int) {
	arrRes := [][]int{}
	arrRes = build(numRows, 0, arrRes)
	return arrRes
}

func build(n int, level int, arrRes [][]int) [][]int {
	if(n == level) {
		return arrRes 
	}

	arr := []int{}
	if level == 0 {
		arr = append(arr, 1)
	} else if level == 1 {
        arr = append(arr, 1, 1)
	} else {
		// get it out
		tmp := arrRes[level-1]
		arr = comb(tmp)
	}

	arrRes = append(arrRes, arr)
	arrRes = build(n, level+1, arrRes)
	return arrRes;
}

func comb(arr []int) []int{
	// arr type init
	tmpArr := []int{1}
	for i:=1;  i<len(arr); i++ {
		sum := arr[i-1] + arr[i]
		tmpArr = append(tmpArr, sum)	
	}

	// go use val, not ref
	tmpArr = append(tmpArr, 1)
	return tmpArr;
}
