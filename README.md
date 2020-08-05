<p align="center">
  <img alt="gitlab" src="https://natanfelles.github.io/assets/img_posts/gitlab.png" width="250px" float="center"/>
</p>

<h1 align="center">Welcome to Glabby repository</h1>

<p align="center">
  <a href="https://github.com/lpmatos/glabby">
    <img alt="Open Source" src="https://badges.frapsoft.com/os/v1/open-source.svg?v=102">
  </a>

  <a href="https://github.com/lpmatos/glabby/graphs/contributors">
    <img alt="GitHub Contributors" src="https://img.shields.io/github/contributors/lpmatos/glabby">
  </a>

  <a href="https://github.com/lpmatos/glabby">
    <img alt="GitHub Language Count" src="https://img.shields.io/github/languages/count/lpmatos/glabby">
  </a>

  <a href="https://github.com/lpmatos/glabby">
    <img alt="GitHub Top Language" src="https://img.shields.io/github/languages/top/lpmatos/glabby">
  </a>

  <a href="https://github.com/lpmatos/glabby/stargazers">
    <img alt="GitHub Stars" src="https://img.shields.io/github/stars/lpmatos/glabby?style=social">
  </a>

  <a href="https://github.com/lpmatos/glabby/commits/master">
    <img alt="GitHub Last Commit" src="https://img.shields.io/github/last-commit/lpmatos/glabby">
  </a>

  <a href="https://github.com/lpmatos/glabby">
    <img alt="Repository Size" src="https://img.shields.io/github/repo-size/lpmatos/glabby">
  </a>

  <a href="https://github.com/lpmatos/glabby/issues">
    <img alt="Repository Issues" src="https://img.shields.io/github/issues/lpmatos/glabby">
  </a>

  <a href="https://github.com/lpmatos/glabby/blob/master/LICENSE">
    <img alt="MIT License" src="https://img.shields.io/github/license/lpmatos/glabby">
  </a>
</p>

> Glabby is a CLI to interact with your pipeline (Make complience to block jobs in pipeline...)

## Menu

<p align="left">
  <a href="#pre-requisites">Pre-Requisites</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#how-to-contribute">How to contribute</a>
</p>

## Getting Started

If you want use this repository you need to make a **git clone**:

```bash
git clone --depth 1 https://github.com/lpmatos/glabby.git -b master
```

This will give access on your **local machine**.

## Pre-Requisites

To this project you yeed:

### Requirement

* Git.
* Golang 1.13.
* Docker and Docker Compose.

### Optional 

* NPM | Yarn (package tool)
* Install Packages
  * Husky
  * Commitlint
  * Commitizen
  * Standard-Version

## Description

This project builds a simple CLI to interact with your GitLab pipeline.

## Install

```sh
go get github.com/lpmatos/glabby
```

## Usage

#### Global

```text
   ____   _           _       _
  / ___| | |   __ _  | |__   | |__    _   _
 | |  _  | |  / _` | | '_ \  | '_ \  | | | |
 | |_| | | | | (_| | | |_) | | |_) | | |_| |
  \____| |_|  \__,_| |_.__/  |_.__/   \__, |
                                      |___/
Description:

Glabby is a CLI library to interact with your pipeline and make/analyze some stuffs.

Usage:
  glabby [command]

Available Commands:
  compliance  Run compliance in pipeline
  help        Help about any command
  version     Version outputs the version of CLI

Flags:
  -h, --help      help for glabby
  -s, --silence   enable silence mod without logs in stdout terminal.

Use "glabby [command] --help" for more information about a command.
```

#### Complience

```text
   ____   _           _       _
  / ___| | |   __ _  | |__   | |__    _   _
 | |  _  | |  / _` | | '_ \  | '_ \  | | | |
 | |_| | | | | (_| | | |_) | | |_) | | |_| |
  \____| |_|  \__,_| |_.__/  |_.__/   \__, |
                                      |___/
Description:

Compliance to your GitLab pipeline.

Usage:
  glabby compliance [command]

Aliases:
  compliance, c

Available Commands:
  sonar       Run sonar compliance
  speedio     Run speedio compliance

Flags:
  -h, --help   help for compliance

Global Flags:
  -s, --silence   enable silence mod without logs in stdout terminal.

Use "glabby compliance [command] --help" for more information about a command.
```

## Commit Lint

We all did bad commit messages. Lucky us, Conventional Commits specification exists, and with it a set of powerful tools to help us.

To enforce a standard every time we make a commit, we have husky and commitlint. Husky listen to git hooks, and we will use it to trigger the commitlint when we type a commit message.

Commitizen is a package that makes it easier to create commit messages following the previous specification.

* husky
* commitlint
* commitizen

<strong>Requirements:</strong>

> OBS: Required .git folder in project
* .git in folder
* Node
* yarn | npm

<strong>Install by default [package.json](package.json):</strong>

```cmd
yarn install
```

<strong>Manual Installment:</strong>

```cmd
yarn init -y

yarn add @commitlint/config-conventional @commitlint/cli husky commitizen -D

echo module.exports = {extends: ['@commitlint/config-conventional']} > commitlint.config.js
```

add configuration in package.json:

```json
"husky": {
    "hooks": {
        "commit-msg": "commitlint -E HUSKY_GIT_PARAMS"
    }
},
"config": {
    "commitizen": {
        "path": "./node_modules/cz-conventional-changelog"
    }
}
```

<strong>Use:</strong>

with dependencies already installed, commits that do not follow the semmantic commit rules will be automatically blocked in the development environment

```cmd
C:\>  git add .
C:\>  git commit -m "commit"


husky > commit-msg (node v12.14.0)
⧗   input: commit
✖   subject may not be empty [subject-empty]
✖   type may not be empty [type-empty]

✖   found 2 problems, 0 warnings
ⓘ   Get help: https://github.com/conventional-changelog/commitlint/#what-is-commitlint

husky > commit-msg hook failed (add --no-verify to bypass)
```

using the commitzen, previously installed, an auxiliary service will be available to build the commits

```cmd
C:\>  git add .
C:\>  npm run commit


cz-cli@4.0.3, cz-conventional-changelog@3.2.0

? Select the type of change that you're committing: (Use arrow keys)
> feat:     A new feature
  fix:      A bug fix
  docs:     Documentation only changes
  style:    Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  refactor: A code change that neither fixes a bug nor adds a feature
  perf:     A code change that improves performance
  test:     Adding missing tests or correcting existing tests
  build:    Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
  ci:       Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs) 
  chore:    Other changes that don't modify src or test files
  revert:   Reverts a previous commit
```

## Standard version

Once you have conventional commits, you can make use of them with the next tool. Standard-version will usually do the following things:

* updates semantic version according to the scope of changes described in the commits
* updates CHANGELOG.md file with the new version and list of changes
* commits both changes and tags them with the new version

## Dockle

Container Image Linter for Security, Helping build the Best-Practice Docker Image, Easy to start

`Dockle` helps you:

1. Build [Best Practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) Docker images
2. Build secure Docker images
    - Checkpoints includes [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/)

BasicAuth server needs `DOCKLE_USERNAME` and `DOCKLE_PASSWORD`.

```bash
export DOCKLE_USERNAME={USERNAME}
export DOCKLE_PASSWORD={PASSWORD}

# if you'd like to use 80 port, use NonSSL
export DOCKLE_NON_SSL=true
```

Run command:

```bash
dockle --username ${DOCKLE_USERNAME} --password ${DOCKLE_PASSWORD} -o results.json -f json -c 1 -l warn -d registry
```

For more information access [Dockle](https://github.com/goodwithtech/dockle).

## How to contribute

1. Make a **Fork**.
2. Follow the project organization.
3. Add the file to the appropriate level folder - If the folder does not exist, create according to the standard.
4. Make the **Commit**.
5. Open a **Pull Request**.
6. Wait for your pull request to be accepted.. 🚀

Remember: There is no bad code, there are different views/versions of solving the same problem. 😊

## Add to git and push

You must send the project to your GitHub after the modifications

```bash
git add -f .
git commit -m "Added - Fixing somethings"
git push origin master
```

## Versioning

- [CHANGELOG](CHANGELOG.md)

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Author

👤 **Lucca Pessoa**

Hey!! If you like this project or if you find some bugs feel free to contact me in my channels:

> * Email: luccapsm@gmail.com
> * Website: https://github.com/lpmatos
> * Github: [@lpmatos](https://github.com/lpmatos)
> * LinkedIn: [@luccapessoa](https://www.linkedin.com/in/lucca-pessoa-4abb71138/)

## Show your support

Give a ⭐️ if this project helped you!

## Project Status

* ✔️ Finish

---

<p align="center">Make with ❤️ by <strong>Lucca Pessoa :wave:</p>
