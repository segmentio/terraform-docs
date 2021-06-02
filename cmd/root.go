/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/cmd/asciidoc"
	"github.com/terraform-docs/terraform-docs/cmd/completion"
	"github.com/terraform-docs/terraform-docs/cmd/json"
	"github.com/terraform-docs/terraform-docs/cmd/markdown"
	"github.com/terraform-docs/terraform-docs/cmd/pretty"
	"github.com/terraform-docs/terraform-docs/cmd/tfvars"
	"github.com/terraform-docs/terraform-docs/cmd/toml"
	versioncmd "github.com/terraform-docs/terraform-docs/cmd/version"
	"github.com/terraform-docs/terraform-docs/cmd/xml"
	"github.com/terraform-docs/terraform-docs/cmd/yaml"
	"github.com/terraform-docs/terraform-docs/internal/cli"
	"github.com/terraform-docs/terraform-docs/internal/version"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := NewCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}
	return nil
}

// NewCommand returns a new cobra.Command for 'root' command
func NewCommand() *cobra.Command {
	config := cli.DefaultConfig()
	cmd := &cobra.Command{
		Args:          cobra.MaximumNArgs(1),
		Use:           "terraform-docs [PATH]",
		Short:         "A utility to generate documentation from Terraform modules in various output formats",
		Long:          "A utility to generate documentation from Terraform modules in various output formats",
		Version:       version.Full(),
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations:   cli.Annotations("root"),
		PreRunE:       cli.PreRunEFunc(config),
		RunE:          cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().StringVarP(&config.File, "config", "c", ".terraform-docs.yml", "config file name")

	cmd.PersistentFlags().StringSliceVar(&config.Sections.Show, "show", []string{}, "show section ["+cli.AllSections+"]")
	cmd.PersistentFlags().StringSliceVar(&config.Sections.Hide, "hide", []string{}, "hide section ["+cli.AllSections+"]")

	cmd.PersistentFlags().StringVar(&config.Output.File, "output-file", "", "File path to insert output into (default \"\")")
	cmd.PersistentFlags().StringVar(&config.Output.Mode, "output-mode", "inject", "Output to file method ["+cli.OutputModes+"]")
	cmd.PersistentFlags().StringVar(&config.Output.Template, "output-template", cli.OutputTemplate, "Output template")
	cmd.PersistentFlags().BoolVar(&config.Output.Check, "output-check", false, "Check if content of output file is up to date (default false)")

	cmd.PersistentFlags().BoolVar(&config.Sort.Enabled, "sort", true, "sort items")
	cmd.PersistentFlags().StringVar(&config.Sort.By, "sort-by", "name", "sort items by criteria ["+cli.SortTypes+"]")

	// deprecated flags ==>
	cmd.PersistentFlags().BoolVar(new(bool), "show-all", true, "show all sections")
	cmd.PersistentFlags().BoolVar(new(bool), "hide-all", false, "hide all sections (default false)")
	cmd.PersistentFlags().MarkDeprecated("show-all", "more information: https://terraform-docs.io/user-guide/how-to/#visibility-of-sections\n\n") //nolint:errcheck,gosec
	cmd.PersistentFlags().MarkDeprecated("hide-all", "more information: https://terraform-docs.io/user-guide/how-to/#visibility-of-sections\n\n") //nolint:errcheck,gosec

	cmd.PersistentFlags().BoolVar(&config.Sort.Criteria.Required, "sort-by-required", false, "sort items by name and print required ones first (default false)")
	cmd.PersistentFlags().BoolVar(&config.Sort.Criteria.Type, "sort-by-type", false, "sort items by type of them (default false)")
	cmd.PersistentFlags().MarkDeprecated("sort-by-required", "use '--sort-by required' instead\n\n") //nolint:errcheck,gosec
	cmd.PersistentFlags().MarkDeprecated("sort-by-type", "use '--sort-by type' instead\n\n")         //nolint:errcheck,gosec
	// <==

	cmd.PersistentFlags().StringVar(&config.HeaderFrom, "header-from", "main.tf", "relative path of a file to read header from")
	cmd.PersistentFlags().StringVar(&config.FooterFrom, "footer-from", "", "relative path of a file to read footer from (default \"\")")

	cmd.PersistentFlags().BoolVar(&config.OutputValues.Enabled, "output-values", false, "inject output values into outputs (default false)")
	cmd.PersistentFlags().StringVar(&config.OutputValues.From, "output-values-from", "", "inject output values from file into outputs (default \"\")")

	// formatter subcommands
	cmd.AddCommand(asciidoc.NewCommand(config))
	cmd.AddCommand(json.NewCommand(config))
	cmd.AddCommand(markdown.NewCommand(config))
	cmd.AddCommand(pretty.NewCommand(config))
	cmd.AddCommand(tfvars.NewCommand(config))
	cmd.AddCommand(toml.NewCommand(config))
	cmd.AddCommand(xml.NewCommand(config))
	cmd.AddCommand(yaml.NewCommand(config))

	// other subcommands
	cmd.AddCommand(completion.NewCommand())
	cmd.AddCommand(versioncmd.NewCommand())

	return cmd
}
