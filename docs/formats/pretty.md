# Pretty

Generate a colorized pretty of inputs and outputs.

## Usage

```text
Usage:
  terraform-docs pretty [PATH] [flags]

Flags:
  -h, --help       help for pretty
      --no-color   do not colorize printed result

Global Flags:
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

## Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs pretty --no-color ./examples/
```

generates the following output:

    Usage:

    Example of 'foo_bar' module in `foo_bar.tf`.

    - list item 1
    - list item 2

    Even inline **formatting** in _here_ is possible.
    and some [link](https://domain.com/)

    * list item 3
    * list item 4

    ```
    module "foo_bar" {
        source = "github.com/foo/bar"

        id   = "1234567890"
        name = "baz"

        zones = ["us-east-1", "us-west-1"]

        tags = {
            Name         = "baz"
            Created-By   = "first.last@email.com"
            Date-Created = "20180101"
        }
    }
    ```

    Here is some trailing text after code block,
    followed by another line of text.

    | Name | Description     |
    |------|-----------------|
    | Foo  | Foo description |
    | Bar  | Bar description |



    provider.aws (>= 2.15.0)

    provider.aws.ident (>= 2.15.0)

    provider.null

    provider.tls



    input.input-with-code-block ([
    "name rack:location"
    ])
    This is a complicated one. We need a newline.
    And an example in a code block
    ```
    default     = [
    "machine rack01:neptune"
    ]
    ```

    input.input-with-pipe ("v1")
    It includes v1 | v2 | v3

    input.input_with_underscores (required)
    A variable with underscores.

    input.list-1 ([
    "a",
    "b",
    "c"
    ])
    It's list number one.

    input.list-2 (required)
    It's list number two.

    input.list-3 ([])
    n/a

    input.long_type ({
    "bar": {
        "bar": "bar",
        "foo": "bar"
    },
    "buzz": [
        "fizz",
        "buzz"
    ],
    "fizz": [],
    "foo": {
        "bar": "foo",
        "foo": "foo"
    },
    "name": "hello"
    })
    This description is itself markdown.

    It spans over multiple lines.

    input.map-1 ({
    "a": 1,
    "b": 2,
    "c": 3
    })
    It's map number one.

    input.map-2 (required)
    It's map number two.

    input.map-3 ({})
    n/a

    input.no-escape-default-value ("VALUE_WITH_UNDERSCORE")
    The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.

    input.string-1 ("bar")
    It's string number one.

    input.string-2 (required)
    It's string number two.

    input.string-3 ("")
    n/a

    input.unquoted (required)
    n/a

    input.with-url ("")
    The description contains url. https://www.domain.com/foo/bar_baz.html



    output.output-0.12
    terraform 0.12 only

    output.output-1
    It's output number one.

    output.output-2
    It's output number two.

    output.unquoted
    It's unquoted output.
