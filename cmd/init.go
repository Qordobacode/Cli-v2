package cmd

import (
	"bufio"
	"fmt"
	"github.com/qordobacode/cli-v2/pkg/general/log"
	"github.com/qordobacode/cli-v2/pkg/types"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:         "init",
	Short:       "Init configuration for.ConfigConfig CLI from STDIN or file",
	RunE:        RunInitRoot,
	Example:     "qor init",
	Annotations: map[string]string{"version": APIVersion},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// RunInitRoot function starts config initialization
func RunInitRoot(cmd *cobra.Command, args []string) error {
	fmt.Println("init called")
	fileName := ""
	if len(args) > 0 {
		fileName = args[0]
	}
	var newConfig *types.Config
	var err error
	if fileName != "" {
		newConfig, err = ConfigurationService.ReadConfigInPath(fileName)
		if err != nil {
			return err
		}
	}
	if newConfig == nil {
		newConfig = buildConfigFromStdin()
	}

	ConfigurationService.SaveMainConfig(newConfig)
	return nil
}

func buildConfigFromStdin() *types.Config {
	scanner := bufio.NewScanner(os.Stdin)
	accessToken := readVariable("ACCESS TOKEN: ", "Access token can't be empty\n", scanner)
	organizationID := readIntVariable("ORGANIZATION ID: ", "Organization ID can't be empty\n", scanner)
	projectID := readIntVariable("PROJECT ID: ", "Project ID can't be empty\n", scanner)
	return &types.Config{
		Qordoba: types.QordobaConfig{
			AccessToken:    accessToken,
			ProjectID:      projectID,
			OrganizationID: organizationID,
		},
	}
}

func readVariable(header, errMessage string, scanner *bufio.Scanner) string {
	for {
		fmt.Print(header)
		scanner.Scan()
		text := scanner.Text()
		if text != "" {
			return text
		}
		log.Infof("%s", errMessage)
	}
}

func readIntVariable(header, errMesage string, scanner *bufio.Scanner) int64 {
	for {
		fmt.Print(header)
		scanner.Scan()
		text := scanner.Text()
		if text != "" {
			num, err := strconv.ParseInt(text, 10, 64)
			if err == nil {
				return num
			}
		}
		fmt.Printf(errMesage)
	}
}
