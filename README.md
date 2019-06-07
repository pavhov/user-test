**Init case**
    
    1. You will run `make init-dep` for install dep package manager for this service
    2. You will run `make init-services` for copy your local enviroment variables or `env.local`
        you will set postgres databese your enviroments on this file
**Services**

    1. You will run `make build-services` for build services binaries
    2. You will run `make run-services` for run initial-job and user-service
        on initial-job service I insert 20.000.000 conn_logs records for test
        on this logs is insert (
                            "127.0.0.1",
                            "127.0.0.2",
                            "127.0.0.3",
                            "127.0.0.4",
                            "127.0.0.5",
                            "127.0.0.6",
                            "127.0.0.7",
                            "127.0.0.8",
                            "127.0.0.9",
                            "127.0.0.10",
                            "127.0.0.11",
                            "127.0.0.12",
                            "127.0.0.13",
                            "127.0.0.14",
                            "127.0.0.15",
                            "127.0.0.16",
                            "127.0.0.17",
                            "127.0.0.18",
                            "127.0.0.19",
                            "127.0.0.20",
                        ) ip addreses
        and 22 users on ids 1-22
        1-20 id users have duplicates
        21 and 22 id users don't have duplicates  

**Test steps**

    `Get request: http://localhost:12345/1/2 Response:
     Responce { "dupes": true }

     Get request: http://localhost:12345/1/3 Response:
     
     Responce { "dupes": true }

     Get request: http://localhost:12345/2/1 Response:
     
     Responce { "dupes": true }

     Get request: http://localhost:12345/2/21 Response:
     
     Responce { "dupes": false }

     Get request: http://localhost:12345/3/22 Response:
     
     Responce { "dupes": false }

     Get request: http://localhost:12345/1/23 Response:
     
     Responce { "dupes": false }`