# URL SHORTENER

## 1. Business Rules
|   ID  | Description |
| --- | --- |
| BR-01 | The URL is composed by numbers (0-9) and chars (a-z, A-Z) |


## 2. Functional Requirements
> 01. URL shortening
> 02. URL redirecting

|   ID  | Description |
| --- | --- |
| RF-01 | The user will receive a shorter url given the original one |

|   ID  | Description |
| --- | --- |
| RF-02 | Given the shortest url, the user will be redirected to the original url |

## 3. Non-Functional Requirements
> 01. 10 million users per day (just for study purposes)
> 02. Availability and scalability
> 03. Security

## 4. HTTP requests
| HTTP | URI | Description |
| -- | -- | -- |
|   _POST_ | /api/v1/data/shorten | Return the shortened url with the status code 201 | 
|   _GET_ | /api/v1/:short-url | Redirect the user with status code 301/302 | 

### JSON representation
#### POST
```
{    
    long-url: "www.longUrl.com/test"
} 
```

## 5. User stories
### User Story #1: Create a shorter URL (CREATE)
| US #1 | Create a shorter url | 
| -- | -- |
| **Description** |  As a user, I want to receive a shorter URL |
| **Case 1 - URL does not exist** | If the original url has not been shortened before and doesn't exist in the database, the system will create a new shortened url based on the ID of the new entry |
| **Case 2 - The URL exists** | If the original url already exists in the database, we'll return the existing shortened url to the user |
| **Validation** | Validate that there are only chars and numbers in the url (BR-01) |

### User Story #2: Redirected with the new URL (READ)
| US #2 | Visit a site | 
| -- | -- |
| **Description** |  As a user, I want to be redirected with my new shortened url |
| **Case 1 - URL does not exist in database** | If the shortened url doesn't exist in the database, the system will return an error |
| **Case 2 - The URL exists** | If the shortened url already exists in the database, we'll redirect the user to the original url |
| **Validation** | Validate that the url already exists |

## 6. Back of the envelope estimation

## 7. Technologies
> 01. Sql database (Mysql)
> 02. Cache - Redis
> 03. Nginx server
> 04. Nginx load balancer
> 05. Nginx Rate limiter
> 06. Message Broker - RabbitMQ
> 07. Observability - Prometheus/Grafana


## 8. System Design
#### High Level Design
![url-shortener (5)](https://github.com/user-attachments/assets/d8af5e24-f5ca-4d3f-8346-325471ec63fd)

#### Expiration Cache Policy 
#### Load Balancer x App Load Balancer x API Gateway
#### Mysql Master and Slaves instancies



