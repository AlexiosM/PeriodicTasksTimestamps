# Welcome to PeriodicTasksTimestamps!

# Files

A periodic task is described by the following properties: 
* Period (every hour, every day, ...) 
* Invocation point (where inside the period should be invoked) 
* Timezone (days/months/years are timezone-depended) 

The service returns all matching timestamps of a periodic task (ptlist) between 2 points in time (t1, t2). t1, t2 and the entries of ptlist are in UTC with seconds accuracy, in the following form: 20060102T150405Z The supported periods should be: 1h, 1d, 1mo, 1y. The invocation timestamp should be at the start of the period (e.g. for 1h period a matching timestamp is considered the 20210729T010000Z).

## Run the service
### Run in docker
> make all

To clean the environment
>make clean

### Run locally
> go run src/*

### Perform a request
From an http client run command as the following 
>curl --location --request GET 'http://localhost:3000/ptlist?period=200d&tz=Europe/Athens&t1=20200714T204603Z&t2=20230715T204603Z'

### Perform a request
Responses from the service

|Result|Code|Response|
|---|---|---|
|Success|200 OK| ["20201022T200000Z","20210130T200000Z","20210510T200000Z","20210818T200000Z","20211126T200000Z","20220306T200000Z"]|
|Failure|400 Bad Request|{"Status":"error","Desc":"Unsupported period"}|
## API specifications

Service API endpoint is at **/plist**

Service parameters are the following
t1=&t2=

|    period    | tz    |      t1      | t2 |
|----------|-------------------|-----------|------|
| A number of hours/days/months/years | TimeZone   |   start time  (format: YYYYMMDD**T**HHMMSS**Z**)   | end time (format: YYYYMMDD**T**HHMMSS**Z**)|
| \<number\>h,d,mo,y| Europe/Athens   |   20200714T204603Z   | 20230715T204603Z|

## Design Notes
Interfaces have been used for the available periods extensibility
Periods support various numbers (e.g 2d,10h etc.)
Input validation for security reasons is supported
[*zerolog*](https://pkg.go.dev/github.com/rs/zerolog) has been used for logging for easier parsing and troubleshooting
