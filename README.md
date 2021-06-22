# Stop Analyzing API

[![This is a Poppins project](https://raw.githubusercontent.com/bancodobrasil/poppins/master/badge-poppins.svg)](https://github.com/bancodobrasil/poppins)
[![first-timers-only](https://img.shields.io/badge/first--timers--only-friendly-blue.svg?style=flat-square)](https://www.firsttimersonly.com/)

Stop Analyzing is a tool that uses Tinder like interaction to help your customers make up their mind when choosing something that has lot of options, like a product of an e-commerce. [Check this repo for more details](https://github.com/bancodobrasil/stop-analyzing) and [this issue explains how this idea was born](https://github.com/bancodobrasil/stop-analyzing/issues/2). **Stop Analyzing API** has the core functions to make it happen.

## This project was made for first-time contributors and open source beginners

This project follows the [Poppins manifesto guidelines](https://github.com/bancodobrasil/poppins) as part of it's community principles and policies, focusing all the decisions and interactions on providing open source beginners mentorship with real and relevant experiences, respecting each learning pace, background experience, academic formation, questions, suggestions, doubts and opinion.

## Contribute now!

So, let's start contributing! **Open an issue asking for a task to be done by you**. A mentor/maintainer will come and provide a technical overview of the project and what are the possibles ways of contributing to the project. You will discuss the options and a suitable issue will be assigned or created to you.

That's it. Just make yourself at home and good luck!

## Getting Started

### Configure pre-commit code validations

This project includes pre-commit git hooks that make sure the code you commit is valid.
In order to use them, you should rewire the default git hooks folder in your local repository by running (only once) 
```bash
git config core.hooksPath githooks
```


### Generate database using the Prisma migration tool

To generate the database and the go class to handle the connection run the migration tool:

1. Start database and pgadmin using the command

```bash 
docker-compose up -d postgres pgadmin
```

2. At the scripts/prisma folder run the **generate.sh** file using the command 

```bash 
sh generate.sh
```

3. Your database is started and the tables are generated

### Build API using Docker

To build stop-analyzing-api run the below code at the project root folder to run with Docker (suggested):

```
docker-compose build
```

### Run API using Docker

To run stop-analyzing-api run the below code at the project root folder:

```
$ docker-compose up stop-analyzing-api
```

### Build API environment using Docker

To build stop-analyzing-api environment execute the bash script below at the project root folder:

```
$ local_build.sh
```

### Run API using local Go installation

Execute the Go command bellow at the project root folder:

```
$ DATABASE_URL=postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api go run main.go
```

## Awesome list of other Poppins projects for you to go

[![Awesome](https://camo.githubusercontent.com/1997c7e760b163a61aba3a2c98f21be8c524be29/68747470733a2f2f617765736f6d652e72652f62616467652e737667)](https://github.com/sindresorhus/awesome)

- [First Contributions Repository](https://github.com/firstcontributions/first-contributions): Help beginners to contribute to open source projects
- [Contribute to this Project](https://github.com/Syknapse/Contribute-To-This-Project): This is for absolute beginners. If you know how to write and edit an anchor tag <a href="" target=""></a> then you should be able to do it.
- [Contribute to open source](https://github.com/danthareja/contribute-to-open-source):
  Learn the GitHub workflow by contributing code in a fun simulation project
