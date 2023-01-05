# order

Please make sure MySQL already running and had a database with name gank_order 


Create .env file with format below: 

DB_USERNAME = 

DB_PASSWORD = 

DB_HOST = 

DB_PORT = 

DB_NAME = 

SERVICE_PORT = :8030

CUSTOMER_SERVICE_URI = # set this url to where is customer service api run

PRODUCT_SERVICE_URI = # set this url to where is product service api run


Run with command 
$ go run main.go