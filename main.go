package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}

	var sparseArr []ValNode

	valNode := ValNode{ //记录原始数组大小和默认值
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v1 := range chessMap { //遍历原始数组
		for j, v2 := range v1 {
			if v2 != 0 { //当值不为0时
				valNode = ValNode{ //创建节点结构体，记录位置和值
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Println("转换后的稀疏数组如下")
	for _, v := range sparseArr {
		fmt.Printf("%d\t%d\t%d\n", v.row, v.col, v.val)
	}

	var chessMap2 [11][11]int

	for i, v := range sparseArr {
		if i != 0 { //跳过第一行记录值
			chessMap2[v.row][v.col] = v.val
		}
	}

	fmt.Println("恢复后的原数组如下")
	for _, v1 := range chessMap2 {
		for _, v2 := range v1 {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}
}
