package algorithm

import (
	"fmt"
)

const N = 8

var result [][N][N]int

type Point struct {
	x int
	y int
}

//生成正确的括号 n个左括号n个右括号
func GenerateRightBrackets(n int, left int, right int, str string) {
	if left == n && right == n {
		fmt.Println(str)
		return
	}
	if left < n {
		GenerateRightBrackets(n, left+1, right, str+"(")
	}
	if left > right && right < n {
		GenerateRightBrackets(n, left, right+1, str+")")
	}
}



//n皇后问题解法1
func SolveNQueens(n int){
    var tmpResult [N][N]int
    //PrintArray(N, result)
    //表示是是否右皇后
	var cur [N][2 * N]int
	
	DFS(N, 0, cur,tmpResult)
    Print3Array(N, result)
}

func Print3Array(n int, arr [][N][N]int) {
    for i := 0; i < len(arr); i ++ {
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                fmt.Print(arr[i][j][k], " ")
            }
            fmt.Println()
        }
        fmt.Println()
    }
}
func Print2Array(n int, arr [N][2 * N]int) {
    for i := 0; i < n; i ++ {
        for j := 0; j < n * 2; j++ {
            fmt.Print(arr[i][j], " ")
        }
        fmt.Println()
    }
}
func DFS(n int, row int, cur [N][2 * N]int, tmpResult [N][N]int){
    if row >= n {
		result = append(result, tmpResult)
		return 
    }

    for col := 0 ; col < n ; col ++ {
        //没有皇后攻击
        if cur[0][row] == 0 && cur[1][col] == 0 && cur[2][row + col] == 0 && cur[3][n + row - col] == 0 {
            // fmt.Printf("(%v,%v) \n", row, col)
            // fmt.Println("######")
            // Print2Array(N, cur)
            // fmt.Println("######")
            cur[0][row] = 1
            cur[1][col] = 1
            cur[2][row + col] = 1
            cur[3][n + row - col] = 1
            tmpResult[row][col] = 1

            // fmt.Println("*******")
            // Print2Array(N, cur)
            // fmt.Println("*******")
			DFS(n, row + 1, cur,tmpResult)
			cur[0][row] = 0
            cur[1][col] = 0
            cur[2][row + col] = 0
            cur[3][n + row - col] = 0
            tmpResult[row][col] = 0
        }
    }
    //return num, result
}
//乘积最大子序列
func F1() {
    arr := []int {2, 3, -10, 5, -1}
    //最终结果集
    arr1 := make([]int, 0)

    //中间结果集
    arr2 := make([]int, 1)
    arr2[0] = arr[0]
    arr1, arr2 = f11(arr, 1, arr1, arr2)
    fmt.Println(maxArr(arr1))
}


//暴力求解
func f11(arr []int, n int, arr1 []int, arr2 []int) ([]int, []int) {
    if n == len(arr) {
        for i := 0; i < len(arr2); i ++ {
            arr1 = append(arr1, arr2[i])
        }
        return arr1, arr2
    }
    for i := 0; i < len(arr2); i ++ {
        arr1 = append(arr1, arr2[i])
        arr2[i] *= arr[n]
    }
    arr2 = append(arr2, arr[n])
    return f11(arr, n + 1, arr1, arr2)
}

func maxArr(arr []int) int {
    max := arr[0]
    for i := 0; i < len(arr); i ++ {
        if max < arr[i] {
            max = arr[i]
        }
    }
    return max
}

//取数组最小值
func minArr(arr []int) int {
    min := arr[0]
    for i := 1; i < len(arr); i ++ {
        if min > arr[i] {
            min = arr[i]
        }
    }
    return min
}

//零钱兑换
func F2(arr []int, n int) int{
    min := minArr(arr)
    max := maxArr(arr)
    if min <= 0 {
        return -1
    }
    length := len(arr)
    maxCount := n / min + 1
    dp := make([]int, maxCount)
    //初始化
    for i := 0; i < length; i ++ {
        dp[arr[i]] = 1 
    }

    for i := max + 1; i < maxCount; i ++ {
        arr1 := make([]int, 0)
        for j := 0; j < length; j ++ {
            count := dp[i - arr[j]]
            if count != 0 {
                arr1 = append(arr1, count)
            }   
        }
        minCount := 0
        if len(arr1) != 0 {
            minCount = minArr(arr1) + 1
        }
        dp[i] = minCount
    }
    fmt.Println(dp)
    return dp[n]
}
