/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"

	"github.com/gitpod-io/gitpod/previewctl/pkg/preview"
	"github.com/spf13/cobra"
)

func getNameCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get-name",
		Short: "Returns the name of the preview for the corresponding branch.",
		Run: func(cmd *cobra.Command, args []string) {
			p := preview.New(branch)

			fmt.Println(p.GetPreviewName())
		},
	}

	return cmd
}
