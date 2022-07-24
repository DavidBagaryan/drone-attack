<h3>Drone attack!</h3>

So be careful, and make sure your Drone Navigation Services are deployed on sectors, and you ready to link them

First, make sure your application runs, just `docker compose up -d` and check logs<br/>
Second, if something went wrong, check **.env** file<br/>
Third, lets check out existed endpoints:

for ver 1.0.0

- add a bunch of sectors `POST host:port/sectors/add`
- get list of sectors `GET host:port/sectors/list`
- locate DNS on sector `POST host:port/sector/locate?id={existed_sector_id}`
- get list of DNS `GET host:port/dns/list`

for ver ^1.1.0

- add a bunch of sectors `POST host:port/sectors`
- get list of sectors `GET host:port/sectors`
- locate DNS on sector `POST host:port/sector/{sectorID}/locate`
- get list of DNS `GET host:port/dns`

The application uses IMDB resources, so it will clean up when stopped
