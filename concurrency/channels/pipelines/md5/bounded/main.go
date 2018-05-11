package bounded

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// A result is the product of reading and summing a file using MD5.
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// Emit the paths of regular files in the tree
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// Close paths channel after Walk returns
		defer close(paths)
		// No select needed, since errc is buffered
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk cancelled")
			}
			return nil
		})
	}()
	return paths, errc
}

// Start a fixed number of digester goroutines that recieve file names
// from paths and send results on channel c
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}: // Send data on c
		case <-done: //recieve on done
			return
		}
	}
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All closes done on return; may do so before
	// recieving all valuse from c and errc
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	// Starts a fixed number of goroutines to read and digest files
	c := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 850
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	// End of popeline

	m := make(map[string][md5.Size]byte)
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
