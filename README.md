<p align="center">
  <img alt="gitlab" src="https://natanfelles.github.io/assets/img_posts/gitlab.png" width="250px" float="center"/>
</p>

<h1 align="center">Welcome to Glabby repository</h1>

> Glabby is a CLI to interact with yout pipeline (Make complience to block jobs in pipeline...)

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

## Show your support

Give a ⭐️ if this project helped you!
