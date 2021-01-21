package bd

import (
	"context"
	"time"

	"github.com/Oscar02-94/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro en la prte final con la BD para insertar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("usuario")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
