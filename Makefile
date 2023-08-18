# Ensure 'vscode-tmgrammar-test' is installed before running these tests.

tests:
	vscode-tmgrammar-test ./test/semantic_tokens.go
snap:
	vscode-tmgrammar-snap ./test/semantic_tokens.go
