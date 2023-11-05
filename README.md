# gRPC Client Server architecture README


This Go-based application harnesses the capabilities of gRPC to facilitate a client-server architecture.The application offers two indispensable methods, **GetUserByUserId** and **GetUserListByIds** to cater to the efficient retrieval of user data. With **GetUserByUserId**, clients can fetch specific user information by providing their unique user ID. On the other hand, **GetUserListByIds** allows clients to retrieve data for multiple users by specifying an array of user IDs. These methods are designed with performance and ease of use in mind, ensuring that users can obtain the data they need quickly and conveniently. What truly distinguishes this application is its pioneering use of bidirectional streaming. This communication method empowers real-time interactions between clients and the server. Unlike traditional request-response patterns, bidirectional streaming allows for concurrent, independent data transmission in both directions. 


## Prerequisites
* Go (1.21.3 or later)
* Docker (for containerization)
* Postman or any gRPC client tool for testing

## Quick Start

You can get this application up and running with just one command, assuming you have Docker installed.
```bash
docker run -it -p 5501:5501 krishh182307/totality_golang_project
```

## Getting Started
Follow these steps to run and test the gRPC Client-Server application.

### 1. Clone the Repository
Clone the repository using the command
~~~ 
git clone https://github.com/Kkrishna19/totality_golang_project.git
~~~
```
cd totality-golang-project
```

### 2. Run the Docker Container
Launch a Docker container using the image, and ensure that the gRPC service port (e.g., 5501) is mapped to a port on your host machine (e.g., 5501).

```bash
docker run -it -p 5501:5501 krishh182307/totality_golang_project
```
### 4. Test the gRPC Service
Use a gRPC client (e.g., BloomRPC, grpcurl) or Postman to connect to the gRPC service.

* Method 1: GetUserByUserId
Submit a gRPC request to "localhost:5501" using the "GetUserByUserId" method. Specify a User_ID to retrieve information about a specific user.
#### Sample Request
```
{
    "User_Id" : 1
}
```

* Method 2: GetUserListByIds
Submit a gRPC request to the address "localhost:5501" using the "GetUserListByIds" method. Supply an array of User_IDs to retrieve data for multiple users.

#### Sample Request
```
{
   "UserRequestList": [{"UserId":1}, {"UserId":5}, {"UserId":10},  {"UserId":100}]
}
```

## Acknowledgments
Special thanks to the gRPC community for their support and resources.