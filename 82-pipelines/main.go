package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func main() {
	//examplePipeline()
	//exampleMerge()
	//exampleDigest()
	exampleDigestWithChannels()
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}

		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}

		}
	}()

	return out
}

func examplePipeline() {
	done := make(chan struct{})
	defer close(done)

	c := gen(done, 2, 3, 4)
	out := sq(done, c)
	for n := range out {
		fmt.Println(n)
	}
}

func merge(done <-chan struct{}, channels ...<-chan int) <-chan int {
	out := make(chan int)

	var waitGroup sync.WaitGroup

	// function that merges a single channel into out
	mergeSingle := func(c <-chan int) {
		defer waitGroup.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}

		}
	}

	waitGroup.Add(len(channels))

	for _, c := range channels {
		go mergeSingle(c)
	}

	// wait in a new goroutine, so that the merge function returns before it
	go func() {
		waitGroup.Wait()
		close(out)
	}()

	return out
}

func exampleMerge() {
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3, 4)

	// fan-out
	// Divide processing between two channels, c1 and c2
	c1 := sq(done, in)
	c2 := sq(done, in)

	// fan-in
	for n := range merge(done, c1, c2) {
		fmt.Println(n)
	}
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte, 0)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}

func exampleDigest() {
	m, err := MD5All(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%s %x\n", path, m[path])
	}
}

type Result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan Result, <-chan error) {
	c := make(chan Result)
	errc := make(chan error, 1) // buffered channel, won't block in the first errc<-

	go func() {
		var waitGroup sync.WaitGroup

		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				data, err := os.ReadFile(path)
				select {
				case c <- Result{path: path, sum: md5.Sum(data), err: err}:
				case <-done:
					return
				}
			}()

			select {
			case <-done:
				return errors.New("walk cancelled")
			default:
				return nil
			}
		})

		go func() {
			waitGroup.Wait()
			close(c)
		}()

		errc <- err
	}()

	return c, errc
}

func MD5All2(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)
	m := make(map[string][md5.Size]byte, 0)

	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil
}

func exampleDigestWithChannels() {
	m, err := MD5All2(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%s %x\n", path, m[path])
	}
}
