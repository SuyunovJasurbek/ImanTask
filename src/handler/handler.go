package handler

import "days_limit/src/service"

type Handler struct {
	service *service.Service
}

func  NewHandler(s *service.Service) *Handler {
	return  &Handler{service: s}
}