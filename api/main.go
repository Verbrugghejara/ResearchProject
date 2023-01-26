package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

    "os"
	// "errors"
	"context"

	"github.com/gofiber/fiber/v2"

	// "strconv"
	// "reflect"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	// "github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Scroll struct {
	Data int
	Id   primitive.ObjectID `bson:"_id"`
}
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  } 
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	// dotenv := goDotEnvVariable("mongoDB_password")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jaraver:Jonasjara1@research.9wqbygj.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1000000000*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	app := fiber.New()
	app.Use(cors.New())
	var ConfigDefault = cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	// Or extend your config for customization
	app.Use(cors.New(ConfigDefault))

	app.Get("/api/:id", func(c *fiber.Ctx) error {

		objID, err := primitive.ObjectIDFromHex(c.Params("id"))

		if err != nil {
			return c.Status(400).SendString(string(err.Error()))

		}
	

		// 	// Give the database collection
		aDatabase := client.Database("api")
		theCollection := aDatabase.Collection("Scroll")

		// 	// get the data from the database
		// filter, err := theCollection.Find(ctx, bson.M{"_id" :objID})
		filter := bson.D{{"_id", objID}}
		var result Scroll
		err = theCollection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				return c.Status(400).SendString("your query did not match any documents")
			}
			panic(err)
		}
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println("result")
		fmt.Println(string(res))

		return c.JSON(result)

	})


	app.Put("/api/update/:id", func(c *fiber.Ctx) error {
		scroll := &Scroll{}
		if err := c.BodyParser(scroll); err != nil {
			return err
		}

		objID, err := primitive.ObjectIDFromHex(c.Params("id"))
		fmt.Println(scroll.Data)

		if err != nil {
			return c.Status(400).SendString(string(err.Error()))

		}

		// 	// Give the database collection
		aDatabase := client.Database("api")
		theCollection := aDatabase.Collection("Scroll")

		results, errdb := theCollection.UpdateOne(ctx, bson.M{ // this is the
			"_id": objID,
		}, bson.D{
			{"$set", bson.D{
				{"Data", scroll.Data}}},
		})

		if errdb != nil {
			log.Print("errdb")
			log.Print(errdb)
			// app.Connect(":3333")

			// log.Fatal(app.Listen(":3333"))
		}

		return c.JSON(results)

	})
	log.Fatal(app.Listen(":3333"))
	// Start the server
	// http.HandleFunc("/api/HttpExample", updateData)
	// http.HandleFunc("/api/HttpExample2", getData)
	// data :=1080
	// response := "/api/update" + strconv.Itoa(data)

	// err := http.ListenAndServe(":3333", nil)
	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// 	} else if err != nil {
	// 		fmt.Printf("error starting server: %s\n", err)
	// 		os.Exit(1)
	// 	}
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jaraver:<password>@research.9wqbygj.mongodb.net/test"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// aDatabase := client.Database("api")
	// theCollection := aDatabase.Collection("Scroll")

	// insertResult, err := theCollection.InsertOne(ctx, bson.D{
	// 	{"Data", 0},
	//    })

	//    if err != nil {
	// 	log.Println("There was an errr in trying to migrate the data into the database")
	//    }

	//    fmt.Println(insertResult.InsertedID)
	// objID, _ := primitive.ObjectIDFromHex("63c161b3f92ef4f1658fa6af")
	// // filter, err := theCollection.Find(ctx, bson.M{"_id" :objID})
	// filter := bson.D{{"_id", objID}}
	// var result Music
	// err = theCollection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		// This error means your query did not match any documents.
	// 		return
	// 	}
	// 	panic(err)
	// }
	// res, _ := bson.MarshalExtJSON(result, false, false)
	// fmt.Println(string(res))

	// fileCount := res
	// fmt.Println(string(fileCount))

}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fileCount := map[string]string{
// 		"message": "Hello World",
//     }
//     bytes, _ := json.Marshal(fileCount)
//     // fmt.Println(string(bytes))
// 	fmt.Fprint(w, string(bytes))
// }

// func getData(w http.ResponseWriter, r *http.Request){
//     // if origin := r.Header.Get("Origin"); origin != "" {
//     //     w.Header().Set("Access-Control-Allow-Origin", "*")
//     // }

// 	enableCors(&w)
// // connect with database
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jaraver:<password>@research.9wqbygj.mongodb.net/test"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Give the database collection
// 	aDatabase := client.Database("api")
// 	theCollection := aDatabase.Collection("Scroll")

// 	// get the data from the database
// 	objID, _ := primitive.ObjectIDFromHex("63c190b806a8b3cd9276ed4b")
// 	// filter, err := theCollection.Find(ctx, bson.M{"_id" :objID})
// 	filter := bson.D{{"_id", objID}}
// 	var result Scroll
// 	err = theCollection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// This error means your query did not match any documents.
// 			return
// 		}
// 		panic(err)
// 	}
// 	res, _ := bson.MarshalExtJSON(result, false, false)
//     fmt.Println(string(res))

// 	// return the data

// 	fileCount := res

//     fmt.Println(string(fileCount))
// 	fmt.Fprint(w, string(fileCount))

// }

// func updateData(w http.ResponseWriter, r *http.Request){

//     // w.Header().Set("Access-Control-Allow-Origin", "http://http://localhost:5174")
//     // if origin := r.Header.Get("Origin"); origin != "" {
//     //     w.Header().Set("Access-Control-Allow-Origin", "*")
//     // }

// 	scrollData := r.URL.Query().Get("scrollData")
// 	fmt.Println(string("updateData"+scrollData))
// 	marks, err := strconv.Atoi(scrollData)

// 	if err != nil {
// 		fmt.Println("Error during conversion")
// 		return
// 	}
// 	fmt.Println(reflect.TypeOf(marks))
// 	// connect with database
// 	enableCors(&w)
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jaraver:<password>@research.9wqbygj.mongodb.net/test"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Give the database collection
// 	aDatabase := client.Database("api")
// 	theCollection := aDatabase.Collection("Scroll")

// 	// Change the data in the database
// 	objID, _ := primitive.ObjectIDFromHex("63c190b806a8b3cd9276ed4b")

// 	results, err := theCollection.UpdateOne(ctx, bson.M{ // this is the
// 		"_id": objID,
// 	   }, bson.D{
// 		{"$set", bson.D{
// 		 {"Data", marks}}},
// 	   })

// 	   if err != nil {
// 		log.Print(err)
// 	   }

// 	  fmt.Println(results.ModifiedCount,"data",marks)

// 	fmt.Fprint(w, marks)
// }
