package middleware
import(
	"context"
	"encoding/json"
	"log"
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/sunshiine1225/TODO-APPLICATION/models"



	
)
var collection *mongo.Collection 
func init(){
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv(){
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("error in loading env file")

	}
}

func createDBInstance(){
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	coll:= os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err!=nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("connected to mangodb")
	collection = client.Database(dbName).Collection(coll)
	fmt.Println("collection instance created")
} 

func GetAllTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	payload := getAllTasks()
	fmt.Println("in all tasks")
	json.NewEncoder(w).Encode(payload)

}
func CreateTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-ALlow-Methods" , "POST")
	w.Header().Set("Access-Control-ALlow-Headers","Content-Type")
  
	var task models.ToDoList
	err := json.NewDecoder(r.Body).Decode(&task)
	fmt.Println(err)
	insertResult, err := collection.InsertOne(context.Background(),task)
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("imserter a single record",insertResult.InsertedID)
	json.NewEncoder(w).Encode(task)
}
func TaskComplete(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-ALlow-Methods" , "PUT")
	w.Header().Set("Access-Control-ALlow-Headers","Content-Type")
    params := mux.Vars(r)
	taskComplete(params["id"])
	
	json.NewEncoder(w).Encode(params["id"])
}


func UndoTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-ALlow-Methods" , "PUT")
	w.Header().Set("Access-Control-ALlow-Headers","Content-Type")
	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-ALlow-Methods" , "DELETE")
	w.Header().Set("Access-Control-ALlow-Headers","Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
}
func DeleteAllTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
    count := deleteAllTasks()
	json.NewEncoder(w).Encode(count)
}
func getAllTasks()[]primitive.M{
	curr , err := collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)

	}
	var results []primitive.M
	for curr.Next(context.Background()){
var result bson.M
e:=curr.Decode(&result)
if e!=nil{
	log.Fatal(e)
}
results = append(results,result)
}

curr.Close(context.Background())
return results


}

func taskComplete(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	update :=bson.M{"$set":bson.M{"status":true}}
	result, err := collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("modified :",result.ModifiedCount)


}

func insertOneTask(task  models.ToDoList){

	


}
func deleteOneTask(task string){
	id,_ :=primitive.ObjectIDFromHex(task)

	filter := bson.M{"_id":id}
	result, err := collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("modified :",result.DeletedCount)

}
func undoTask(task string){
	id,_ :=primitive.ObjectIDFromHex(task)
	fmt.Println(id)
	filter := bson.M{"_id":id}
	update :=bson.M{"$set":bson.M{"status":true}}
	result, err := collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("modified :",result.ModifiedCount)

}
func deleteAllTasks() int64{
	
	result, err := collection.DeleteMany(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("modified :",result.DeletedCount)
	return result.DeletedCount


}