package controllers

//     "fmt"

import (
    "log"
    "fmt"
    "context"
    "gin-mongo-api/configs"
    "gin-mongo-api/models"
    "gin-mongo-api/responses"
    "net/http"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    helper "gin-mongo-api/helpers"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}



func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

       // //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }


        // check email exists 

        count, err := userCollection.CountDocuments(ctx, bson.M{"email" : user.Email})

        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error" : "error ocurred while checking email"})
        }

        if count > 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error" : "Email or User already exits! "})
            return
        }


        user_id := primitive.NewObjectID()


        hash, _ := HashPassword(user.Password) // ignore error for the sake of simplicity
        time_at , _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))


        token, refreshToken , _ := helper.GenerateAllTokens(user.Email, user.Name, user_id, user.Role)

        

        newUser := models.User{
            Id:       user_id,
            Name:     user.Name,
            Location: user.Location,
            Title:    user.Title,
            Email:    user.Email,
            Password: hash,
            Role :    user.Role,
            Phone:    user.Phone,
            Created_at : time_at,
            Updated_at : time_at,
            Token : token,
            Refresh_token : refreshToken,  
            FirstName : user.FirstName,
            LastName : user.LastName,
        }  

      
        result, err := userCollection.InsertOne(ctx, newUser)


        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

         c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}


func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {

        userId := c.Param("userId")


fmt.Println(time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)))

        if err := helper.MatchUserTypeToUid(c, userId); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
        }


        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
    }
}


func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        var user models.User
        defer cancel()
        objId, _ := primitive.ObjectIDFromHex(userId)

        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
        result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //get updated user details
        var updatedUser models.User
        if result.MatchedCount == 1 {
            err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }

        c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
    }
}

func DeleteAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound,
                responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
            )
            return
        }

        c.JSON(http.StatusOK,
            responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
        )
    }
}

func GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context) {







        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

         
        var users []models.User
        defer cancel()

        results, err := userCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleUser models.User
            if err = results.Decode(&singleUser); err != nil {
                c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
          
            users = append(users, singleUser)
        }

        c.JSON(http.StatusOK,
            responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
        )
    }
}

     
 