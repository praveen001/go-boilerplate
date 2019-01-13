package router

func userRouter(r *CustomRouter) {
	r.post("/register", r.RegisterUser)
}
