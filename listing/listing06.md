Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3,2,3]
Рассмотрим работу функции modifySlice
Получаем слайс s и копируем его в i
т.к слайс это структура, то указатель на базовый массив также скопируется
i[0] = "3" // Изменили слайс s и i
s : [3,2,3]
i : [3,2,3]
i = append(i, 4) // Теперь i больше не ссылается на тот же массив что и s, т.к мы переполнили capacity и создался новый массив
s : [3,2,3]
i : [3,2,3,4]
i[1] = 5
s : [3,2,3]
i : [3,5,3,4]
i = append(i, "6") //Опять пересоздаём слайс и получаем внтури ссылку на новый массив
s : [3,2,3]
i : [3,5,3,4,6]

```
