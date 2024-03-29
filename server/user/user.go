package user

type Handler struct {
	controller *Controller
}

func NewHandler() *Handler {
	dao := NewDao()
	usecase := NewUsecase(dao)
	controller := NewController(usecase)

	return &Handler{
		controller: controller,
	}
}

// func (h *Handler) Process(action string, data []byte) error {
// 	switch action {
// 	case "auth":
// 		return h.controller.Register(data)

// 	default:
// 		return errors.New("invalid action")
// 	}
// }
