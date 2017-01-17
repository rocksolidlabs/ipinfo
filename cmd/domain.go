// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"net"
	"reflect"

	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Enter a domain name to get its IP info",
	Long:  "Enter a domain name to get its IP info",
	Run: func(cmd *cobra.Command, args []string) {
		// curl ipinfo.io
		fmt.Printf("Getting IP info for domain: %s\n", args[0])
		fmt.Println()
		ips, err := net.LookupIP(args[0])
		if err != nil {
			fmt.Printf("Error looking up ip for hostname: %s\n", args[0])
		}

		for _, IP := range ips {
			ip := IP.String()
			info := IPInfo{}
			err := getJson("http://ipinfo.io/"+ip+"/json", &info)
			if err != nil {
				fmt.Printf("Error getting ipinfo for ip: %s\n", ip)
			}
			s := reflect.ValueOf(&info).Elem()
			typeOfT := s.Type()

			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				fmt.Printf("%s = %v\n", typeOfT.Field(i).Name, f.Interface())
			}
			fmt.Println()
		}

	},
}

func init() {
	RootCmd.AddCommand(domainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
