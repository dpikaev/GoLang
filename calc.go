package main
import "fmt"
func main(){
    var a int
    fmt.Println("Введите число 1")
    fmt.Scan(&a)
    var b int
    fmt.Println("Введите число 2")
    fmt.Scan(&b)
    var c string
    fmt.Println("Введите знак операции (+,-,/,*)")
    fmt.Scan(&c)
    switch{
    case c=="+":
        fmt.Println("Сумма чисел равна",a+b)
    case c=="-":
        fmt.Println("Разность чисел равна",a-b)
    case c=="/":
        fmt.Println("Частное чисел равно",a/b)
    case c=="*":
        fmt.Println("Произведение чисел равно",a*b)
    }
}