package main

/*func main() {
    var c1, c2, c3 chan int
    var i1, i2 int

    select {
        case i1 = <-c1:
            fmt.Printf("recived ", i1, "from c1\n")
        case c2 <- i2:
            fmt.Printf("sent ", i2, " to c2\n ")
        case i3, ok := (<-c3):
            if ok {
                fmt.Printf("recived ", i3, " from c3\n")
            } else {
                fmt.Printf("c3 is closed\n")
        }
        default:
            fmt.Printf("no communication\n")
   }

}*/
/*func main(){
    var a int = 100
    var b int = 200
    var ret int

    ret = max( a, b )
    fmt.Printf("最大值是：%d\n", ret)
}

func max( num1, num2 int ) int {
    var result int

    if ( num1 > num2 ){
        result = num1
    } else {
        result = num2
    }
    return result
}*/

/*func swap(x int, y string) (int, string){
    return x, y
}

func main(){
    a, b := swap(1, "xiangk")
    fmt.Println( a, b )
}*/

/*func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main(){
    nextNumber := getSequence()

    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())

    nextNumber1 := getSequence()
    fmt.Println( nextNumber1() )
    fmt.Println( nextNumber1() )
}*/

/*func main(){
    add_func := add(1, 2)
    fmt.Println( add_func() )
    fmt.Println( add_func() )
    fmt.Println( add_func() )
}

func add( x1, x2 int ) func() ( int, int ){
    i := 0
    return func() ( int, int ){
        i++
        return i, x1 + x2
    }
}*/

/*func main(){
    var n [10]int
    var i,j int

    for i = 0; i < 10; i++{
        n[i] = i + 10
    }

    for j = 0; j < 10; j++{
        fmt.Printf("Element[%d] = %d\n", j, n[j])
    }
}*/

/*func main(){
    var a int = 20
    var ip *int
    ip = &a

    fmt.Printf("a 變量地址是：%x\n", &a)
    fmt.Printf("ip 變量存儲的指針地址：%x\n", ip)
    fmt.Printf("*ip 變量的值：%d\n", *ip)
}*/

/*const MAX int = 3

func main(){
    a := [] int {10, 100, 200}
    var i int
    for i = 0; i < MAX; i++ {
        fmt.Printf("a[%d] = %d\n", i, a[i])
    }
}*/

/*const MAX int = 3

func main() {
    a := []int{10, 100, 200}
    var i int
    var ptr [MAX]*int

    for i = 0; i < MAX; i++ {
        ptr[i] = &a[i]
    }

    for i = 0;i <MAX; i++ {
        fmt.Printf("a[%d] = %d\n", i, *ptr[i])
    }
}*/

/*func main(){
    var a int
    var ptr *int
    var pptr **int

    a = 300

    ptr = &a
    pptr = &ptr

    fmt.Printf("var a = %d\n", a )
    fmt.Printf("point var a = %d\n", *ptr )
    fmt.Printf("point point var a = %d\n", **pptr )
}*/

/*type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main(){
    fmt.Println( Books{"Go 語言", "www.runoob.com", "Go 語言教程", 6495407} )
    fmt.Println(Books{title:"Go 語言", author: "www.runoob.com"})
}*/

/*type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main(){
    var Book1 Books
    var Book2 Books

    Book1.title = "Go 語言"
    Book1.author = "ww.runoob.com"
    Book1.subject = "Go 語言教程"
    Book1.book_id = 6495407

    Book2.title = "Go 語言"
    Book2.author = "ww.runoob.com"
    Book2.subject = "Go 語言教程"
    Book2.book_id = 6495470

    printBook(1, &Book1 )

    printBook( 2, &Book2 )
}

func printBook( i int, book *Books ){
    fmt.Printf( "Book%d title : %s\n",   i, book.title);
    fmt.Printf( "Book%d author : %s\n",  i, book.author);
    fmt.Printf( "Book%d subject : %s\n", i, book.subject);
    fmt.Printf( "Book%d book_id : %d\n", i, book.book_id);
}*/


/*func main() {
    //var numbers = make([]int,3,5)
    var numbers []int

    printSlice(numbers)

    if( numbers == nil ){
        fmt.Printf("切片是空的")
    }
}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}*/

/*func main(){
    var numbers []int
    printSlice(numbers)

    numbers = append( numbers, 0 )
    printSlice(numbers)

    numbers = append( numbers, 1 )
    printSlice( numbers )

    numbers1 := make([]int, len(numbers), (cap(numbers))*2)

    copy(numbers1, numbers)
}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}*/

/*func main(){
    nums := []int{2, 3, 4}
    sum  := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3{
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a":"apple", "b":"banana"}

    for k, v := range kvs {
        fmt.Printf( "%s -> %s\n", k, v )
    }

    for i, c := range "xiangk" {
        fmt.Printf("go %d, %d\n",i, c )
    }
}*/

/*func main(){
    var countryCapitalMap map[string]string
    countryCapitalMap = make(map[string]string)

    countryCapitalMap["x"] = "m"
    countryCapitalMap["k"] = "n"

    for k, v := range countryCapitalMap{
        fmt.Println( k, "值是:", v)
    }

    captial, ok := countryCapitalMap["x"]

    if (ok) {
        fmt.Println("x is", captial)
    } else {
        fmt.Println("x m", captial)
    }

    delete(countryCapitalMap, "x")

    for k, v := range countryCapitalMap{
        fmt.Println("k", k, "v", v )
    }
}*/

/*func Factorial( n uint64 )( result uint64 ){
    if( n > 0 ){
        result = n * Factorial( n - 1 )
        return result
    }
    return 1
}

func main(){
    var i uint64 = 15
    fmt.Printf("%d 的階乘是 %d\n", i, Factorial(i))
}*/

/*func fibonacci( n int ) int {
    if n < 2 {
        return n
    }
    return fibonacci(n - 2) + fibonacci( n - 1 )
}

func main(){
    var i int
    for i = 0; i < 10; i++{
        fmt.Printf("%d\t", fibonacci( i ) )
    }
}*/

/*type Phone interface{
    call()
}

type NokiaPhone struct {
}

func ( nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call(){
    fmt.Println("I am iPhone, I can call you!")
}

func main(){
    var phone Phone
    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()
}*/

/*type DividerError struct {
    dividee int
    divider int
}

func ( de *DividerError ) Error() (int, string) {
    strFormat := `Cannot proceed, the divider is zero.
                  devidee: %d
                  devider: 0`
    return -1, fmt.Sprintf( strFormat, de.dividee )
}

func Divide( varDividee int, varDivider int ) ( result int, errorMsg string ){
    if varDivider == 0{
        dData := DividerError{
            dividee: varDividee,
            divider: varDivider,
        }
        result, errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider, ""
    }
}

func main(){
    if result, errorMsg := Divide( 100, 10 ); errorMsg == ""{
        fmt.Println("100/10 = ", result)
    }
    if result, errorMsg := Divide(100, 0);errorMsg != ""{
        fmt.Println("errorMsg is:", errorMsg, "result is", result )
    }
}*/

/*func say( s string ) {
    for i := 0; i < 5; i++ {
        time.Sleep( 100 * time.Millisecond )
        fmt.Println( s )
    }
}

func main(){
    go say( "world" )
    say("hello" )
}*/

/*func sum( s []int, c chan int ){
    sum := 0
    for _, v := range s {
        sum += v
    }
    println( sum )
    c <- sum
}

func main() {
    s := []int{7, 2, 8, -9 ,4, 0}

    c := make( chan int )
    go sum( s[:len(s)/2], c )
    go sum( s[len(s)/2:], c )
    x, y := <- c, <- c

    fmt.Println( x, y, x+y )
}*/

/*func main(){
    ch := make( chan int )

    go func() {
        v := <- ch
        fmt.Println( v )
    }()

    ch <- 1

    fmt.Println("2")
}*/

/*func produce( p chan <- int, wg *sync.WaitGroup ){
    for i := 0; i < 10; i++ {
        p <- i
        fmt.Println("send:", i)
    }
    wg.Done()
}

func consumer( c <-chan int, wg *sync.WaitGroup ){
    for i := 0; i < 10; i++{
        v := <- c
        fmt.Println("receive:", v)
    }
    wg.Done()
}

func main(){
    var wg sync.WaitGroup
    for i := 1; i < 3; i++{
        wg.Add(1 )
    }
    ch := make( chan int, 10 )
    go produce( ch, &wg )
    go consumer( ch, &wg )

    wg.Wait()
    fmt.Println("Game Over")
}*/



/*func main() {
    file, err := os.OpenFile("./test.log",  os.O_CREATE | os.O_WRONLY | os.O_APPEND,os.ModePerm)
    if err != nil {
        log.Fatalln(err)
    }
    logger := log.New(file, "", log.LstdFlags|log.Llongfile)
    logger.Println("日志1.")
    logger.Println("日志23")
}*/

/*import (
    log "./lib"
)

func main(){
    log.Config("./test.log", 1)
    log.Error("Is a bug, %d, %d", 1, 2)
}*/

