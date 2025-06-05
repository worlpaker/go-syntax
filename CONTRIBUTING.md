# Contributing Guide

Thanks for your interest in contributing to Go Syntax!

Any kind of PR (e.g., tests, performance improvements, bug fixes, features, typos) is welcome!

We include a simple scenario demonstrating each change in [semantic_tokens.go](test/semantic_tokens.go). We take advantage of Git version control to test changes. Eventually, running tests will generate snapshot files, and we will be able to see the differences. This helps us to ensure that we are not breaking anything in the existing sample code.

## Test

Please add a sample of Go code related to your PR in [semantic_tokens.go](test/semantic_tokens.go). Refer to the [Makefile](Makefile) for detailed test commands.

You need to install [vscode-tmgrammar-test](https://www.npmjs.com/package/vscode-tmgrammar-test) before running the tests.

Install it globally:

```sh
npm i -g vscode-tmgrammar-test
```

In the go-syntax directory, run the following tests:

1. Stress Test

    ```sh
    make stress
    ```

2. Semantic Tokens Test

    ```sh
    make ready
    ```

> Many changes may occur in the tests depending on your PR, which is fine.
