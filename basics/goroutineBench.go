package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

func goroutineBench() {
	recreate()

	start := time.Now()
	concurrent()
	elapsed := time.Since(start)

	fmt.Println("Concurrent time: ", elapsed)

	recreate()

	start = time.Now()
	synchronous()
	elapsed = time.Since(start)

	fmt.Println("Synchronous time: ", elapsed)

	recreate()
}

func synchronous() {
	for b := 0; b < 1000000; b++ {
		f, err := os.Create(fmt.Sprintf("files/%d.txt", b))

		if err != nil {
			log.Fatalln(err)
		}

		f.Close()
	}
}

func concurrent() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			for b := 0; b < 10000; b++ {
				f, err := os.Create(fmt.Sprintf("files/%d%d.txt", b, num))

				if err != nil {
					log.Fatalln(err)
				}

				f.Close()
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func recreate() {
	files, _ := ioutil.ReadDir("files")

	for _, f := range files {
		if err := os.Remove(fmt.Sprintf("files/%s", f.Name())); err != nil {
			fmt.Println(err)
			log.Fatalln(err)
		}
	}

	if err := os.RemoveAll("files"); err != nil {
		log.Fatalln(err)
	}

	if err := os.Mkdir("files", os.ModePerm); err != nil {
		log.Fatalln(err)
	}

}
