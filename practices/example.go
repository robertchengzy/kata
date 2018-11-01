package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main22() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex:sepIndex] //full slice expression
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)

	fmt.Println("new path =>", string(path))
}

func main3() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2, 4)

	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}

type myMutex sync.Mutex
type test sync.Once

func main4() {
	var mtx myMutex
	fmt.Printf("%+v", mtx)
}

type myLocker1 struct {
	sync.Mutex
}

func main5() {
	var lock myLocker1
	lock.Lock()   //ok
	lock.Unlock() //ok
}

type myLocker2 sync.Locker

func main6() {
	var lock myLocker2 = new(sync.Mutex)
	lock.Lock()   //ok
	lock.Unlock() //ok
}

func main7() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
			//fallthrough
		case false:
			fmt.Printf("1")
			goto retry
		}
	}
retry:
	fmt.Println("out!")
}

func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		fmt.Println(1, v)
		go func() {
			fmt.Println(v)
		}()
		/*go func(i string) {
			fmt.Println(i)
		}(v)*/
	}

	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main9() {
	data := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		v := v
		go v.print()
	}

	time.Sleep(3 * time.Second)
}

func main10() {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }())
	i++
	//prints: result => 2 (not ok if you expected 4)
}

func main1111() {
	var s []*string
	for i := 0; i < 10; i++ {
		func() {
			defer fmt.Println("test", i)
			fmt.Println("testz")
		}()
		fmt.Println("testy")
		//s = append(s, "1")
	}
	s = nil
	fmt.Println("testx")
	fmt.Println(s)
}

func main111() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}

	start, err := os.Stat(os.Args[1])
	if err != nil || !start.IsDir() {
		os.Exit(-1)
	}

	var targets []string
	filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, fpath)
		return nil
	})

	for _, target := range targets {
		func() {
			f, err := os.Open(target)
			if err != nil {
				fmt.Println("bad target:", target, "error:", err)
				return
			}
			defer f.Close() //ok
			//do something with the file...
		}()
	}
}
