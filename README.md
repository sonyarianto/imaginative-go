# Imaginative Go (Work in Progress)
Imaginative Go is a free code sample in the form of Go web application.

We believe that when learning programming language is by directly implement the knowledge into a project. This imaginative project will show doing web application in Go language. We don't use any framework and forgive us if the code still not efficient or optimal, since this sample just will give you an idea how to achieve something in Go language.

## Requirements
- Docker Engine (version 17.03.0 or later)
- Docker Compose (version 1.22.0 or later)

Docker is used since it will create additional service to mimic realistic web sample. Imaginative Go is using MySQL and MongoDB with pre-populated data. It's useful to show code that doing query to those database service (SQL and noSQL).

For those who need documentation of Docker installation, please refer to [Docker CE](https://store.docker.com/search?type=edition&offering=community) and [Docker Compose](https://docs.docker.com/compose/install/)

## Usage
#### Linux/MacOS
Just do this.

```
git clone https://github.com/sonyarianto/imaginative-go.git
cd imaginative-go
sudo docker-compose up
```

After that go to your browser and type
```
http://localhost:9899
```
or
```
http://<YOUR_LOCAL_IP_ADDRESS>:9899
```
or
```
http://<YOUR_DOCKER_MACHINE_IP_ADDRESS>:9899
```

## Docker Images Used
- mysql:latest
- mongodb:latest
- adminer (port exposed 8989) (host: mysql, username: root, password: mysqlpassword)

## Credits
### Themes
- [Editorial](https://html5up.net/editorial) from [HTML5 UP](https://html5up.net)
- [Phantom](https://html5up.net/phantom) from [HTML5 UP](https://html5up.net)

## Contributors
Sony Arianto Kurniawan - sony@sony-ak.com - original author, project maintainer
