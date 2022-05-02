package leetcode

func rotate(matrix [][]int) {

	if len(matrix) == 1 {
		return
	}

	//先沿左上到右下对角线翻转
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//再左右翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
	return
}

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func findDiagonalOrder(mat [][]int) []int {
	rowLen := len(mat)    //多少行
	colLen := len(mat[0]) //多少列
	res := make([]int, 0, rowLen*colLen)
	k, l := 0, 0 //每次的起点
	for i := 0; i < rowLen+colLen-1; i++ {
		if i%2 == 0 {
			for j := k; j >= i-l; j-- {
				res = append(res, mat[j][i-j])
			}
		} else {
			for j := l; j >= i-k; j-- {
				res = append(res, mat[i-j][j])
			}
		}

		if k >= rowLen-1 {
			k = rowLen - 1
		} else {
			k++
		}

		if l >= colLen-1 {
			l = colLen - 1
		} else {
			l++
		}
	}
	return res
}
