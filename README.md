# Golang_CRUD_API
A Simple RESTFUL API build using golang  version 1.22 net/http standard library for performing Cread, Read, Update, Delete (CRUD) operation. GORM for ORM and PostgreSQL for database.  
# Preprerequisite
1. Have PostgreSQL installed and run on your machine
2. Create new database 
3. Edit ``.env`` by applying your configuration  
# How To Use
1. Clone this repo by run following script
```
git clone https://github.com/W-ptra/Golang_CRUD_API.git
```
2. Change Directory to main folder & Download all dependency
```
cd Golang_CRUD_API
go mod download
```
3. Run the main.go & api.go
```
go run main.go api.go
```
# Using Docker
Run following script  
```
docker run -d -p 8080:8080 \
	--name golang_crud_api \
	--network {your_network_name *optional} \
	-e HOST="0.0.0.0" \
	-e PORT="8080" \
	-e DB_HOST="{your postgreSQL host}" \
	-e DB_PORT="{your postgreSQL port}" \
	-e DB_USER="{your postgreSQL user}" \
	-e DB_PASSWORD="{your postgreSQL password}" \
	-e DB_NAME="{your database name}" \
	-e DB_SSLMODE="disable" \
	wisnup001binus/golang_crud_api:1.0
```
# API Endpoint
1. ``GET http://127.0.0.1:8080/api/student``  
Retrive all student data  
``Example Respond:``  
```
[
    {
        "Id": 1,
        "Name": "zaza",
        "Age": 22,
        "GPA": 3.12,
        "Street": "street 1",
        "Province": "province 1",
        "Country": "indonesia",
        "CreatedAt": "2024-09-22T12:49:20.106021Z"
    },
    {
        "Id": 2,
        "Name": "renaldi",
        "Age": 19,
        "GPA": 3.5,
        "Street": "street 1",
        "Province": "Ohio",
        "Country": "America",
        "CreatedAt": "2024-09-22T12:49:20.949791Z"
    }
]
```  
2. ``GET http://127.0.0.1:8080/api/student/{id}``  
Retrive specifict student data identify by id  
``Example Respond:``  
```
{
        "Id": 1,
        "Name": "zaza",
        "Age": 22,
        "GPA": 3.12,
        "Street": "street 1",
        "Province": "province 1",
        "Country": "indonesia",
        "CreatedAt": "2024-09-22T12:49:20.106021Z"
}
```  
3. ``POST http://127.0.0.1:8080/api/student``  
Create a new student  
``Example Request:``  
```
{
    "name":"zaza",
    "age":22,
    "gpa":3.12,
    "street":"street 1",
    "province":"province 1",
    "country":"indonesia"
}
```  
``Example Respond:``  
```
{
    "message": "successfully created new student"
}
```  
4. ``PUT http://127.0.0.1:8080/api/student/{id}``  
Update existing student specify by id  
``Example Request:``  
```
{
    "name":"dimas",
    "age":35,
    "gpa":3.12,
    "street":"street 1",
    "province":"province Z",
    "country":"Wakanda"
}
```  
``Example Respond:``  
```
{
    "message": "successfully update student"
}
```  
5. ``DELETE http://127.0.0.1:8080/api/student/{id}``  
Delete a student specify by id  
``Example Respond:``  
```
{
    "message": "successfully update student"
}
``` 