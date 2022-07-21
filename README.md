<h3>Drone attack!</h3>

So be careful, and make sure your Drone Navigation Services are deployed on sectors, and you ready to link them

First, make sure your application runs, just `docker compose up -d` and check logs<br/>
Second, if something went wrong, check **.env** file<br/>
Third, lets check out existed endpoints:

- add a bunch of sectors `host:port/sectors/add`
- get list of sectors `host:port/sectors/list`
- locate DNS on sector `host:port/sector/locate?id={existed_sector_id}`
- get list of DNS `host:port/dns/list`

The application uses IMDB resources, so it will clean up when stopped
