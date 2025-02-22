# URL SHORTENER

## Business Rules
|   ID  | Description |
| --- | --- |
| BR-01 | The URL is composed by numbers (0-9) and chars (a-z, A-Z) |


## Functional Requirements
> 01. URL shortening
> 02. URL redirecting

|   ID  | Description |
| --- | --- |
| RF-01 | The user will receive a shorter url given the original one |

|   ID  | Description |
| --- | --- |
| RF-02 | Given the shortest url, the user will be redirected to the original url |

## Non-Functional Requirements
> 01. 10 million users per day (just for study purposes)
> 02. Availability and scalability
> 03. Security

## HTTP requests
| HTTP | URI | Description |
| -- | -- | -- |
|   _POST_ | /api/v1/data/shorten | Return the shortened url with the status code 200 | 
|   _GET_ | /api/v1/:short-url | Return the shortened url with the status code 200 | 


## User stories
### User Story #1: Create a shorter URL (CREATE)
| US #1 | Create a shorter url | 
| -- | -- |
| **Description** |  As a user, I want to receive a shorter URL |
| **Case 1 - URL does not exist** | If the original url has not been shortened before and doesn't exist in the database, the system will create a new shortened url based on the ID of the new entry |
| **Case 2 - The URL exists** | If the original url already exists in the database, we'll return the existing shortened url to the user |
| **Validation** | |

## Back of the envelope estimation

## System Design

## Challenges
> 01. Creating a distributed id isn`t easy. Here we'll explore the Twitter snowflake approach.
> 02. Caching the data


