package set_qr

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/types/known/emptypb"

	"english-ai-full/logger"
	"english-ai-full/quanqr/proto_qr/set"
	"english-ai-full/token"
)


type SetHandlerController struct {
    ctx        context.Context
    client     set.SetServiceClient
    TokenMaker *token.JWTMaker
    logger     *logger.Logger
}

func NewSetHandler(client set.SetServiceClient, secretKey string) *SetHandlerController {
    return &SetHandlerController{
        ctx:        context.Background(),
        client:     client,
        TokenMaker: token.NewJWTMaker(secretKey),
        logger:     logger.NewLogger(),
    }
}

func (h *SetHandlerController) CreateSetProto(w http.ResponseWriter, r *http.Request) {
    var setReq CreateSetRequest

    if err := json.NewDecoder(r.Body).Decode(&setReq); err != nil {
        http.Error(w, "error decoding request body", http.StatusBadRequest)
        return
    }

    h.logger.Info(fmt.Sprintf("Creating new set:CreateSetProto handler %+v", setReq))
    createdSetResponse, err := h.client.CreateSetProto(h.ctx, ToPBCreateSetProtoRequest(setReq))
    if err != nil {
        h.logger.Error("Error creating set: " + err.Error())
        http.Error(w, "error creating set", http.StatusInternalServerError)
        return
    }

    res := ToSetResFromPbSetResponse(createdSetResponse)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(res)
}

func (h *SetHandlerController) GetSetProtoDetail(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    i, err := strconv.ParseInt(id, 10, 32)
    if err != nil {
        http.Error(w, "error parsing ID", http.StatusBadRequest)
        return
    }

    h.logger.Info(fmt.Sprintf("Fetching set detail for ID: %d", i))
    setResponse, err := h.client.GetSetProtoDetail(h.ctx, &set.SetProtoIdParam{Id: int32(i)})
    if err != nil {
        h.logger.Error("Error fetching set detail: " + err.Error())
        http.Error(w, "error getting set", http.StatusInternalServerError)
        return
    }

    res := ToSetResFromPbSetResponse(setResponse)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}

func (h *SetHandlerController) GetSetProtoList(w http.ResponseWriter, r *http.Request) {
    h.logger.Info("Fetching set list")
    setsResponse, err := h.client.GetSetProtoList(h.ctx, &emptypb.Empty{})
    if err != nil {
        h.logger.Error("Error fetching set list: " + err.Error())
        http.Error(w, "failed to fetch sets: "+err.Error(), http.StatusInternalServerError)
        return
    }

    res := ToSetResListFromPbSetListResponse(setsResponse)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(res); err != nil {
        http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *SetHandlerController) UpdateSetProto(w http.ResponseWriter, r *http.Request) {
    var setReq UpdateSetRequest
    if err := json.NewDecoder(r.Body).Decode(&setReq); err != nil {
        http.Error(w, "error decoding request body", http.StatusBadRequest)
        return
    }

    h.logger.Info(fmt.Sprintf("Updating set: %d", setReq.ID))
    updatedSetResponse, err := h.client.UpdateSetProto(h.ctx, ToPBUpdateSetProtoRequest(setReq))
    if err != nil {
        h.logger.Error("Error updating set: " + err.Error())
        http.Error(w, "error updating set", http.StatusInternalServerError)
        return
    }

    res := ToSetResFromPbSetResponse(updatedSetResponse)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}

func (h *SetHandlerController) DeleteSetProto(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    i, err := strconv.ParseInt(id, 10, 32)
    if err != nil {
        http.Error(w, "error parsing ID", http.StatusBadRequest)
        return
    }

    h.logger.Info(fmt.Sprintf("Deleting set: %d", i))
    deletedSetResponse, err := h.client.DeleteSetProto(h.ctx, &set.SetProtoIdParam{Id: int32(i)})
    if err != nil {
        h.logger.Error("Error deleting set: " + err.Error())
        http.Error(w, "error deleting set", http.StatusInternalServerError)
        return
    }

    res := ToSetResFromPbSetResponse(deletedSetResponse)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}

func ToPBCreateSetProtoRequest(req CreateSetRequest) *set.CreateSetProtoRequest {
    protoReq := &set.CreateSetProtoRequest{
        Name:        req.Name,
        Description: req.Description,
        Dishes:      ToPBSetProtoDishes(req.Dishes),
        IsPublic:    req.IsPublic,
        Image:       req.Image,
    }
    if req.UserID != nil {
        protoReq.UserId = *req.UserID
    }
    return protoReq
}

func ToPBUpdateSetProtoRequest(req UpdateSetRequest) *set.UpdateSetProtoRequest {
    return &set.UpdateSetProtoRequest{
        Id:          int32(req.ID),
        Name:        req.Name,
        Description: req.Description,
        Dishes:      ToPBSetProtoDishes(req.Dishes),
        IsPublic:    req.IsPublic,
        Image:       req.Image,
    }
}

func ToPBSetProtoDishes(dishes []struct {
    Dish     Dish `json:"dish"`
    Quantity int  `json:"quantity"`
}) []*set.SetProtoDish {
    pbDishes := make([]*set.SetProtoDish, len(dishes))
    for i, dish := range dishes {
        pbDishes[i] = &set.SetProtoDish{
            DishId:   int32(dish.Dish.ID),
            Quantity: int32(dish.Quantity),
        }
    }
    return pbDishes
}

func ToSetResFromPbSetResponse(pbRes *set.SetProtoResponse) SetResponse {
    return SetResponse{
        Data: ToSetFromPbSetProto(pbRes.Data),
    }
}

func ToSetResListFromPbSetListResponse(pbRes *set.SetProtoListResponse) SetListResponse {
    sets := make([]Set, len(pbRes.Data))
    for i, pbSet := range pbRes.Data {
        sets[i] = ToSetFromPbSetProto(pbSet)
    }
    return SetListResponse{
        Data: sets,
    }
}

func ToSetFromPbSetProto(pbSet *set.SetProto) Set {
    var userID *int32
    if pbSet.UserId != nil {
        userID = pbSet.UserId
    }

    return Set{
        ID:          int(pbSet.Id),
        Name:        pbSet.Name,
        Description: pbSet.Description,
        Dishes:      ToSetDishesFromPbSetProtoDishes(pbSet.Dishes),
        UserID:      userID,
        CreatedAt:   pbSet.CreatedAt.AsTime(),
        UpdatedAt:   pbSet.UpdatedAt.AsTime(),
        IsFavourite: pbSet.IsFavourite,
        LikeBy:      pbSet.LikeBy,
        IsPublic:    pbSet.IsPublic,
        Image:       pbSet.Image,
    }
}

func ToSetDishesFromPbSetProtoDishes(pbDishes []*set.SetProtoDish) []SetDish {
    dishes := make([]SetDish, len(pbDishes))
    for i, pbDish := range pbDishes {
        dishes[i] = SetDish{
            DishID:   int64(pbDish.DishId),
            Quantity: int(pbDish.Quantity),
            Dish:     ToDishFromPbDish(pbDish.Dish),
        }
    }
    return dishes
}

func ToDishFromPbDish(pbDish *set.Dish) Dish {
    if pbDish == nil {
        return Dish{}
    }
    return Dish{
        ID:          pbDish.Id,
        Name:        pbDish.Name,
        Price:       int(pbDish.Price),
        Description: pbDish.Description,
        Image:       pbDish.Image,
        Status:      pbDish.Status,
        CreatedAt:   pbDish.CreatedAt.AsTime(),
        UpdatedAt:   pbDish.UpdatedAt.AsTime(),
    }
}

// Helper function to convert int32 to *int32
func Int32Ptr(i int) *int32 {
    if i == 0 {
        return nil
    }
    i32 := int32(i)
    return &i32
}

// Helper function to convert *int32 to *int
func Int32PtrToIntPtr(i *int32) *int {
    if i == nil {
        return nil
    }
    intVal := int(*i)
    return &intVal
}
// type SetHandlerController struct {
//     ctx        context.Context
//     client     set.SetServiceClient
//     TokenMaker *token.JWTMaker
// }

// func NewSetHandler(client set.SetServiceClient, secretKey string) *SetHandlerController {
//     return &SetHandlerController{
//         ctx:        context.Background(),
//         client:     client,
//         TokenMaker: token.NewJWTMaker(secretKey),
//     }
// }

// func (h *SetHandlerController) CreateSetProto(w http.ResponseWriter, r *http.Request) {
//     var setReq CreateSetRequest
//     if err := json.NewDecoder(r.Body).Decode(&setReq); err != nil {
//         http.Error(w, "error decoding request body", http.StatusBadRequest)
//         return
//     }

//     log.Println("Creating new set:", setReq.Name)
//     createdSetResponse, err := h.client.CreateSetProto(h.ctx, ToPBCreateSetProtoRequest(setReq))
//     if err != nil {
//         log.Println("Error creating set:", err)
//         http.Error(w, "error creating set", http.StatusInternalServerError)
//         return
//     }

//     res := ToSetResFromPbSetResponse(createdSetResponse)
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(res)
// }

// func (h *SetHandlerController) GetSetProtoDetail(w http.ResponseWriter, r *http.Request) {
//     id := chi.URLParam(r, "id")
//     i, err := strconv.ParseInt(id, 10, 32)
//     if err != nil {
//         http.Error(w, "error parsing ID", http.StatusBadRequest)
//         return
//     }

//     log.Println("Fetching set detail for ID:", i)
//     setResponse, err := h.client.GetSetProtoDetail(h.ctx, &set.SetProtoIdParam{Id: int32(i)})
//     if err != nil {
//         log.Println("Error fetching set detail:", err)
//         http.Error(w, "error getting set", http.StatusInternalServerError)
//         return
//     }

//     res := ToSetResFromPbSetResponse(setResponse)
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     json.NewEncoder(w).Encode(res)
// }

// func (h *SetHandlerController) GetSetProtoList(w http.ResponseWriter, r *http.Request) {
//     log.Println("Fetching set list")
//     setsResponse, err := h.client.GetSetProtoList(h.ctx, &emptypb.Empty{})
//     if err != nil {
//         log.Println("Error fetching set list:", err)
//         http.Error(w, "failed to fetch sets: "+err.Error(), http.StatusInternalServerError)
//         return
//     }

//     res := ToSetResListFromPbSetListResponse(setsResponse)
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     if err := json.NewEncoder(w).Encode(res); err != nil {
//         http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
//         return
//     }
// }

// func (h *SetHandlerController) UpdateSetProto(w http.ResponseWriter, r *http.Request) {
//     var setReq UpdateSetRequest
//     if err := json.NewDecoder(r.Body).Decode(&setReq); err != nil {
//         http.Error(w, "error decoding request body", http.StatusBadRequest)
//         return
//     }

//     log.Println("Updating set:", setReq.ID)
//     updatedSetResponse, err := h.client.UpdateSetProto(h.ctx, ToPBUpdateSetProtoRequest(setReq))
//     if err != nil {
//         log.Println("Error updating set:", err)
//         http.Error(w, "error updating set", http.StatusInternalServerError)
//         return
//     }

//     res := ToSetResFromPbSetResponse(updatedSetResponse)
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     json.NewEncoder(w).Encode(res)
// }

// func (h *SetHandlerController) DeleteSetProto(w http.ResponseWriter, r *http.Request) {
//     id := chi.URLParam(r, "id")
//     i, err := strconv.ParseInt(id, 10, 32)
//     if err != nil {
//         http.Error(w, "error parsing ID", http.StatusBadRequest)
//         return
//     }

//     log.Println("Deleting set:", i)
//     _, err = h.client.DeleteSetProto(h.ctx, &set.SetProtoIdParam{Id: int32(i)})
//     if err != nil {
//         log.Println("Error deleting set:", err)
//         http.Error(w, "error deleting set", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusNoContent)
// }

// // Helper functions for mapping between protobuf and Go structs
// func ToPBCreateSetProtoRequest(req CreateSetRequest) *set.CreateSetProtoRequest {
//     return &set.CreateSetProtoRequest{
//         Name:        req.Name,
//         Description: req.Description,
//         Dishes:      ToPBSetProtoDishes(req.Dishes),
//         UserId:      int32(req.UserID),
//     }
// }

// func ToPBUpdateSetProtoRequest(req UpdateSetRequest) *set.UpdateSetProtoRequest {
//     return &set.UpdateSetProtoRequest{
//         Id:          int32(req.ID),
//         Name:        req.Name,
//         Description: req.Description,
//         Dishes:      ToPBSetProtoDishes(req.Dishes),
//     }
// }

// func ToPBSetProtoDishes(dishes []SetDish) []*set.SetProtoDish {
//     pbDishes := make([]*set.SetProtoDish, len(dishes))
//     for i, dish := range dishes {
//         pbDishes[i] = &set.SetProtoDish{
//             Dish:     ToPBDish(dish.Dish),
//             Quantity: int32(dish.Quantity),
//         }
//     }
//     return pbDishes
// }

// func ToPBDish(dish Dish) *set.Dish {
//     return &set.Dish{
//         Id:          dish.ID,
//         Name:        dish.Name,
//         Price:       dish.Price,
//         Description: dish.Description,
//         Image:       dish.Image,
//         Status:      dish.Status,
//         CreatedAt:   timestamppb.New(dish.CreatedAt),
//         UpdatedAt:   timestamppb.New(dish.UpdatedAt),
//     }
// }

// func ToSetResFromPbSetResponse(pbRes *set.SetProtoResponse) SetResponse {
//     return SetResponse{
//         Data: ToSetFromPbSetProto(pbRes.Data),
//     }
// }

// func ToSetResListFromPbSetListResponse(pbRes *set.SetProtoListResponse) SetListResponse {
//     sets := make([]Set, len(pbRes.Data))
//     for i, pbSet := range pbRes.Data {
//         sets[i] = ToSetFromPbSetProto(pbSet)
//     }
//     return SetListResponse{
//         Data: sets,
//     }
// }

// func ToSetFromPbSetProto(pbSet *set.SetProto) Set {
//     return Set{
//         ID:          int(pbSet.Id),
//         Name:        pbSet.Name,
//         Description: pbSet.Description,
//         Dishes:      ToSetDishesFromPbSetProtoDishes(pbSet.Dishes),
//         UserID:      int(pbSet.UserId),
//         CreatedAt:   pbSet.CreatedAt.AsTime(),
//         UpdatedAt:   pbSet.UpdatedAt.AsTime(),
//     }
// }

// func ToSetDishesFromPbSetProtoDishes(pbDishes []*set.SetProtoDish) []SetDish {
//     dishes := make([]SetDish, len(pbDishes))
//     for i, pbDish := range pbDishes {
//         dishes[i] = SetDish{
//             Dish:     ToDishFromPbDish(pbDish.Dish),
//             Quantity: int(pbDish.Quantity),
//         }
//     }
//     return dishes
// }

// func ToDishFromPbDish(pbDish *set.Dish) Dish {
//     return Dish{
//         ID:          pbDish.Id,
//         Name:        pbDish.Name,
//         Price:       pbDish.Price,
//         Description: pbDish.Description,
//         Image:       pbDish.Image,
//         Status:      pbDish.Status,
//         CreatedAt:   pbDish.CreatedAt.AsTime(),
//         UpdatedAt:   pbDish.UpdatedAt.AsTime(),
//     }
// }