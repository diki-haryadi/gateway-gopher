package cmd

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "main",
	Short: `
   Database tunnel dialer
    `,
}

func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	rootCmd.AddCommand(GatewayCmd())

	GatewayCmd().PersistentFlags().StringVarP(&Gateway.SSHAuthType, "ssh-auth-type", "a", "", "SSH authentication type (key or userpass)")
	GatewayCmd().PersistentFlags().StringVarP(&Gateway.SSHHost, "ssh-host", "s", "", "SSH jump server hostname")
	GatewayCmd().PersistentFlags().IntVarP(&Gateway.SSHPort, "ssh-port", "p", 22, "SSH port")
	GatewayCmd().PersistentFlags().StringVarP(&Gateway.SSHUser, "ssh-user", "u", "", "SSH username")
	GatewayCmd().PersistentFlags().StringVarP(&Gateway.SSHPassword, "ssh-password", "w", "", "SSH password")
	GatewayCmd().PersistentFlags().StringVarP(&Gateway.SSHKeyPath, "ssh-key", "k", "~/.ssh/id_rsa", "Path to SSH private key")

	GatewayCmd().Flags().StringVarP(&Gateway.DBHost, "db-host", "d", "private-db.internal", "Target example hostname")
	GatewayCmd().Flags().IntVarP(&Gateway.DBPort, "db-port", "b", 5432, "Target example port")

	GatewayCmd().Flags().StringVarP(&Gateway.LocalHost, "local-host", "l", "127.0.0.1", "Local binding address")
	GatewayCmd().Flags().IntVarP(&Gateway.LocalPort, "local-port", "o", 5432, "Local binding port")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(1)
	}
}
