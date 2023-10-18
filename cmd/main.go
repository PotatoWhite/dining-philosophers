package main

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Kant", leftFork: 4, rightFork: 0},
	{name: "Marx", leftFork: 0, rightFork: 1},
	{name: "Hegel", leftFork: 1, rightFork: 2},
	{name: "Nietzsche", leftFork: 2, rightFork: 3},
	{name: "Aristotle", leftFork: 3, rightFork: 4},
}

var hunger = 30

func main() {
	fmt.Println("모여서 밥먹자~")

	dine()

	fmt.Println("다먹었네~")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go eat(philosophers[i], forks, wg, seated)
	}

	wg.Wait()
}

func eat(p Philosopher, forks map[int]*sync.Mutex, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()

	seated.Done()
	seated.Wait()

	fmt.Printf("%s 먹기 시작\n", p.name)

	for i := 0; i < hunger; i++ {

		if p.leftFork > p.rightFork {

			forks[p.rightFork].Lock()
			fmt.Printf("%s 오른쪽 포크 집기 %d \n", p.name, p.rightFork)
			forks[p.leftFork].Lock()
			fmt.Printf("%s 왼쪽 포크 집기 %d \n", p.name, p.leftFork)

		} else {

			forks[p.leftFork].Lock()
			fmt.Printf("%s 왼쪽 포크 집기 %d \n", p.name, p.leftFork)
			forks[p.rightFork].Lock()
			fmt.Printf("%s 오른쪽 포크 집기 %d \n", p.name, p.rightFork)
		}

		fmt.Printf("%s 먹는 중\n", p.name)
		// time.Sleep(1 * time.Second)

		fmt.Printf("%s 왼쪽 포크 내려놓기 %d \n", p.name, p.leftFork)
		forks[p.leftFork].Unlock()

		fmt.Printf("%s 오른쪽 포크 내려놓기 %d \n", p.name, p.rightFork)
		forks[p.rightFork].Unlock()
	}

	fmt.Printf("%s 다 먹었습니다.\n", p.name)
}
