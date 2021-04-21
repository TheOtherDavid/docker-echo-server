# docker-echo-server
A simple echo server, Dockerized!

To build the code, in the code's base directory, run
```
docker build . -t docker-echo-server
```
To run the code, in the code's base directory, run
```
docker run -p 10000:10000 docker-echo-server
```
When the server is running, as a GET request, hit:
http://localhost:10000/echo/example
with any text in place of "example", and the app will echo it back at you.