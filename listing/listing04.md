Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод: от 0 до 9 и deadlock. 
Дедлок связан с тем, что канал не закрывается, а главное горутина main продолжает ждать из канала данные

```
