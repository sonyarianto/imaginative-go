# Imaginative Go (Work in Progress)
## What is Imaginative Go?
Imaginative Go is a collection of free code samples in the form of Go web application.

We believe that when learning programming language is by directly implement the knowledge into a project. This imaginative project will show doing web application in Go language as well as other samples that not related to web. We don't use any framework and forgive us if the code still not efficient or optimal, since this project will just give you an idea how to achieve something in Go language.

## Why Imaginative Go?
- Ideas about achieve something with Go language
- Plenty of working code samples
- Easy to run (with or without Docker, personally I prefer with Docker)
- You can contribute

## Requirements
- Docker Engine (version 17.03.0 or later)
- Docker Compose (version 1.22.0 or later)

Docker is used since it will create additional service to mimic realistic web sample. Imaginative Go is using MySQL and MongoDB with pre-populated data. It's useful to show code that doing query to those database service (SQL and noSQL).

For those who need documentation of Docker installation, please refer to [Docker CE](https://store.docker.com/search?type=edition&offering=community) and [Docker Compose](https://docs.docker.com/compose/install/).

> **NOTE:** `sudo` used throughout this doc, since mainly we use Linux/MacOS.

## Usage
#### Linux/MacOS
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

> To clean containers (stop and remove), clean network (remove), pull, build (with no cache) and bring up all containers you can type like below
> ```
> sudo docker-compose down && sudo docker-compose pull && sudo docker-compose build --no-cache && sudo docker-compose up -d --build --force-recreate
> ```
> Above command will make sure you will get create and run fresh all containers needed to run Imaginative Go web project. This is useful after you are doing `git pull` on Imaginative Go repository.

> **Note 2**

> You can access the MySQL table using Adminer container that available on the following address.
> ```
> http://localhost:8989
> ```

> **Note 3**

> MySQL and MongoDB expose random port to host machine. You can see it by typing this after all containers are running.
> ```
> sudo docker ps -f "name=mysql-imaginative-go" -f "name=mongodb-imaginative-go"
> ```
> Sample output is like below
> ```
> CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS         >                        NAMES
> 8932a72252b0        imaginative-go_mongodb   "docker-entrypoint.s…"   20 minutes ago      Up 20 minutes       0.0.0.0:32818->27017/tcp             mongodb-imaginative-go
> 6516ccc7418d        imaginative-go_mysql     "docker-entrypoint.s…"   20 minutes ago      Up 20 minutes       33060/tcp, 0.0.0.0:32817->3306/tcp   mysql-imaginative-go
> ```

## Docker Images Used
- `mysql:latest` (see `docker-compose.yml` for default credential, port exposed to host is using random port, see on `docker ps` command)
- `mongodb:latest` (see `docker-compose.yml` for default credential, port exposed to host is using random port, see on `docker ps` command)
- `adminer` (see `docker-compose.yml` for default credential to MySQL database, port exposed to host at 8989)

## Credits
### Themes
- [Editorial](https://html5up.net/editorial) from [HTML5 UP](https://html5up.net)
- [Phantom](https://html5up.net/phantom) from [HTML5 UP](https://html5up.net)

## Contributors
- Sony Arianto Kurniawan - sony@sony-ak.com - original author, project maintainer
- Prasetyama Hidayat - prasetyama@gmail.com

## Screen Shots
### Home Page
![Imaginative Go - Screenshot 1](/src/assets/images/screenshot1.png?raw=true "Imaginative Go - Screenshot 1")
### Sample Code List
![Imaginative Go - Screenshot 2](/src/assets/images/screenshot2.png?raw=true "Imaginative Go - Screenshot 2")
