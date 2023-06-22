package cmd

import (
	"fmt"
	"os"
	"io"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use: "find",
	Short: "Find local private key(s)",
	Long: `Find local private key(s). Use flags to set custom search locations`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := os.UserHomeDir()

		// scan results
		var(
			gnupg string = ""
			private string = ""
			priv_ex bool = false
			keys string = ""
		)

		if _, err := os.Stat(usr + "/.gnupg"); !os.IsNotExist(err) {
			// .gnupg exists
			gnupg = (usr + "/.gnupg exists ... ✅")
		} else {
			gnupg = (usr + "/.gnupg exists ... ❌")
		}

		if _, err := os.Stat(usr + "/.gnupg/private-keys-v1.d"); !os.IsNotExist(err) {
			// private-keys-v1.d exists
			private = (usr + "/.gnupg/private-keys-v1.d exists ... ✅")
			priv_ex = true
		} else {
			private = (usr + "/.gnupg/private-keys-v1.d exists ... ❌")
		}

		if (priv_ex) {
			f, err := os.Open(usr + "/.gnupg/private-keys-v1.d")
			defer f.Close()
			_, err = f.Readdirnames(1)
			if err == io.EOF {
				keys = ("private key(s) present ... ❌")
			} else {
				keys = ("private key(s) present ... ✅")
			}
		} else {
			fmt.Println("Cannot open private-keys-v1.d ... Does not exist")
		}

		fmt.Println(gnupg)
		fmt.Println(private)
		fmt.Println(keys)
	},
}