func searchMatrix(matrix [][]int, target int) bool {

    row := len(matrix)
    col := len(matrix[0])

    for i := 0; i < row; i++ {

        if target >= matrix[i][0] &&  target <= matrix[i][col-1] {

            s := 0
            e := col-1

            for s <= e {

                mid := s + (e-s)/2

                if matrix[i][mid] == target {

                    return true
                }else if matrix[i][mid] < target {

                    s =  mid +1


                }else{

                    e = mid-1
                }
            }

            return false

        } 

    }

    return false

}
