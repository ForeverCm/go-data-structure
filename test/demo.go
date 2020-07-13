package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//realSingle 读入单个变量
func realSingle() {
	var a int
	var b string
	fmt.Scan(&a, &b)

}
func getCubeRoot(m float64) float64{
	min,max:=0.0,m
	for max-min>1e-7{
		mid:=(max+min)/2
		if mid*mid*mid>m{
			max=mid
		}else if mid*mid*mid<m{
			min=mid
		}else{
			return mid
		}
	}
	return min
}

func demo(str string) int{
	if len(str)==0{
		return 0
	}
	var res int
	for _,v:=range str{
		if v>='A'&&v<='Z'{
			res++
		}
	}
	return res
}



func Getbig(a, b int) int {
	if a < b{
		a, b = b, a
	}
	if a % b == 0{
		return b
	}
	return Getbig(b, a % b)
}

//readMutiLine 读入多行
func readMutilLine() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	temp := scanner.Text()
	//}
}

func Demo2() {
	var n int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp := scanner.Text()
		if len(temp) <= 0 {
			return
		}
		if n <= 0 {
			n, err = strconv.Atoi(temp)
			if err != nil {
				return
			}
		} else if n > 0 {
			n--
			for len(temp) >= 8 {
				fmt.Println(temp[:8])
				temp = temp[8:]
			}
			if len(temp) > 0 {
				for len(temp) < 8 {
					temp += "0"
				}
				fmt.Println(temp)
			}
		}
	}
}

//readOneLine  控制台读入一行，回车换行
func readOneLine() {
	//reader := bufio.NewReader(os.Stdin)
	//a, _, _ := reader.ReadLine()
	//b:=demo([]byte(a))

}
func Demo(src []byte)string{
	i:=0
	j:=len(src)-1
	for i<=j{
		src[i],src[j]=src[j],src[i]
		i++
		j--
	}
	//fmt.Printf("----1%s\n", src)
	return string(src)
}


//subSequenceIndex 没有找到则返回-1
func subSequenceIndex(target, source string) int {
	if len(target) == 0 || len(source) <= 0 || len(target) > len(source) {
		return -1
	}
	i := 0
	j := 0
	for i < len(source) && j < len(target) {
		if source[i] == target[j] {
			i++
		}
		j++
		if j == len(target) {
			return i+1
		}
	}
	return -1
}


//subSequenceIndex 没有找到则返回-1
func subSequenceIndex2(target, source string) int {
	if len(target) == 0 || len(source) <= 0 || len(target) > len(source) {
		return -1
	}
	i := 0
	j := 0
	lastIndex := -1
	for  i < len(target) && j < len(source) {
		if target[i] == source[j] {
			if i == 0 {
				lastIndex = j
			}
			i++
			if i == len(target) {
				i = 0
			}
		}
		j++
	}

	return lastIndex

}




