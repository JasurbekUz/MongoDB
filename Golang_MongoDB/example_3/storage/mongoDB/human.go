package mongoDB

import (
	"context"
	"time"

	"monGo/models"
	"monGo/pkg/utils"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *humanRepo) Create (human *models.Human) (*models.Human, *utils.RestErr) {
	dbHumans := r.db.Collection("humans")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := dbHumans.InsertOne(ctx, bson.M{
		"fullname":     human.Firstname,
		"lastname":    human.Lastname,
		"age": human.Age,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert human to the database.")
		return nil, restErr
	}
	human.ID = result.InsertedID.(primitive.ObjectID)

	return human, nil
}

func (r *humanRepo) Find(ID primitive.ObjectID) (*models.Human, *utils.RestErr) {
	var human models.Human
	usersC := r.db.Collection("humans")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"ID": ID}).Decode(&human)
	if err != nil {
		restErr := utils.NotFound("human not found.")
		return nil, restErr
	}
	return &human, nil
}

func (r *humanRepo) Update(human *models.Human) (*models.Human, *utils.RestErr) {
	humans := r.db.Collection("humans")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := humans.UpdateOne(
		ctx, 
		bson.M{"ID": human.ID}, 
		bson.M{"$set": bson.M{
			"firstname": human.Firstname, 
			"lastname": human.Lastname,
			"age": human.Age,
			},
		},
	)

	/*result, err := humans.UpdateOne(
	    ctx,
	    bson.M{"_id": human.ID},
	    bson.D{
	        {"$set", bson.D{{human.Firstname, human.Lastname, human.Age}}},
	    },
	)
*/
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("human not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	human, restErr := r.Find(human.ID)
	if restErr != nil {
		return nil, restErr
	}
	return human, restErr
}

func (r *humanRepo) Delete(ID primitive.ObjectID) *utils.RestErr {
	usersC := r.db.Collection("humans")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		restErr := utils.NotFound("faild to delete.")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := utils.NotFound("human not found.")
		return restErr
	}
	return nil
}

