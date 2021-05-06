---
title: "tfvars"
description: "Generate terraform.tfvars of inputs."
menu:
  docs:
    parent: "terraform-docs"
weight: 959
toc: true
---

## Synopsis

Generate terraform.tfvars of inputs.

## Options

```console
  -h, --help   help for tfvars
```

## Inherited Options

```console
  -c, --config string               config file name (default ".terraform-docs.yml")
      --footer-from string          relative path of a file to read footer from (default "")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --output-check                Check if content of output file is up to date (default false)
      --output-file string          File path to insert output into (default "")
      --output-mode string          Output to file method [inject, replace] (default "inject")
      --output-template string      Output template (default "<!-- BEGIN_TF_DOCS -->\n{{ .Content }}\n<!-- END_TF_DOCS -->")
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --show strings                show section [data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --sort                        sort items (default true)
      --sort-by string              sort items by criteria [name, required, type] (default "name")
```

## Subcommands

- [terraform-docs tfvars hcl]({{< ref "tfvars-hcl" >}})
- [terraform-docs tfvars json]({{< ref "tfvars-json" >}})
