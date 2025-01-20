package channels

import "fmt"

func UnbufferedChannel() {
	fmt.Println("Создаем небуферизованный канал")
	ch := make(chan int)

	fmt.Println("Если попытаться записать в небуферизованный канал в единственной горутине, то будет deadlock")
	// ch <- 1

	fmt.Println("Записываем из другой горутины")
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	fmt.Println("deadlock не происходит, так как не все горутины заблокированы")

	fmt.Println("Вычитываем данные из канала")

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Если попытаться вычитать данные, которые не были записаны, получим deadlock")
	// <-ch

	close(ch)

	fmt.Println("Если поптытаться записать данные в закрытый канал, будет паника")
	// ch <- 2
	fmt.Println("Читать из закрытого канала разрешено, если канал пуст, вернется нулевое значение")
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Благодаря возожности чтения из закрытого канала, можно проверять, является ли канал закрытым")

	if val, ok := <-ch; !ok {
		fmt.Println("Получаем нулевое значение", val)
		fmt.Println("ok == false, значит канал закрыт")
	}

}
