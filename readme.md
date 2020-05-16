##Actant test task

- Running server 

    `go run . -send_port (port int)`
    
    Server starting on port 8899
    
- Running server in docker.

    1. `docker build -t image_name .`
    2. `docker run -p 127.0.0.1:8899:8899 image_name`
   
    Server starting in docker container with exposed port 8899 
    
- Run test client.

    Test client sends couple of requests to server on port 8899
    1. `cd project_root/test_client`
    2. `go run .`
    