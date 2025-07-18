## sfeir_cli completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(sfeir_cli completion zsh)

To load completions for every new session, execute once:

#### Linux:

	sfeir_cli completion zsh > "${fpath[1]}/_sfeir_cli"

#### macOS:

	sfeir_cli completion zsh > $(brew --prefix)/share/zsh/site-functions/_sfeir_cli

You will need to start a new shell for this setup to take effect.


```
sfeir_cli completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [sfeir_cli completion](sfeir_cli_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 3-Jul-2025
