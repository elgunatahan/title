# Running 

- install docker
- run `docker compose up` in terminal 
- Application will start at port 8080
- Before using the application , run sql script for creating required tables and records.

# Technical Details
- Go
- GORM
- Postgre
- Docker

# Example Requests:

```sh 

GET /series?name=dark&genre3&page=1&pagesize=1      # Search series
GET /series/1                                       # Get series by Id
DELETE /series/1                                    # Soft delete series by Id

GET /movies?name=a&genre3&page=1&pagesize=1         # Search movies
GET /movies/1                                       # Get movie by Id
DELETE /movie/1                                     # Soft delete movie by Id. By the soft delete we can stream an event. In consumer with mutliple options we can process hard delete of the title.

GET /accounts/2/favorites?name=a&genre3&page=1&pagesize=2         # Search favorites
DELETE /accounts/2/favorites                                      # Hard delete favorite by Id
POST /accounts/2/favorites                                        # Create favorite
{
    "TitleId":3
}

POST /titles                                        # Create title. After than we can create episodes, seasons ...

#for creating movie
{
    "Name":"elgun movie",
    "Type":0,
    "IMDBId":7,
    "Duration":240,
    "Description":"what is about?",
    "Rating":9.5,
    "ReleaseDate":"2008-06-25T00:00:00Z"
}

#for creating series
{
    "Name":"elgun series",
    "Type":1,
    "IMDBId":8,
    "Duration":60,
    "Description":"what is about this series?",
    "Rating":9.4,
    "ReleaseDate":"2008-06-26T00:00:00Z"
}

```