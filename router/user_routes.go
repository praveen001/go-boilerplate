package router

func userRouter(r *CustomRouter) {
	r.get("/register", r.RegisterUser)
}
