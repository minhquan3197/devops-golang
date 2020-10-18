package convert

import (
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StringToInt func parse string to int
func StringToInt(payload string) (int, error) {
	var result int
	result, err := strconv.Atoi(payload)
	if err != nil {
		return result, errors.New("Parse failed")
	}
	return result, nil
}

// StringToInt64 func parse string to int64
func StringToInt64(payload string) (int64, error) {
	var result int64
	result, err := strconv.ParseInt(payload, 0, 64)
	if err != nil {
		return result, errors.New("Parse failed")
	}
	return result, nil
}

// StringToObjectID func parse string to objectID
func StringToObjectID(payload string) (primitive.ObjectID, error) {
	var result primitive.ObjectID
	result, err := primitive.ObjectIDFromHex(payload)
	if err != nil {
		return result, errors.New("Parse failed")
	}
	return result, nil
}
