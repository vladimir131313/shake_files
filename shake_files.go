package main

import (
	"fmt"
	"os"
	"log"
	"crypto/rand"
	"math/big"
	"flag"
	"regexp"
	"path"
)


func main() {

	var dirname string
	var rename bool
	var rollback bool
	var renamedir bool
	flag.StringVar(&dirname,"p", ".", "path to directory with files")
	flag.BoolVar(&rename,"w", false, "Rename files (Default: Just print result without rename files)")
	flag.BoolVar(&rollback,"r", false, "Delete prefix in filenames if it exist")
	flag.BoolVar(&renamedir,"d", false, "Rename dirs too")
	flag.Parse()

	fdirname, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	fnames, err := fdirname.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}


	for _,fname := range fnames {
		if ! renamedir {
			fi, err := os.Stat(path.Join(dirname,fname))
				if err != nil {
				log.Fatal(err)
			}
			if fi.Mode().IsDir() {
				continue
			}
		}
		prf := "$2"
		if ! rollback {
			r,err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				log.Fatal(err)
			}
			prf = r.String()
			prf += "RND-$2"
		}
		re := regexp.MustCompile(`^(\d+RND-)?(.+)$`)
		new_fname := re.ReplaceAllString(fname, prf)
		if rename {
			err = os.Rename(path.Join(dirname,fname),path.Join(dirname,new_fname))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("%-50s --->       %s\n",path.Join(dirname,fname),path.Join(dirname,new_fname))
		}
	}
}
