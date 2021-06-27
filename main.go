package main

import (
	"bufio"
	"container/list"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	calc "helloworld/functions" // $GOPATH/src开始继续写
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//声明全局变量
var (
	name string
	age  int = 15
	isOk bool
	//类型推导，有赋值可省略类型
	name2 = "abcd"
	wg    sync.WaitGroup
	lock  sync.Mutex //互斥锁
	wg2   sync.WaitGroup
)

//声明全局常量
const (
	pi       = 3.1415926
	NOTFOUND = 404
)

//iota是从0开始计数的常量计数器，每增加1行加1
const (
	a1     = iota       //0
	a2                  //1
	a3     = 100        //100
	_                   //3
	a4                  //4
	a5, a6 = iota, iota //5,5
)

func main() {
	fmt.Println("------------------------ fmt ------------------------")
	//赋值
	name = "理想"
	age = 16
	isOk = true
	//格式化输出
	fmt.Printf("name:%s\n", name)

	//简短变量声明,仅函数内使用
	name3 := "qwe"
	fmt.Printf("name3:%s\n", name3)

	//字符长度len（）
	fmt.Println(len(name3))

	//类型转换
	n1 := 100
	fmt.Printf("%T %v\n", strconv.Itoa(n1), strconv.Itoa(n1)) //int to array, n1的int类型转为string类型

	n2 := "10000"
	retInt, _ := strconv.Atoi(n2) ////array to int, n2的string类型转为int类型
	fmt.Printf("%T %v\n", retInt, retInt)

	fmt.Println("------------------------ if for switch ------------------------")

	age2 := 19
	if age2 > 18 {
		fmt.Println("已成年")
	} else if age2 > 12 {
		fmt.Println("青少年")
	} else {
		fmt.Println("小孩")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %v\n", i)
		if i == 2 {
			break //跳出for循环
		}
	}

	var i2 = 0
	for i2 < 2 { //作用大致为while
		fmt.Printf("i2 = %v\n", i2)
		i2++
	}

	s2 := "hello"
	for i, v := range s2 {
		fmt.Printf("%v %c\n", i, v)
	}

	var i3 = 3
	switch i3 {
	case 1:
		fmt.Println("i3 = 1")
	case 2:
		fmt.Println("i3 = 2")
	case 3, 4, 5:
		fmt.Println("i3 = 3 or 4 or 5")
	default:
		fmt.Println("nil")
	}

	fmt.Println("------------------------ Array，切片,追加，复制，删除，排序 ------------------------")
	//声明方式1
	var a3 [3]int = [3]int{1, 2, 3}
	fmt.Println("a3: ", a3)
	fmt.Println("a3[1]: ", a3[1])
	//声明方式2
	a4 := [...]int{2, 3, 4, 5, 6}
	fmt.Println("a4: ", a4)
	//切片
	a5 := []int{5, 4, 3, 2, 1}
	s5 := a5[0:4] //数组切割第0至4位，左包含右不包含（ 从头或尾切割可省略，等于a5[:4] ）
	fmt.Println("a5: ", a5)
	fmt.Println("s5: ", s5)
	//追加
	a5 = append(a5, 999)
	fmt.Println("a5: ", a5)
	//copy, 复制而非引用
	a6 := make([]int, 6, 6)
	copy(a6, a5)
	fmt.Println("a6: ", a6)
	//删除，如要删除位置为1的元素
	a6 = append(a6[:1], a6[2:]...)
	fmt.Println("a6: ", a6)
	//排序
	sort.Ints(a6[:])
	fmt.Println("a6: ", a6)
	//切割
	a7 := "how,do,you,do"
	a8 := strings.Split(a7, ",") //以","切割
	fmt.Println(a8)

	fmt.Println("------------------------ List链表，删除，遍历输出 ------------------------")
	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushFront(2)

	for element := mylist.Front(); element != nil; element = element.Next() {
		if element.Value != 1 {
			mylist.Remove(element) //删除链表中值不为1的元素
		}
	}

	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value) //遍历输出当前链表中的元素
	}

	fmt.Println("------------------------ 指针 ------------------------")
	b1 := 18
	fmt.Println("&b1: ", &b1)   //b1的地址
	fmt.Println("*&b1: ", *&b1) //取b1的地址的值

	fmt.Println("------------------------ map ------------------------")
	//声明方式1
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["age"] = 18
	m1["ID"] = 123
	fmt.Println(m1)
	//声明方式2
	m2 := make(map[string]int)
	m2["age"] = 10
	m2["ID"] = 456
	fmt.Println("the ID of m2: ", m2["ID"])
	//删除
	delete(m2, "ID")

	fmt.Println("------------------------ 函数,defer ------------------------")

	fmt.Println(sum1(1, 2))
	abc3("test", 1, 2, 3, 4, 5) //可变长参数
	deferExample()              //defer

	calc.Add(10, 20)

	fmt.Println("------------------------ 结构体，方法，继承 ------------------------")
	//声明方式1
	var p person
	p.name = "Alex"
	p.age = 18
	p.gender = "Male"
	p.hobby = []string{"basketball", "football", "swimming"}
	fmt.Println("p: ", p)
	fmt.Println("p.name: ", p.name)
	//声明方式2
	var p2 = person{
		"Bob",
		72,
		"Female",
		[]string{"archer", "play"},
	}
	fmt.Println("p2 ", p2)

	//方法， 仅允许接收person类型
	p.personSayHello()

	//继承, man 也属于person，可调用person类型方法，但是多了strongPower这个值
	p3 := man{
		strongPower: 1,
		person:      person{name: "Cesar"},
	}
	p3.personSayHello()

	//JSON
	c1 := car{
		"Benz",
		"Black",
		10000,
	}
	// JSON 序列化， 结构体struct 转为 JSON
	b, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json.Marshal failed,err: ", err)
		return
	}
	fmt.Println("b: ", string(b))
	// JSON 反序列化， JSON 转为 结构体struct
	jsonString := `{"Logo":"BMW","Color":"Green","Price":100000}`
	var c2 car
	json.Unmarshal([]byte(jsonString), &c2)
	fmt.Println("c2: ", c2)

	fmt.Println("------------------------ 接口 ------------------------")

	//不管什么结构体，有run()方法就可以，都算runable接口的
	p3.run()
	c1.run()

	//空接口，可接收任意类型
	interfaceShow("abc")
	interfaceShow(777)

	fmt.Println("------------------------ IO ------------------------")

	// 获取输入
	// fmt.Println("Please Input some String")
	// var inputString string
	// fmt.Scan(&inputString)
	// fmt.Println("you input: ", inputString)

	// 获取输入（Bufio版）
	// fmt.Println("Please Input some String")
	// reader := bufio.NewReader(os.Stdin)
	// readerString, _ := reader.ReadString('\n')
	// fmt.Println("you input: ", readerString)

	//写入文件
	//writeStringtoFile()

	//写入文件（Bufio版）
	//writeStringtoFilebyBufio()

	//写入文件（Ioutil版）
	//writeStringtoFilebyIoutil()

	//读取文件（Bufio版），一行一行处理
	//readFromFilebyBufio()

	//读取文件（Ioutil版），整个文件处理
	//readFromFilebyIoutil()

	fmt.Println("------------------------ Time ------------------------")
	//时间
	fmt.Println("Year now is: ", time.Now().Year())
	fmt.Println("Minute now is: ", time.Now().Minute())

	//时间戳, 1970年1月1日至现在的秒数
	fmt.Println("时间戳： ", time.Now().Unix())

	//时间戳转变为时间
	fmt.Println("时间戳转变为时间: ", time.Unix(1613304836, 0))

	// now + 24小时
	fmt.Println("now + 24小时: ", time.Now().Add(24*time.Hour))

	// 定时器, 1秒钟执行一次
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	//等待Sleep
	//time.Sleep(5 * time.Second)

	//时间格式化，2021-02-14
	fmt.Println(time.Now().Format("2006-01-02"))
	//时间格式化，2021/02/14 20:31:04
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))

	fmt.Println("------------------------ 并发 goroutine channel ------------------------")
	//并发hello()这个函数
	go hello()

	//同步 sync.WaitGroup, 保证所有go线程结束后再继续wg.Wait()后内容
	wg.Add(1) //go线程计数器+1
	go helloSyncWait()
	wg.Wait() //go线程计数器归0后再继续

	//通道 channel
	chan1 := make(chan int, 1) //chan1通道最多放1个int值
	chan1 <- 20
	chan1Value := <-chan1
	fmt.Println("chan1: ", chan1Value)
	close(chan1) //关闭通道，可继续读，不可继续写

	//worker pool,指定开启goroutine的数量
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启2个goroutine
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}
	//一共3个任务
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	//输出结果
	for i := 1; i <= 3; i++ {
		<-results
	}

	//select,哪个case能走通就走哪个, 多个case能走通则随机走一个
	chan2 := make(chan int, 10)
	for i := 0; i < 5; i++ {
		select {
		case x := <-chan2:
			fmt.Println(x)
		case chan2 <- i:
		}
	}

	//sync.Mutex , 互斥锁, lock.Lock()后只有1个goroutine进入临界区运行,其他的goroutine等待. lock.Unlock()后解锁.
	//atomic, 原子操作,并发安全版互斥锁, 写入 atomic.StoreInt64() , 相加 atomic.AddInt64(&n, 20) , 交换 atomic.SwapInt64()

	//sync.Map, 并发安全版map
	var syncMap1 = sync.Map{}
	wg2 := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncMap1.Store(key, n)                 // 存储
			syncMap1Value, _ := syncMap1.Load(key) // 读取
			fmt.Printf("key: %v, value: %v\n", key, syncMap1Value)
			wg2.Done()
		}(i)
	}
	wg2.Wait()

	//contaxt, 控制子goroutine,timeout后直接结束
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	go contaxtWorker(ctx)
	time.Sleep(time.Second * 3)
	ctxCancel()
	fmt.Println("contaxt goroutine is cancelled")

	fmt.Println("------------------------ flag ------------------------")
	//flag,执行文件时可传入参数, 如命令行执行: helloworld.exe -name alex
	name4 := flag.String("name", "默认名", "请输入名字")
	flag.Parse()
	fmt.Println("name4: ", *name4)

}

func sum1(x int, y int) int {
	return x + y
}

func sum2(x int, y int) (ret int) {
	ret = x + y
	return
}

//可变长参数
func abc3(x string, y ...int) {
	fmt.Println(x, " + ", y)
}

func deferExample() {
	fmt.Print("One")
	defer fmt.Print("Two")   //defer语句会延迟到函数即将返回的时候执行
	defer fmt.Print("Three") //defer执行顺序为先进后出
	fmt.Print("Four")
}

type person struct { //结构体
	name   string
	age    int
	gender string
	hobby  []string
}

//方法， 仅允许接收person类型
func (p person) personSayHello() {
	fmt.Printf("%s:Hello\n", p.name)
}

//继承, man 也属于person，可调用person类型方法，但是多了strongPower这个值
type man struct {
	strongPower int
	person
}

//JSON, 变量名需要首字母大写，以便json.Marshal和json.Unmarshal方法访问
type car struct {
	Logo  string
	Color string
	Price int
}

//接口interface,有run()方法就可以，都算runable接口的
type runable interface {
	run()
}

func (m man) run() {
	fmt.Println("man run")
}

func (c car) run() {
	fmt.Println("car run")
}

//空接口，可接收任意类型
func interfaceShow(a interface{}) {

	//类型断言，根据接收类型分别处理
	switch t := a.(type) {
	case string:
		fmt.Printf("case string: type: %T , value: %v\n", t, t)
	case int:
		fmt.Printf("case int: type: %T , value: %v\n", t, t)
	default:
		fmt.Printf("cannot find")
	}

}

//写入文件
func writeStringtoFile() {
	fileObj, err := os.OpenFile("./somewords.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()

	fileObj.WriteString("I would like to add a string\n")
}

//写入文件（Bufio版）
func writeStringtoFilebyBufio() {
	fileObj, err := os.OpenFile("./somewords.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()

	wr := bufio.NewWriter(fileObj)
	wr.WriteString("I would like to add a string by bufio\n") //写到缓存中
	wr.Flush()                                                //将缓存内容写入文件
}

//写入文件（Ioutil版）
func writeStringtoFilebyIoutil() {
	str := "I would like to add a string by ioutil\n"
	err := ioutil.WriteFile("./somewords.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed, err: %v", err)
		return
	}
}

//读取文件（Bufio版），一行一行处理
func readFromFilebyBufio() {
	fileObj, err := os.Open("./somewords.txt") //打开文件
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close() //最后执行关闭文件

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') //文件每行输出为String类型，即变量line
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err: %v", err)
			return
		}
		fmt.Print(line)
	}
}

//读取文件（Ioutil版），整个文件处理
func readFromFilebyIoutil() {
	content, err := ioutil.ReadFile("./somewords.txt") //文件直接读取
	if err != nil {
		fmt.Printf("read line failed, err: %v", err)
		return
	}
	fmt.Print(string(content))

}

func hello() {
	fmt.Println("hello")
}

//go同步 sync.WaitGroup
func helloSyncWait() {
	defer wg.Done() //go线程计数器-1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		r1 := rand.Intn(10) // 0 <= r1 < 10 随机数
		fmt.Printf("random Int: %v \n", r1)
	}
}

//worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker: %v start job:%v\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker: %v end job:%v\n", id, j)
		results <- j * 2
	}
}

//contaxt, 控制子goroutine,timeout后直接结束
func contaxtWorker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("do something......")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50ms后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("contaxt goroutine is timeout")
}
