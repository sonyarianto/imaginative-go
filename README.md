# Imaginative Go (ALPHA)
Imaginative Go is a self-hosted website that contains real world code example using real world beautiful layout to make code sample realistic!

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
Sony Arianto Kurniawan - sony@sony-ak.com - original author
