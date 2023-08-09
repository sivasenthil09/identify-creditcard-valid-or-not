package main

import "fmt"

func main() {
    var num ,temp,count ,count2,sum1 ,sum2 int64
    count=0
    count2=1
    sum1 =0
    sum2=0
	fmt.Println("enter the credit card number")
    fmt.Scan(&num)
    temp = num
    for temp!=0{
        temp=temp/10
        count++
    }
    //fmt.Print(count)
    if count <16{
        fmt.Print("not valid")

    }else{
        temp=num
        //fmt.Print(temp)
        for temp!=0{
            
        if (count2 % 2)!=0{
            sum1=sum1+(temp%10)
            //fmt.Print(sum1," ")
            count2++
            temp=temp/10

        }else{
            mul:=(temp%10)*2
            if mul>9{
                mul = mul-9
            }
            sum2=sum2+mul
            //fmt.Print(sum2," ")
            count2++
            temp=temp/10
        }
        

    }
    ans:=sum1+sum2
if ans%10==0{
    fmt.Print("valid")

}else{
    fmt.Print("invalid")
}
}

}