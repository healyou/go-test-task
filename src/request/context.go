package request

import (
	"sync"
)

// класс осуществляет запуск функций как горутины, при этом можно ограничить максимальное кол-во горутин
type SearchContext struct {
	wg              *sync.WaitGroup
	ch              chan int
	sum             *int
	runGorCnt       *int
	maxGoroutineCnt *int
}

// создаём новый контекст для запуска горутин, channel ограничен максимальным кол-вом горутин
func NewContext(maxGoroutineCnt int) SearchContext {
	sum := 0
	runGorCnt := 0
	return SearchContext{&sync.WaitGroup{}, make(chan int, maxGoroutineCnt), &sum, &runGorCnt, &maxGoroutineCnt}
}

// запуск функции поиска как горутины
func (context SearchContext) runFuncAsGoroutine(f func(string, string) int, search string, url string) {
	*context.runGorCnt++
	context.wg.Add(1)
	context.runAsGoroutine(f, search, url)
}

// запуск горутины с освобожением WaitGroup и записью результата в channel
func (context SearchContext) runAsGoroutine(f func(string, string) int, search string, url string) {
	defer context.wg.Done()
	var count = f(url, search)
	context.ch <- count
}

// ожидание завершения всех горутин и подстчёт значений channel
func (context SearchContext) waitAndRelease() {
	context.wg.Wait()
	context.releaseGoroutines()
}

// освобождение данных канала для создания новых горутин
func (context SearchContext) releaseGoroutines() {
	if len(context.ch) > 1 {
		for len(context.ch) > 0 {
			*context.sum += <-context.ch
			*context.runGorCnt--
		}
	} else {
		*context.sum += <-context.ch
		*context.runGorCnt--
	}
}

// можно ли создать ещё горутины
func (context SearchContext) isLimitGoroutine() bool {
	return *context.runGorCnt >= *context.maxGoroutineCnt
}

// вернуть сумму подстёча всех горутин
func (context SearchContext) getSum() int {
	return *context.sum
}
