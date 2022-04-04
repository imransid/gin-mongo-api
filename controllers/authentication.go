



package controllers

//     "fmt"

import (
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
   // "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
   "golang.org/x/crypto/bcrypt"
 // "reflect"
  
)

// reflect.TypeOf()

var authCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var authValidate = validator.New()




func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func Authentication() gin.HandlerFunc {
  
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var auth models.Auth

        defer cancel()
     
               // //validate the request body
            if err := c.BindJSON(&auth); err != nil {
                c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
    
            //use the validator library to validate required fields
            if validationErr := validate.Struct(&auth); validationErr != nil {
                c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
                return
            }


            var password string = auth.Password

            var result bson.M

            usernameErr := userCollection.FindOne(ctx, bson.M{"name": auth.Username}).Decode(&result)

            emailErr := userCollection.FindOne(ctx, bson.M{"email": auth.Username}).Decode(&result)

            

        if usernameErr != nil && emailErr != nil {

            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Invalid username"}})
            return
        }

        // if str, ok := result["password"].(string); ok {

        //    return result["password"]
        //     /* act on str */
        // } else {
        //    return ""
        //     /* not string */
        // }

        hashkey := fmt.Sprint(result["password"])

        checkPassword :=  CheckPasswordHash(password, hashkey)


        if checkPassword {
            c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
        }else{
            c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "Failed", Data: map[string]interface{}{"data": "Invalid Password"}})
        }


        fmt.Println("checkPassword" , checkPassword)


        // return data
       
    }


}