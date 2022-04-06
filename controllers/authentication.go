



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
   "github.com/dgrijalva/jwt-go"
   "os"
   "github.com/twinj/uuid"

 // "reflect"
  
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenDetails struct {
    AccessToken  string
    RefreshToken string
    AccessUuid   string
    RefreshUuid  string
    AccExpires    int64
    RefExpires    int64
    username     string
    email        string
    userid       string
  }
  
  

// reflect.TypeOf()

var authCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var authValidate = validator.New()

var jwtKey = []byte("RaFa1234@@secret")




func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreateToken(userid string) (*TokenDetails, error) {

    td := &TokenDetails{}
    td.AccExpires = time.Now().Add(time.Minute * 15).Unix()
    td.AccessUuid = uuid.NewV4().String()
  
    td.RefExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
    td.RefreshUuid = uuid.NewV4().String()

    td.username = userid

    var err error
    //Creating Access Token
    os.Setenv("ACCESS_SECRET", "raFa1234@@&&06") //this should be in an env file
    atClaims := jwt.MapClaims{}
    atClaims["authorized"] = true
    atClaims["access_uuid"] = td.AccessUuid
    atClaims["user_id"] = userid
    atClaims["exp"] = td.AccExpires
    at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
    td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
    if err != nil {
       return nil, err
    }
    //Creating Refresh Token
    os.Setenv("REFRESH_SECRET", "RaFa2341@@&&60") //this should be in an env file
    rtClaims := jwt.MapClaims{}
    rtClaims["refresh_uuid"] = td.RefreshUuid
    rtClaims["user_id"] = userid
    rtClaims["exp"] = td.RefExpires
    rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
    td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
    
    if err != nil {
       return nil, err
    }  
    return td, nil
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

        hashkey := fmt.Sprint(result["password"])
        

        checkPassword :=  CheckPasswordHash(password, hashkey)
      

        if checkPassword {



            finaleToken, err := CreateToken("email")

                var token map[string]interface{}
                /* create a map*/
                token = make(map[string]interface{})
                
                /* insert key-value pairs in the map*/
                token["AccessToken"] = finaleToken.AccessToken
                token["RefreshToken"] = finaleToken.RefreshToken
                token["AccessUuid"] = finaleToken.AccessUuid
                token["RefreshUuid"] = finaleToken.RefreshUuid
                token["AccExpires"] = finaleToken.AccExpires
                token["RefExpires"] = finaleToken.RefExpires
                token["username"] = "Tokyo"
                token["email"] = "New Delhi"
                token["userid"] = "New Delhi"

           // &finaleToken.email = "example@.com"

            fmt.Println("token", token, finaleToken)

            if err != nil {
                c.JSON(http.StatusUnprocessableEntity, err.Error())
                return
            }


         
            // fmt.Println(claims, tokenString)

            c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": token}})
        }else{
            c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "Failed", Data: map[string]interface{}{"data": "Invalid Password"}})
        }


        fmt.Println("checkPassword" , checkPassword)


        // return data
       
    }


}