package set_qr

import "time"

type Dish struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Price       int32     `json:"price"`
    Description string    `json:"description"`
    Image       string    `json:"image"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type SetSnapshot struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Dishes      []SetDish `json:"dishes"`
    UserID      int       `json:"userId"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
	SetID      int       `json:"setId"`
}


type Set struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Dishes      []SetDish `json:"dishes"`
    UserID      int       `json:"userId"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type SetDish struct {
    Dish     Dish `json:"dish"`
    Quantity int  `json:"quantity"`
}





type CreateSetRequest struct {
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Dishes      []SetDish `json:"dishes"`
    UserID      int       `json:"userId"`
}

type UpdateSetRequest struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Dishes      []SetDish `json:"dishes"`
}

type SetResponse struct {
    Data    Set    `json:"data"`

}

type SetListResponse struct {
    Data    []Set  `json:"data"`

}



