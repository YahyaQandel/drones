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

### Features
what i have developed is marked by ✔ and what i wish i had time and could have developed is marked by X.
` i will still add to that repo on my github account even after deadline to make it useful and in proper state`

    - [✔] Register drone
    - [✔] Register medication 
    - [✔] Load drone with medication
    - [✔] Checking loaded medication items for a given drone;
    - [✔] checking available drones for loading;
    - [✔] check drone battery level for a given drone;
    - [✔] add swagger for apis reference
    - [✔] A drone battery decreases each 5 seconds by `3%` if his state is `LOADED`
    - [X] periodic task checks drone battery level
    - [X] add audit/history for all drones battery levels


### System behaviors
1- A drone once loaded its state should be changed from `IDLE` to `LOADED`
2- A drone while registering his battery capacity is `100%` default
3- You can load a drone with only one medication in each call , if you want to load it with multiple medications
you will need to call `/drone/load` multiple times each with the new medication 
3- Medication code is a unique identifier
4- images saved with medication code
5- Images are not uploaded to any server it is just saved on your repo , what should be done is to push images to remote server specific for images (amazon bucket for example)


### Functionality 

* Register drone request
```
curl --request POST \
  --url http://localhost:9090/api/drone/ \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME' \
  --header 'Content-Type: application/json' \
  --data '{
	"serial_number": "firstDrone",
	"model": "Lightweight",
	"weight": 60
}'
```

* Register medication
```
curl --request POST \
  --url http://localhost:9090/api/medication/ \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME' \
  --header 'Content-Type: multipart/form-data; boundary=---011000010111000001101001' \
  --form image=@/home/yahia/Pictures/2022-04-05_12-55.png \
  --form name=panadol \
  --form weight=20.6 \
  --form code=PDL1
  ```
  
  * Load drone with medication
  ```
  curl --request POST \
  --url http://localhost:9090/api/drone/load \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME' \
  --header 'Content-Type: application/json' \
  --data '{
	"drone_serial_number":"firstDrone",
	"medication_code":"PDL1"
}'
```

* Get all loaded medication items
```
curl --request GET \
  --url 'http://localhost:9090/api/drone/medication?drone_serial_number=firstDrone' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME'
  ````

* Get all available drones ( state IDLE )
```
curl --request GET \
  --url http://localhost:9090/api/drone/available \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME'
```
* Get drone battery level
```
curl --request GET \
  --url 'http://localhost:9090/api/drone/battery?drone_serial_number=secondDrone' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME'
```

### System diagram
![Graph](/drone.jpg "system design")
