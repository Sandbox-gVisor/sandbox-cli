# gVisor Sandbox CLI

`sandbox-cli` is a command-line tool designed for in-time configuration of gVisor Sandbox, an open-source container runtime sandbox. This CLI tool allows users to perform various operations related to gVisor's callback functions and state management.

## Usage

```
sandbox-cli <Command> [-h|--help] [-a|--address "<value>"]
```

## Arguments

The following arguments are available for the CLI and every command:

- `-h`, `--help`: Print help information. Use this argument to display the CLI's usage and available commands.
- `-a`, `--address "<value>"`: Socket address. This argument allows users to specify the socket address for communication.

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

## Example Usage

### Get Current Callbacks with verbose output:

```
sandbox-cli get --verbose
```

This command retrieves and displays the currently active callback functions.

### Change State:

```
sandbox-cli state -a "127.0.0.1:8080" -c change_state.js
```

Using this command, users can change the state of gVisor by providing the desired socket address (e.g., "127.0.0.1:8080").

### Unregister Callback after syscall 59(exec):

```
sandbox-cli delete -s 59 -t after
```

This command allows users to unregister specific callback functions from gVisor.

## Additional Information

For more details about gVisor and the available hooks, use the following command:

```
sandbox-cli man
```

This will provide comprehensive information about the hooks and their usage in the JavaScript engine integrated into gVisor.
