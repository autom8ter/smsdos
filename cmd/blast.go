// Copyright © 2019 Coleman Word
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/autom8ter/smsdos/cats"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// blastCmd represents the blast command
var blastCmd = &cobra.Command{
	Use:   "cats",
	Short: "start an cats sms blast",
	Run: func(cmd *cobra.Command, args []string) {
		send := viper.GetInt("send")
		if send == 0 {
			send = 1
		}
		for x := 0; x < send; x++ {
			cats.Blast().Attack()
		}
	},
}

func init() {
	rootCmd.AddCommand(blastCmd)
}
