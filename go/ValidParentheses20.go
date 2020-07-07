func isValid(s string) bool {
    if len(s) == 0{
        return true
    }
    mapping := map[rune]rune{
        '(' : ')',
        '[' : ']',
        '{' : '}',
    }
    otherWay := map[rune]rune{
        ')' : '(',
        ']' : '[',
        '}' : '{',
    }
    stack := []rune{}
    for _, char := range(s){
        if _, eval := mapping[char]; eval{
            stack = append(stack,char)
        }
        if val, eval := otherWay[char]; eval{
            if(len(stack) > 0 && stack[len(stack)-1] == val){
                stack = stack[:len(stack)-1]
            }else{
                return false
            }
        }
    }
    return len(stack) == 0
}
