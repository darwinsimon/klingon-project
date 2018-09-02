[![Build Status](https://travis-ci.org/darwinsimon/klingon-project.svg?branch=master)](https://travis-ci.org/darwinsimon/klingon-project) [![Coverage Status](https://coveralls.io/repos/github/darwinsimon/klingon-project/badge.svg?branch=master)](https://coveralls.io/github/darwinsimon/klingon-project?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/darwinsimon/klingon-project)](https://goreportcard.com/report/github.com/darwinsimon/klingon-project)

# Klingon Project

This application will translate any character written in English to Klingon. If the give name is a valid Star Trek character, it will print any supported information related to the character.
The Klingon translation will use unicode as representative. Character information is provided by http://stapi.co

Valid information will be stored in `char.txt` to reduce API calls for the same character.

Available information:
- Species
- Gender (hidden)
- Year of Birth (hidden)

To show hidden information, add `showas` in the `Character` struct. [Go to file](stapi/types.stapi.go) 

## How To Use

### Build
- Install Go. Downloads and instructions are available [here](https://golang.org/dl/).
- For Windows, run `go get` and `go build` in repository main folder
- For Unix, run `make`

### Run
Run the application with character's name as parameter
```bash
./klingon-project Uhura

KLINGON PROJECT
---------------

Processing...

Input          : Uhura
Klingon Name   : 0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0
Species        : Human
```

### Flags
Add `-v` for verbose logging

## Restrictions
- Translation will fail if the name consists any unavailable characters (c, g, k, x, z, etc.)
- All characters are treated as case-insensitive, except for Q
- If Star Trek API returns more than 10 results, species will be treated as inconclusive.
- Space will use unicode `0x0020`

### Klingon Text
![Klingon Text](./doc/klingon-text.png)

### Klingon Unicode
![Klingon Unicode](./doc/unicode.png)