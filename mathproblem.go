package main 

import
(
    "fmt"
)

func isDivisible(num1, num2 uint32) bool {
    divisor := float64(float64(num1) / float64(num2))
    remainder := float64(divisor - float64(uint32(divisor))) 
    
    if (remainder == 0) {
        // fmt.Printf("Num1: %v N: %v divisor: %v\n", num1, num2, divisor)
        return true
    } 
    
    return false
}

func findNumber(num uint32, n uint16, numbers []uint16) (bool, uint32) {
    // fmt.Printf("N: %v Num: %v\n", n, num)
    if (10 == n) {
        return true, num
    }
    
    temp := num
    for i := 0; i < len(numbers); i++ {
        num = num + uint32(numbers[i])
        if (isDivisible(num, uint32(n))) {
            
            
            newArray := make([]uint16, len(numbers) - 1)
            index := 0
            for j := 0; j < len(numbers); j++ {
                if (i == j) {
                    continue
                }
                
                newArray[index] = numbers[j]
                index++
            }
            
            found, result := findNumber(num * 10, n + 1, newArray)
            
            if found {
                return found, result
            }
        }
        
        num = temp
    }
    
    return false, num
}

func main() {
    _, result := findNumber(0, 1, []uint16{1, 2, 3, 4, 5, 6, 7 , 8, 9})
	fmt.Println(result)
}