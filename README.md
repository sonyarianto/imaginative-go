# Imaginative Go

[![Build Status](https://travis-ci.org/sonyarianto/imaginative-go.svg?branch=master)](https://travis-ci.org/sonyarianto/imaginative-go)

Imaginative Go is a website that created with Go. It contains many Go working samples code that useful for web and non-web application. It mainly demonstrate what Go can achieve. You can contribute!

> I believe that the best way to learn new programming language is by directly implement the knowledge into a project.<br>
> -- Sony Arianto Kurniawan, Imaginative Go Project Maintainer

This imaginative project will show doing web application in Go lang as well as other samples that not related to web. We don't use any framework and forgive us if the code still not efficient or optimal, since it just give you an idea how to achieve something in Go lang.

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

> **NOTE:** `sudo` used throughout this doc, since mainly we use Linux/MacOS during the development. We test running on Windows 10 as well with Docker for Windows and Docker Toolbox and seems no problem.

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

## Credits
### Themes
- [Editorial](https://html5up.net/editorial) from [HTML5 UP](https://html5up.net)
- [Phantom](https://html5up.net/phantom) from [HTML5 UP](https://html5up.net)

## Contributors
- [Sony Arianto Kurniawan](https://github.com/sonyarianto) - sony@sony-ak.com - original author, project maintainer
- [Prasetyama Hidayat](https://github.com/prasetyama) - prasetyama@gmail.com

## Screen Shots
### Home Page
![Imaginative Go - Screenshot 1](/src/assets/images/screenshot1.png?raw=true "Imaginative Go - Screenshot 1")
### Sample Code List
![Imaginative Go - Screenshot 2](/src/assets/images/screenshot2.png?raw=true "Imaginative Go - Screenshot 2")

## License
This project is licensed under the MIT License.

License can be found [here](https://github.com/sonyarianto/imaginative-go/blob/master/LICENSE).
