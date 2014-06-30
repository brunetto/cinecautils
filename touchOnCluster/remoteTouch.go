package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
	
	pwd "github.com/seehuhn/password"
	"github.com/brunetto/goutils/connection"
)

// RESOURCES:
// * https://godoc.org/code.google.com/p/go.crypto/ssh
// * http://dave.cheney.net/tag/golang-3/page/2
// * http://play.golang.org/p/3z513UKrOY
// * https://code.google.com/p/go/issues/detail?id=7787
// * https://groups.google.com/forum/#!topic/golang-nuts/QRsZSPqwDhM
// * http://play.golang.org/p/kMhHvbl4SG
// * https://docs.google.com/document/d/1nF2wlkIwuA4AXryOvE2p0hgQUbsyRYklKSot4ahH3Aw/edit
// * http://godoc.org/code.google.com/p/go.crypto/ssh/agent
// * https://groups.google.com/d/topic/golang-nuts/0JfeQ-Qu37U/discussion


func main () () {
	
	
	var (
		usr string
		pw []byte//string
		err error
	)
	
	fmt.Println("!!!EXPERIMENTAL VERSION!!!")
	fmt.Println("Use at your own risk!!!")
	fmt.Println("Your password should be safe, but I don't guarantee!!!")
	fmt.Println("Could harm your cats!!!")
	
	fmt.Println("=====================================")
	fmt.Println("Close with:")
	fmt.Println("kill -9 ", os.Getpid())
	fmt.Println("=====================================")
	
	fmt.Print("Please insert your username: ")
	fmt.Scan(&usr)
	
	fmt.Println()
	
	if pw, err = pwd.Read("Please insert your password: "); err != nil {
		// 	if pw, err = gopass.GetPass("Please insert your password: ") {
		log.Fatal("Error retrieving password; ", err)
	}
	
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	
	servers := []string{"login.eurora.cineca.it:22", "login.plx.cineca.it:22"}
	
	for {
		for _, server := range servers {
		
			// Each ClientConn can support multiple interactive sessions,
			// represented by a Session.
			log.Println("Start new session")
			session, err := connection.SshSessionWithPw(server, usr, string(pw))
			if err != nil {
				panic("Failed to create session: " + err.Error())
			}
			// 		defer session.Close()
			
			// Once a Session is created, you can execute a single command on
			// the remote side using the Run method.
			var b bytes.Buffer
			session.Stdout = &b
			cmd := "python Code/touchMyTrallalla.py" // "/usr/bin/whoami"
			log.Println("Run touchMyTrallalla.py")
			if err := session.Run(cmd); err != nil {
				panic("Failed to run: " + err.Error())
			}
			session.Close()
			log.Println(b.String())
			log.Println("Done")
		}
		waitingTime := time.Duration(24) * time.Hour
		log.Println("Waiting ", waitingTime)
		time.Sleep(waitingTime)
	}
}

