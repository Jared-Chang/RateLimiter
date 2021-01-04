# RateLimiter

### Test with API test
Use following command to see the testing result

```shell
go test -v ApiTest
```

### Test with your own testing tool

Deploy to your local Docker environment by following command

```shell
docker-compose build
docker-compose up -d
```

Then you can use your own testing tool, the entry url is

```shell
127.0.0.1:8088/AccessCount
```

### Test in other environment

Cause this project not consider ip spoofing attack.

So if you want to deploy this project to other environment, like AWS, GCP, Other machine, etc.

Please make sure you have no proxy in front of the rate limiter service.

### TODO

Integrate RedisTimeSeries(https://oss.redislabs.com/redistimeseries/) to increase extensibility and more efficient data access. 


Integrate cache mechanism like freecache(https://github.com/coocood/freecache) to eliminate useless query from time series data.
