package internal

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Gateway struct {
	// SSH connection settings
	SSHAuthType string
	SSHHost     string
	SSHPort     int
	SSHUser     string
	SSHPassword string
	SSHKeyPath  string

	// Target example settings
	DBHost string
	DBPort int

	// Local listener settings
	LocalHost string
	LocalPort int

	// SSH client
	sshClient *ssh.Client
	mutex     sync.RWMutex
}

func NewGateway() *Gateway {
	return &Gateway{
		mutex: sync.RWMutex{},
	}
}

func (g *Gateway) connectSSH() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.sshClient != nil {
		return nil
	}

	// Read private key
	key, err := os.ReadFile(g.SSHKeyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: g.SSHUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: Use ssh.FixedHostKey in production
		Timeout:         10 * time.Second,
	}

	// Connect to SSH server
	g.sshClient, err = ssh.Dial("tcp",
		fmt.Sprintf("%s:%d", g.SSHHost, g.SSHPort),
		config,
	)
	if err != nil {
		return fmt.Errorf("failed to dial SSH server: %v", err)
	}

	return nil
}

func (g *Gateway) connectSSHBasic() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.sshClient != nil {
		return nil
	}

	config := &ssh.ClientConfig{
		User: g.SSHUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(g.SSHPassword), // Using password authentication
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: Use ssh.FixedHostKey in production
		Timeout:         10 * time.Second,
	}

	// Connect to SSH server
	var err error
	g.sshClient, err = ssh.Dial("tcp",
		fmt.Sprintf("%s:%d", g.SSHHost, g.SSHPort),
		config,
	)
	if err != nil {
		return fmt.Errorf("failed to dial SSH server: %v", err)
	}

	return nil
}

func (g *Gateway) Start() error {
	// Create local listener
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", g.LocalHost, g.LocalPort),
	)
	if err != nil {
		return fmt.Errorf("failed to start local listener: %v", err)
	}
	defer listener.Close()

	log.Printf("Gateway listening on %s:%d", g.LocalHost, g.LocalPort)
	log.Printf("Tunneling to %s:%d through %s:%d",
		g.DBHost, g.DBPort, g.SSHHost, g.SSHPort)

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Connection acceptance loop
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Printf("Failed to accept connection: %v", err)
				continue
			}
			go g.handleConnection(client)
		}
	}()

	// Wait for shutdown signal
	<-stop
	log.Println("Shutting down gateway...")
	return nil
}

func (g *Gateway) handleConnection(local net.Conn) {
	defer local.Close()

	// Ensure SSH connection is established
	if g.SSHAuthType == "key" {
		err := g.connectSSH()
		if err != nil {
			log.Printf("Failed to establish SSH connection: %v", err)
			return
		}
	} else {
		err := g.connectSSHBasic()
		if err != nil {
			log.Printf("Failed to establish SSH connection: %v", err)
			return
		}
	}

	// Connect to remote example through SSH tunnel
	remote, err := g.sshClient.Dial("tcp",
		fmt.Sprintf("%s:%d", g.DBHost, g.DBPort),
	)
	if err != nil {
		log.Printf("Failed to connect to remote example: %v", err)
		return
	}
	defer remote.Close()

	// Start bidirectional copy
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(remote, local)
	}()

	go func() {
		defer wg.Done()
		io.Copy(local, remote)
	}()

	wg.Wait()
}

func App(gateway *Gateway) {

	// Parse command line flags
	//flag.StringVar(&gateway.SSHHost, "ssh-host", "jump-server.example.com", "SSH jump server hostname")
	//flag.IntVar(&gateway.SSHPort, "ssh-port", 22, "SSH port")
	//flag.StringVar(&gateway.SSHUser, "ssh-user", "username", "SSH username")
	//flag.StringVar(&gateway.SSHPassword, "ssh-password", "password", "SSH password")
	//flag.StringVar(&gateway.SSHKeyPath, "ssh-key", "~/.ssh/id_rsa", "Path to SSH private key")
	//
	//flag.StringVar(&gateway.DBHost, "db-host", "private-db.internal", "Target example hostname")
	//flag.IntVar(&gateway.DBPort, "db-port", 5432, "Target example port")
	//
	//flag.StringVar(&gateway.LocalHost, "local-host", "127.0.0.1", "Local binding address")
	//flag.IntVar(&gateway.LocalPort, "local-port", 5432, "Local binding port")
	//
	//flag.Parse()

	// Expand home directory in key path
	if gateway.SSHAuthType == "key" {
		if gateway.SSHKeyPath[:2] == "~/" {
			home, err := os.UserHomeDir()
			if err == nil {
				gateway.SSHKeyPath = home + gateway.SSHKeyPath[1:]
			}
		}
	}

	// Start the gateway
	if err := gateway.Start(); err != nil {
		log.Fatal(err)
	}
}
