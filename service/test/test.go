package test

import (
	"context"
	"fmt"

	"github.com/phungvandat/onemilion/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type testService struct {
	mongoDB *mongo.Database
}

// NewTestService func
func NewTestService(mongoDB *mongo.Database) Service {
	return &testService{
		mongoDB: mongoDB,
	}
}

func (w *testService) Test(input domain.Payload) error {
	ctx := context.TODO()

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"flag": bson.M{
					"$ne": -1,
				},
				"$or": []bson.M{
					bson.M{
						"visible": bson.M{
							"$exists": false,
						},
					},
					bson.M{
						"visible": true,
					},
				},
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "treatmentpackages",
				"localField":   "treatmentPackages",
				"foreignField": "_id",
				"as":           "treatmentPackages",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user",
				"foreignField": "_id",
				"as":           "user",
			},
		},
		bson.M{
			"$addFields": bson.M{
				"user": bson.M{
					"$arrayElemAt": []interface{}{"$user", 0},
				},
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "clinics",
				"localField":   "clinic",
				"foreignField": "_id",
				"as":           "clinic",
			},
		},
		bson.M{
			"$addFields": bson.M{
				"clinic": bson.M{
					"$arrayElemAt": []interface{}{"$clinic", 0},
				},
			},
		},
		bson.M{
			"$match": bson.M{
				"clinic.visibleOnFeed": bson.M{
					"$ne": false,
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"createdAt": -1,
			},
		},
		bson.M{
			"$skip": 0,
		},
		bson.M{
			"$limit": 100,
		},
	}

	cur, err := w.mongoDB.Collection("posts").Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	for cur.Next(ctx) {
		_, err := cur.Current.Values()
		if err != nil {
			return err
		}
		postID, _ := primitive.ObjectIDFromHex("5db8121ccb22934406877116")
		cs, err := w.mongoDB.Collection("likes").Find(ctx, bson.M{
			"post": postID,
		})
		for cs.Next(ctx) {
			cs.Current.Values()
		}
		if err != nil {
			return err
		}

	}

	arrPostIDs := []primitive.ObjectID{}
	postID, _ := primitive.ObjectIDFromHex("5d04eb6b6fd3ab314c63810e")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cf90e247a276c3f6de18b44")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d04bfb76fd3ab314c63737f")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cf8c6c87a276c3f6de175d7")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfe8d3c6fe64e5bb1161a97")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cf5d6a2a176de7df244e131")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce9f9b1708aa67603ff5caf")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce9f304708aa67603ff591a")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cef46cf84769d2733299480")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0533326fd3ab314c6394d1")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce79115de6a3a669530b705")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce9fc1d708aa67603ff5dd8")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce97572708aa67603ff51af")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce669c07a930353c31c2e2a")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce65e3d7a930353c31c284e")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cffe0de62d31a3b3406afd0")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d00dfff62d31a3b3407145f")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce8b46d708aa67603ff133f")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d04ee3b6fd3ab314c63841f")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ced16de84769d2733292142")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce791a1de6a3a669530b718")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce65def7a930353c31c2829")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d034429bd1fff5aa5cef539")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cf6a63fa176de7df2455383")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce911a1708aa67603ff3a6a")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce906d9708aa67603ff3476")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfe99446fe64e5bb1161af2")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce90fe0708aa67603ff39ba")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0450b16fd3ab314c630462")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce90f79708aa67603ff3991")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfa8629dc42da034530c9b9")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5ce8f5fd708aa67603ff2cc0")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cffdf0c62d31a3b3406af8f")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0451676fd3ab314c630501")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfb4d94dc42da034530f4c9")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfbd41cdc42da03453119f2")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d04e8e46fd3ab314c637fdf")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfe8cd16fe64e5bb1161a96")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d034098bd1fff5aa5cef371")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cffb15662d31a3b3406a3bb")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0346d9bd1fff5aa5cef6f8")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0394536fd3ab314c62e0fd")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfba4eddc42da0345310f13")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cfe8d3e6fe64e5bb1161a98")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0076f462d31a3b3406d99a")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d01c3a24351fa0d0e5a38d6")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5cede9be84769d2733293dfa")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d03813f86a84929dbb1cd91")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d042f6f6fd3ab314c62f508")
	arrPostIDs = append(arrPostIDs, postID)
	postID, _ = primitive.ObjectIDFromHex("5d0450466fd3ab314c6303e7")
	arrPostIDs = append(arrPostIDs, postID)

	userID, _ := primitive.ObjectIDFromHex("5ce3a0e56229185c9fb9ebcd")
	w.mongoDB.Collection("likes").Find(ctx, bson.M{
		"user": userID,
		"post": bson.M{
			"$in": arrPostIDs,
		},
	})

	fmt.Println("Done work: ", input.Num)
	return nil
}
