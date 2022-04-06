package helper

import (

	"log"
	"os"
	"time"
	"gin-mongo-api/configs"
	jwt "github.com/dgrijalva/jwt-go"
 	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/twinj/uuid"
)



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
	Role 		 string
  }

var SECRET_KEY string = os.Getenv("SECRET_KEY")
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GenerateAllTokens(email string,  name string,  user_id  primitive.ObjectID , role string) (signedToken string, signedRefreshToken string, err error) {
	

	td := &TokenDetails{}
    td.AccExpires = time.Now().Add(time.Minute * 15).Unix()
    td.AccessUuid = uuid.NewV4().String()
  
    td.RefExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
    td.RefreshUuid = uuid.NewV4().String()
	


    atClaims := jwt.MapClaims{}
    atClaims["authorized"] = true
    atClaims["access_uuid"] = td.AccessUuid
    atClaims["user_id"] = user_id
    atClaims["exp"] = td.AccExpires
    at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
    token, err := at.SignedString([]byte(SECRET_KEY))



	rtClaims := jwt.MapClaims{}
    rtClaims["refresh_uuid"] = td.RefreshUuid
    rtClaims["user_id"] = user_id
    rtClaims["exp"] = td.RefExpires
    rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
    refreshToken, err := rt.SignedString([]byte(SECRET_KEY))


	
	if err != nil {
		log.Panic(err)
		return
	}

	return token , refreshToken, err
}