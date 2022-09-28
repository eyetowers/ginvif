package main

import (
	"os"

	"github.com/eyetowers/gonvif/pkg/gonvif"
	"github.com/spf13/cobra"

	"github.com/eyetowers/ginvif/pkg/ginvif"
)

var (
	bind     string
	url      string
	username string
	password string
	verbose  bool
)

var cmd = &cobra.Command{
	Use:   "ginvif",
	Short: "Onvif JSON API Adapter",
	Long:  "HTTP Server translating JSON requests into Onvif (onvif.org) WS/SOAP protocol.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := gonvif.New(url, username, password, verbose)
		if err != nil {
			return err
		}
		return ginvif.Serve(bind, client)
	},
}

func init() {
	cmd.Flags().StringVarP(&bind, "bind", "b", "0.0.0.0:8081", "Address and port to serve requests on.")
	cmd.Flags().StringVarP(&url, "url", "a", "", "Base URL of the Onvif device.")
	cmd.Flags().StringVarP(&username, "username", "u", "", "Username for authentication with the Onvif device.")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Password for authentication with the Onvif device.")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print sent and received requests.")
	mustMarkFlagRequired(cmd, "url")
	mustMarkFlagRequired(cmd, "username")
	mustMarkFlagRequired(cmd, "password")
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func mustMarkFlagRequired(cmd *cobra.Command, flag string) {
	err := cmd.MarkFlagRequired(flag)
	if err != nil {
		panic(err)
	}
}
