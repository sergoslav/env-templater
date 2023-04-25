# env-templater
### Tool to make .env file use template.

This program takes four arguments from the command line:

```
  -c string
        templater config .yaml file (default ".env-templater.yaml")
  -d string
        the name of the new file that will be created or overwritten. (default ".env")
  -e string
        environment name (default "dev")
  -s string
        the name of the source .env file. (default ".env.template")
```

The program opens the source file and creates a new file. And do one of action:

- Removes line if it has prefix with one of env name but not currently used.
- Remove prefix and insert this line if prefix currently used.
- Does replace


## Go Install
```bash
go install github.com/sergoslav/env-templater@latest
```

## Usages Example

[yaml config](./example/env-templater.yaml)

[env template](./example/.env.template)


Run:
```bash
$ env-templater -s example/.env.template -d example/.env.dev -e dev -c example/env-templater.yaml
```
Result: [.env.dev](./example/.env.dev)


Run:
```bash
$ env-templater -s example/.env.template -d example/.env.debug -e debug -c example/env-templater.yaml
```
Result: [.env.dev](./example/.env.debug)