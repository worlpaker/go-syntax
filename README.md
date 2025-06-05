# Go Syntax

Rich Syntax Highlighting for [Go](https://go.dev/) language

With this extension, your favorite theme will be able to color your code better.

> **NOTE:** Go Syntax became the default [Go grammar](https://github.com/microsoft/vscode/blob/main/extensions/go/syntaxes/go.tmLanguage.json) in the official VS Code in v1.86 ([release note](https://code.visualstudio.com/updates/v1_86#_new-go-grammar)).

## About

Go Syntax improves your coding experience by providing client-side syntax highlighting based on [TextMate rules](https://macromates.com/manual/en/language_grammars). This means you can enjoy advanced code coloring without relying on the Language Server Protocol (LSP).

It is compatible with various platforms, including web-based code editors like [vscode.dev](https://vscode.dev) and [github.dev](https://github.dev).

![dark+_theme](examples/dark+_after.png)

**⚠️ Warning:** [Gopls' `ui.semanticTokens` setting](https://github.com/golang/vscode-go/wiki/settings#uisemantictokens) is disabled by default, and enabling semantic highlighting might override the tokens.

## Contributing

Yes, please! Feel free to contribute. Check out [CONTRIBUTING.md](CONTRIBUTING.md).

## Credits

This extension enhances [better-go-syntax](https://github.com/jeff-hykin/better-go-syntax) with new features.

Semantic tokens tested thanks to the [vscode-tmgrammar-test](https://github.com/PanAeon/vscode-tmgrammar-test).

## License

[MIT](https://github.com/worlpaker/go-syntax/blob/master/LICENSE)
