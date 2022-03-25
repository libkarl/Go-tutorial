package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://karel_data:karl@cluster0.zsir6.mongodb.net/myFirstDatabase?retryWrites=true&w=majority" // string pro připojení do databáze
const dbName = "netflix" // název databáze
const colName = "watchlist" // název kolekce

//MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - helpers funkce nebudou exportovány,
// takže začínají malým písmenem

// insert 1 record - přidání nové položky, proměnná předávaná jako argument je moive se strukturou importovanou z modelu s názvem Netflix
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie) // funkce InsertOne z package collection vloží do databáze novou položku movie, má nadefinovaný Background context

	if err != nil {
		log.Fatal(err) // pokud se položku nepodařá přidat vyhodí chybu
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID) // vypíše ID vkládané položky z proměnné inserted
}

// update 1 record - upgrade stavu položky, první je třeba poskytnout filr na základě, kterého budou data odfilrovány a zjistí to, která data budeme upgradovat
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // ObjectIdFromHex převede string na ObjectID které je akceptovatelné MongoDB
	filter := bson.M{"_id": id} 
	// bson.M je pro kratší a čistší výsledky kde nezáleží na velkých a malých písmenech lze i bson.D pokud chci dostat uspořádané prvky
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	// spuštění funkce UpdateOne z kolekce, vytvoří v result hodnotu pro ModifiedCount, která může být později vypsána
	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record- odstranění jednoho záznamu
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} // filtr bude procházet všechna _id V DATABÁZI a hledat shodu s id, které jsme poskytli
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOvie got delete with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllMovie() int64 {
	//int64 je formát ve kterém bude funkce vrace return výstup
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) // místo toho, aby filtr ukládat do proměnné a tu pak dal to této funkce dal tam rovnou filtr s prázným objektem což znamená, že vezme vše
	// když nemám option nemusím na třetí pozici psát nic nebo se tam dá dát nil jako nic

	if err != nil {
		log.Fatal(err)
	}
	// funkce DeleteMany při odstranovaní spočítá kolik toho odstanila a ukládá to do deleteResult
	// počet odstraněných movies pak vypisuje do konzola
	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount // z funkce vrací počet odstraněných položek
}

// get all movies from database

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M // proměnná movies bude typu array ve kterém budou mapy typu M kde key je sting a value cokoli
	// použil for loop ve stylu while loop, smyčka neskončí dokud bude v cur nějaký objekt na který se neaplikovala
	for cur.Next(context.Background()) { //funkce Next bere z cur postupně každý dokument a vrací true dokud, vše neprojde
		var movie bson.M // vytvoří proměnnou movie formátu BSON documentu
		err := cur.Decode(&movie) // dekoduje obsah cur a uloží do movie
		// Decode = “Unmarshalling” is the process of converting some kind of a lower-level representation, often a “wire format”, into a higher-level (object) structure.
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) // přidá aktuálně dekodovaný objekt do pole BSON objektů movies
	}

	defer cur.Close(context.Background()) // uzavře spojení databází
	return movies // vrátí z funkce pole se všemi dekodovanými objekty BSON document
}

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	// mohu nastavit header tak, že nenastavuji jen content typ
	// mohu nastavit také povolenou metodu
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	 
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
