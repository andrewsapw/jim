# jim

jim is a tool for help you with git.

jim is not a git replacement, it just automates some git workflows that I use frequently.

## Commands

- `jim sw [<branch>]` - Checkout a branch. First, it stashes current branch changes. Then, it checks out the specified branch. Finally, it pops the stash on the target branch (if jim created it before).
- `jim sync` - Sync the current branch with the remote branch.
- `jim ignore [<file>]` - Do not track a file. Under the hood, it runs `git update-index --assume-unchanged [<file>]`
- `jim unignore [<file>]` - Track a file. Under the hood, it runs `git update-index --no-assume-unchanged [<file>]`

If jim is not able to run a command, it will run the "raw" git command instead. For example, `jim status` will just run `git status`.

## Installation

Just run:

```bash
go install -u github.com/andrewsapw/jim@latest
```

### Build from source

To install jim, follow these steps:

1. Clone the repository: `git clone https://github.com/andrewsapw/jim`
2. Change to the project directory: `cd jim`
3. Build the project: `go build -o build/jim`
4. Run the executable: `./build/jim`


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

