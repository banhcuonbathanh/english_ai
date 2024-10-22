package order_grpc

import (
	"time"
)

type OrderDish struct {
    DishID   int64 `json:"dish_id"`
    Quantity int64 `json:"quantity"`
}


type OrderSet struct {
    SetID   int64 `json:"set_id"`
    Quantity int64 `json:"quantity"`
}


type OrderType struct {
    ID             int64           `json:"id"`
    GuestID        int64           `json:"guest_id"`
    UserID         int64           `json:"user_id"`
    IsGuest        bool            `json:"is_guest"`
    TableNumber    int64           `json:"table_number"`
    OrderHandlerID int64           `json:"order_handler_id"`
    Status         string          `json:"status"`
    CreatedAt      time.Time       `json:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at"`
    TotalPrice     int32           `json:"total_price"`
    DishItems      []OrderDish `json:"dish_items"`
    SetItems       []OrderSet  `json:"set_items"`
    BowChili       int64           `json:"bow_chili"`
    BowNoChili     int64           `json:"bow_no_chili"`
}

// CreateOrderRequest struct
type CreateOrderRequestType struct {
    GuestID        int64           `json:"guest_id"`
    UserID         int64           `json:"user_id"`
    IsGuest        bool            `json:"is_guest"`
    TableNumber    int64           `json:"table_number"`
    OrderHandlerID int64           `json:"order_handler_id"`
    Status         string          `json:"status"`
    CreatedAt      time.Time       `json:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at"`
    TotalPrice     int32           `json:"total_price"`
    DishItems      []OrderDish `json:"dish_items"`
    SetItems       []OrderSet  `json:"set_items"`
    BowChili       int64           `json:"bow_chili"`
    BowNoChili     int64           `json:"bow_no_chili"`
}

// UpdateOrderRequest struct
type UpdateOrderRequestType struct {
    ID             int64           `json:"id"`
    GuestID        int64           `json:"guest_id"`
    UserID         int64           `json:"user_id"`
    TableNumber    int64           `json:"table_number"`
    OrderHandlerID int64           `json:"order_handler_id"`
    Status         string          `json:"status"`
    TotalPrice     int32           `json:"total_price"`
    DishItems      []OrderDish `json:"dish_items"`
    SetItems       []OrderSet  `json:"set_items"`
    IsGuest        bool            `json:"is_guest"`
    BowChili       int64           `json:"bow_chili"`
    BowNoChili     int64           `json:"bow_no_chili"`
}




// GetOrdersRequest struct
type GetOrdersRequestType struct {
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
	UserID   *int64    `json:"user_id,omitempty"`
	GuestID  *int64    `json:"guest_id,omitempty"`
}

// PayOrdersRequest struct
type PayOrdersRequestType struct {
	GuestID *int64 `json:"guest_id,omitempty"`
	UserID  *int64 `json:"user_id,omitempty"`
}

// OrderResponse struct
type OrderResponse struct {
	Data OrderType `json:"data"`
}

// OrderListResponse struct
type OrderListResponse struct {
	Data []OrderType `json:"data"`
}

// OrderIDParam struct
type OrderIDParam struct {
	ID int64 `json:"id"`
}

// OrderDetailIDParam struct
type OrderDetailIDParam struct {
	ID int64 `json:"id"`
}

// Guest struct
type Guest struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	TableNumber int32     `json:"table_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}


type OrderDetailedDish struct {
    DishID      int64  `json:"dish_id"`
    Quantity    int64  `json:"quantity"`
    Name        string `json:"name"`
    Price       int32  `json:"price"`
    Description string `json:"description"`
    Image       string `json:"image"`
    Status      string `json:"status"`
}

type OrderSetDetailed struct {
    ID          int64            `json:"id"`
    Name        string           `json:"name"`
    Description string           `json:"description"`
    Dishes      []OrderDetailedDish `json:"dishes"`
    UserID      int32            `json:"userId"`
    CreatedAt   time.Time        `json:"created_at"`
    UpdatedAt   time.Time        `json:"updated_at"`
    IsFavourite bool             `json:"is_favourite"`
    LikeBy      []int64          `json:"like_by"`
    IsPublic    bool             `json:"is_public"`
    Image       string           `json:"image"`
    Price       int32            `json:"price"`
}

type OrderDetailedListResponse struct {
    Data []OrderSetDetailed `json:"data"`
}
