package router
import(
		"github.com/sunshiine1225/TODO-APPLICATION/middleware"
		"github.com/gorilla/mux"
)
func Router() *mux.Router{

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks",middleware.GetAllTasks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/tasks",middleware.CreateTask).Methods("POST","OPTIONS")
	router.HandleFunc("/api/tasks/{id}",middleware.TaskComplete).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/undoTask/{id}",middleware.UndoTask).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deleteAllTasks",middleware.DeleteAllTasks).Methods("DELETE","OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}",middleware.DeleteTask).Methods("DELETE","OPTIONS")


	return router

}