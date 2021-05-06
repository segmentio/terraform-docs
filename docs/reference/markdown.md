---
title: "markdown"
description: "Generate Markdown of inputs and outputs."
menu:
  docs:
    parent: "terraform-docs"
weight: 955
toc: true
---

## Synopsis

Generate Markdown of inputs and outputs.

```console
terraform-docs markdown [PATH] [flags]
```

## Options

```console
      --anchor       create anchor links (default true)
      --default      show Default column or section (default true)
      --escape       escape special characters (default true)
  -h, --help         help for markdown
      --indent int   indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
      --required     show Required column or section (default true)
      --sensitive    show Sensitive column or section (default true)
      --type         show Type column or section (default true)
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

- [terraform-docs markdown document]({{< ref "markdown-document" >}})
- [terraform-docs markdown table]({{< ref "markdown-table" >}})
