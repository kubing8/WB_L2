Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error

Связано с тем, что у нас возращается не пустой интерфейс, т.е. [*customError, nil]
А сравнивается это с [nil, nil]. 
Это связано со стуктурами интерфейса и пустого интерфейса

```
