# gVisor Sandbox CLI

`sandbox-cli` is a command-line tool designed for in-time configuration of gVisor, an open-source container runtime sandbox. This CLI tool allows users to perform various operations related to gVisor's callback functions and state management.

## Usage

```
sandbox-cli <Command> [-h|--help] [-a|--address "<value>"]
```

## Hint: you can set address in environment
`export CLI_ADDRESS=localhost:8080`

## Commands

### 1. `change`

Change callbacks.

### 2. `man`

Show man (manual) for hooks. This command provides detailed information and documentation about the available hooks that can be used in gVisor's JavaScript engine.

### 3. `state`

Change state. This command allows users to modify the state of gVisor during runtime.

### 4. `get`

Get current callbacks. Use this command to retrieve the currently active callback functions in gVisor.

### 5. `delete`

Unregister callbacks. This command lets users unregister or remove specific callback functions from gVisor.

## Arguments

The following arguments are available for the CLI:

- `-h`, `--help`: Print help information. Use this argument to display the CLI's usage and available commands.
- `-a`, `--address "<value>"`: Socket address. This argument allows users to specify the socket address for communication.

## Example Usage

### Get Current Callbacks:

```
sandbox-cli get
```

This command retrieves and displays the currently active callback functions.

### Change State:

```
sandbox-cli state -a "127.0.0.1:8080"
```

Using this command, users can change the state of gVisor by providing the desired socket address (e.g., "127.0.0.1:8080").

### Unregister Callbacks:

```
sandbox-cli delete
```

This command allows users to unregister specific callback functions from gVisor.

## Additional Information

For more details about gVisor and the available hooks, use the following command:

```
sandbox-cli man
```

This will provide comprehensive information about the hooks and their usage in the JavaScript engine integrated into gVisor.
 provided information is a fictional representation of a CLI tool for gVisor. The actual CLI tool and its usage may vary based on the specific implementation and requirements.