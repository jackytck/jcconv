package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completion run:

[Linux]:
  sudo yum install bash-completion -y
  echo "source <(jcconv completion)" >> ~/.bashrc
[Mac]:
  brew install bash-completion@2
  jcconv completion > $(brew --prefix)/etc/bash_completion.d/jcconv

Then restart shell.
`,
	Run: func(cmd *cobra.Command, args []string) {
		must(rootCmd.GenBashCompletion(os.Stdout))
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
