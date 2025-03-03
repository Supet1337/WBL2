Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2 1
Defer Вызываются по порярдку, после возврата из функции
В случае с функцией test(), отложенная функция инкрементирует значение x после установки значения x = 1, 
но перед выполнением инструкции return. 
Поэтому значение x увеличивается на 1, и функция возвращает значение 2.

В функции anotherTest(), отложенная функция также выполняется перед инструкцией return. 
Она инкрементирует значение x, но поскольку инструкция return возвращает значение x до выполнения отложенной функции, 
функция возвращает исходное значение x, равное 1.

```
