
func retMin(a int, b int) int {
    
    if( a < b ){
        return a
    }
    return b
        
}

func minPathSum(grid [][]int) int {
    
    y := len(grid[0])
    x := len(grid)
    
    // fmt.Println(x,y)
    ans := make([][]int, x)
    
    for i:=0 ; i<x ; i++{
        ans[i] = make([]int , y)
    }
    
    // fmt.Println(len(ans),len(ans[0]))
    
    ans[x-1][y-1] = grid[x-1][y-1]
    
    for j:=y-2 ; j>=0 ; j--{
        ans[x-1][j] = grid[x-1][j] + ans[x-1][j+1]
    }
    
    for i:=x-2 ; i>=0 ; i--{
        ans[i][y-1] = grid[i][y-1] + ans[i+1][y-1]
    }
    
    // fmt.Println(ans)
    for i := x-2;i>=0 ; i--{
        for j:=y-2 ; j>=0 ; j--{
            ans[i][j] = retMin(ans[i+1][j], ans[i][j+1]) + grid[i][j]
        }
    }
    // fmt.Println(ans)
    return ans[0][0]
}
