package model

import "github.com/google/uuid"

type Subscription struct {
	ID          uuid.UUID  `json:"id" binding:"uuid"`
	ServiceName string     `json:"service_name" binding:"required"`
	Price       int        `json:"price" binding:"required,gte=0"`
	UserID      uuid.UUID  `json:"user_id" binding:"required,uuid"`
	StartDate   MonthYear  `json:"start_date" binding:"required"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
}
type SubscriptionCreate struct {
	ServiceName string     `json:"service_name" binding:"required"`
	Price       int        `json:"price" binding:"required,gte=0"`
	UserID      uuid.UUID  `json:"user_id" binding:"required,uuid"`
	StartDate   MonthYear  `json:"start_date" binding:"required"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
}
type SubscriptionUpdate struct {
	ServiceName *string    `json:"service_name,omitempty"`
	Price       *int       `json:"price,omitempty"`
	UserID      *uuid.UUID `json:"user_id,omitempty"`
	StartDate   *MonthYear `json:"start_date,omitempty"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
}
type SubscriptionsTotal struct {
	ServiceName  string `form:"service_name"`
	UserIDString string `form:"user_id"`
	StartString  string `form:"start" `
	EndString    string `form:"end" `
	UserID       uuid.UUID
	Start        MonthYear
	End          MonthYear
}

type SubscriptionsListRequest struct {
	Page     int `form:"page,default=1" binding:"gte=1"`
	PageSize int `form:"page_size,default=10" binding:"gte=1,lte=20"`
}
