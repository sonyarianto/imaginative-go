# Imaginative Go

[![Build Status](https://travis-ci.org/sonyarianto/imaginative-go.svg?branch=master)](https://travis-ci.org/sonyarianto/imaginative-go) [![Go Report Card](https://goreportcard.com/badge/github.com/sonyarianto/imaginative-go)](https://goreportcard.com/report/github.com/sonyarianto/imaginative-go) [![Maintainability](https://api.codeclimate.com/v1/badges/e8d5f5483ea4c87df280/maintainability)](https://codeclimate.com/github/sonyarianto/imaginative-go/maintainability)

A beautiful website that created with Go. It contains many Go working samples code that useful for web and non-web application. It mainly demonstrate what Go can achieve. You can contribute!

In this imaginative (but real) project, we implement our knowledge during learning Go and we want to share with the community. We don't use any Go framework and forgive us if the code still not efficient or optimal, since we are still learning too in Go language. Any suggestions are welcomed.

## Why Imaginative Go?
- Ideas about achieve something with Go language
- Plenty of working code samples (awwww, currently still not plenty, relaxxx mann, we will add more regularly)
- Easy to run in your local machine (with Docker)
- You can contribute

## Requirements (Optional)
- Go 1.11.x or later
- Docker Engine (version 18.09.x or later)
- Docker Compose (version 1.23.x or later)

Docker is useful if you want to run this website with all its content on your local machine.

For those who need documentation of Docker installation, please refer to [Docker CE](https://store.docker.com/search?type=edition&offering=community) and [Docker Compose](https://docs.docker.com/compose/install/).

> **NOTE:** `sudo` used throughout this doc, since mainly we use Linux/MacOS during the development. We test running on Windows 10 as well with Docker for Windows and Docker Toolbox.

> **NOTE:** For user that using Windows 10 Home that run with Docker Toolbox, I think you should modify IP on `docker-compose.yml` from 127.0.0.1 to your Docker Machine IP. Docker Machine IP can be known by typing `docker-machine ip`.

## Usage
#### Linux/MacOS/Windows
```
git clone https://github.com/sonyarianto/imaginative-go.git
cd imaginative-go
sudo docker-compose up -d
```

After that, go to your browser and type
```
http://localhost:9899
```

> **Note 1**
> 
> To clean containers (stop and remove), remove volume, clean network (remove), pull, build (with no cache) and bring up all containers you can type like below
> ```
> sudo docker-compose down && sudo docker volume rm imaginative-go_volume-mongodb-imaginative-go && sudo docker-compose pull && sudo docker-compose build && sudo docker-compose up -d --build --force-recreate
> ```
> Above command will make sure you will create and run fresh all containers needed to run Imaginative Go web project. This is usualy useful after you are doing `git pull` on Imaginative Go repository.
> 
> Above command will error if volume `imaginative-go_volume-mongodb-imaginative-go` doesn't exists. You can remove the delete volume part if you encounter that error

> **Note 2**
> 
> MongoDB expose random port to host machine. You can see it by typing this after all containers are running.
> ```
> sudo docker ps -f "name=mongodb-imaginative-go"
> ```
> Sample output is like below
> ```
> CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS                                NAMES
> f0c2c2166487        imaginative-go_mongodb   "docker-entrypoint.s…"   About an hour ago   Up About an hour    0.0.0.0:32782->27017/tcp             mongodb-imaginative-go
> ```

## Docker Images Used
- `mongodb:latest` (see `docker-compose.yml` for default credential, port exposed to host is using random port, see on `docker ps` command)

## Contributors
- [Sony Arianto Kurniawan](https://github.com/sonyarianto) - sony@sony-ak.com - original author, project maintainer
- [Prasetyama Hidayat](https://github.com/prasetyama) - prasetyama@gmail.com

Do you want to contribute? Just fork this repository and contribute anything you can (e.g. fix typo, bug fix, add new sample etc.)

## Community/Contributing
Imaginative Go maintains a mailing list, [Imaginative Go][imaginative-go], where you should feel
welcome to ask questions about the project (no matter how simple!) or to talk about Imaginative Go more
generally. Imaginative Go's author (Sony Arianto Kurniawan) also loves to hear from users directly
at his personal email address, which is available on his GitHub profile page.

Contributions to Imaginative Go are welcome.

All interactions in the Imaginative Go community will be held to the high standard of the
broader Go community's [Code of Conduct][conduct].

[imaginative-go]: https://groups.google.com/forum/#!forum/imaginative-go
[conduct]: https://golang.org/conduct

## License
This project is licensed under the MIT License.

License can be found [here](https://github.com/sonyarianto/imaginative-go/blob/master/LICENSE).
