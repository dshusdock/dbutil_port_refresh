package ssh

import (
	//"bufio"
	"bytes"
	//"errors"
	"fmt"
	"log"
	//"os"
	//"path/filepath"
	//"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type SSHClient struct {
}


func NewSSHClient(ip string, user string, pwd string) *ssh.Client {
	// hostKeyCallback, err := knownhosts.New("/home/dave/.ssh/known_hosts")
	// if err != nil {
    // 	log.Fatal(err)
	// }

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //hostKeyCallback /*ssh.FixedHostKey(hostKey)*/,
		// or this --> HostKeyCallback: ssh.InsecureIgnoreHostKey(), // this is not recommended
	}
	//"10.201.205.221:22"
	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		fmt.Println("Failed to dial: ", err)
		return nil
	}
	
	return client
}

func RunCommand(client *ssh.Client, command string) (string, error) {
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		fmt.Println("Failed to run command1:", err)
		//return "error", err
	}

	return b.String(), nil
}

//ssh -o HostkeyAlgorithms=+ssh-rsa,ssh-dss user@hostname
//scp -o HostkeyAlgorithms=+ssh-rsa,ssh-dss user@hostname

func TestSSHClient() {

	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.

	// Found this example on the internet, seems to work
	// https://stackoverflow.com/questions/45441735/ssh-handshake-complains-about-missing-host-key
	hostKeyCallback, err := knownhosts.New("/home/dave/.ssh/known_hosts")
	if err != nil {
    	log.Fatal(err)
	}

	config := &ssh.ClientConfig{
		User: "dave",
		Auth: []ssh.AuthMethod{
			ssh.Password("All4one!"),
		},
		HostKeyCallback: hostKeyCallback /*ssh.FixedHostKey(hostKey)*/,
		// or this --> HostKeyCallback: ssh.InsecureIgnoreHostKey(), // this is not recommended
	}
	client, err := ssh.Dial("tcp", "10.201.205.221:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/ls -la"); err != nil {
		log.Fatalf("Failed to run command1: %s", err)
	}

	fmt.Println(b.String())
}








/*
func getHostKey(host string) (ssh.PublicKey, error) {
    file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var hostKey ssh.PublicKey
    for scanner.Scan() {
        fields := strings.Split(scanner.Text(), " ")
        if len(fields) != 3 {
            continue
        }
        if strings.Contains(fields[0], host) {
            var err error
            hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
            if err != nil {
                return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
            }
            break
        }
    }

    if hostKey == nil {
        return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
    }
    return hostKey, nil
}
*/