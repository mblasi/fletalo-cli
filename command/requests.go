package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var since string

func init() {
	rootCmd.AddCommand(requestsCmd)
	requestsCmd.AddCommand(lastCmd)
	requestsCmd.AddCommand(listCmd)
	requestsCmd.AddCommand(showRequestCmd)
	requestsCmd.PersistentFlags().StringVarP(&impersonalize, "impersonalize", "i", "me", "Run command impersonalized as other user (nickname)")
	lastCmd.PersistentFlags().StringVarP(&since, "since", "s", "1d", "Specifies timeframe to search last requests")
}

var requestsCmd = &cobra.Command{
	Use:               "requests",
	Short:             "Contains various requests subcommands",
	PersistentPreRunE: ensureAuth,
	SilenceErrors:     true,
	SilenceUsage:      true,
}

var lastCmd = &cobra.Command{
	Use:           "last",
	Short:         "Return last requests",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          last,
}

var listCmd = &cobra.Command{
	Use:           "list",
	Short:         "Return all requests",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          list,
}

var showRequestCmd = &cobra.Command{
	Use:           "show",
	Short:         "Show request details",
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.MinimumNArgs(1),
	RunE:          showRequest,
}

func last(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/requests/last?since=%s&authorization=%s", getUri(), since, getToken())
	body, err := GET(url, "last requests")

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", body)

	return nil
}

func list(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/requests?authorization=%s", getUri(), getToken())
	body, err := GET(url, "requests")

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", body)

	return nil
}

func showRequest(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/request/%s?authorization=%s", getUri(), args[0], getToken())
	body, err := GET(url, "requests")

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", body)

	return nil
}
