# Drones task
![example workflow](https://github.com/YahyaQandel/drones/actions/workflows/main.yml/badge.svg)
### Prerequisites to run
##### - install docker compose
follow that [link](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-20-04) to install docker compose (ubuntu 20.04)
then

```
docker compose up
```


### Tests
```
cd docker-tests
docker compose up --abort-on-container-exit --exit-code-from go
docker compose down -v
```