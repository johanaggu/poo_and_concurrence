package main

func simpleArraySum(ar []int32) int32 {
    if(len(ar) == 1){
        return ar[0]
    }
   
    newS := ar[:len(ar)-1] //Extract new slice without ending item 
    
    aux := ar[len(ar)-1] //get ending item
    newS[0] += aux //sum value

    return simpleArraySum(newS)
}


func main() {
	
}