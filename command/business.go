package command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(businessCmd)
	businessCmd.AddCommand(businessRequestsCmd)
}

var businessCmd = &cobra.Command{
	Use:               "business",
	Short:             "Contains various business subcommands",
	PersistentPreRunE: ensureAuth,
	SilenceErrors:     true,
	SilenceUsage:      true,
}

var businessRequestsCmd = &cobra.Command{
	Use:           "requests [status]",
	Short:         "Return business requests",
	Args:          cobra.MinimumNArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          businessRequests,
}

func businessRequests(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/clouds/business/requests/%s?authorization=%s", getUri(), args[0], getToken())
	body, err := GET(url, "business requests")

	if err != nil {
		return err
	}

	var doc map[string]interface{}

	err = json.Unmarshal([]byte(body), &doc)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", body)

	return nil
}
