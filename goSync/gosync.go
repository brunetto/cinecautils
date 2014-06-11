package main

import (
	"fmt"
	"log"
// 	"io"
	"os"
	"os/exec"
	"regexp"
	"time"
	
// 	"code.google.com/p/gopass"
// 	"gopkg.in/pipe.v2"	
	"github.com/brunetto/goutils/debug"
)

// FIXME: see http://godoc.org/labix.org/v2/pipe, http://labix.org/pipe, 

func main () () {
	defer debug.TimeMe(time.Now())
	// kill the process? http://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
	
	var ( 
		waitingTime time.Duration = time.Duration(1) * time.Minute
		
// 		pw string
		err error

		user string 
		host string 
		source string 
		destination string
		
		regString string         = `(\S+)@(\S+):(\S+)`
		regExp    *regexp.Regexp = regexp.MustCompile(regString)
		regRes []string
	)
	
	if len(os.Args) < 3 {
		log.Fatal("Use like: ./gosync user@host:/source dest")
	} 

	log.Println("Try to detect user cli preferences")
	if regRes = regExp.FindStringSubmatch(os.Args[1]); regRes == nil {
		log.Fatal("Can't extract info in ", os.Args[1])
	}
	user = regRes[1]
	host = regRes[2]
	source = regRes[3]
	destination = os.Args[2]

	log.Println("\nIn case of problems run")
	fmt.Println("-------------------")
	fmt.Printf("kill -9 %v\n", os.Getpid())
	fmt.Println("-------------------")
	
	log.Println("No password request because assuming keys exchanges.")
	
// 	if pw, err = gopass.GetPass("Please insert your password: "); err != nil {
// 		log.Fatal("Error retrieving password; ", err)
// 	}
	
	log.Println("Syncing info: ")
	fmt.Println("user: ", user)
	fmt.Println("host: ", host)
	fmt.Println("source: ", source)
	fmt.Println("destination: ", destination)
	
	for {
		cmd := exec.Command("/usr/bin/rsync", "-avuhz", user+"@"+host+":"+source, destination)
		
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
// 		iow, _ := cmd.StdinPipe()
		
		if err = cmd.Start(); err != nil {
			log.Fatal("Sync start: ", err)
		}
		
// 		io.WriteString(iow, pw)
// 		iow.Close()
		
		if err := cmd.Wait(); err != nil {
			log.Println("Sync wait: ", err)
		}
		
		log.Println("Assuming rsync killed by cineca")
		
		log.Println("Waiting ", waitingTime)
		time.Sleep(waitingTime)
	}
}

