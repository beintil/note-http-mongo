package note

import (
	"beintil/mongo-http/domain/note"
	"beintil/mongo-http/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	content     = "content-type"
	application = "application/json"
)

var (
	collection = database.ConnectDB()
	ctx        = context.TODO()
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, application)

	var note note.CreateNoteDTD
	note.UpdateAt, note.CreateAt = time.Now(), time.Now()

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	// Получаем наш notes model
	result, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, application)
	var notes []note.AllNoteDTD

	curser, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	if err = curser.All(ctx, &notes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(notes)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, application)

	var vars = mux.Vars(r)

	// Get id from notes
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	var note note.UpdateNoteDTD
	note.UpdateAt = time.Now()

	var notes_id = bson.M{"_id": id}

	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "notesname", Value: note.Name},
			{Key: "description", Value: note.Description},
		}},
	}

	err = collection.FindOneAndUpdate(ctx, notes_id, update).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	note.UUID = id

	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, application)

	var vars = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	notes_id := bson.M{"_id": id}

	deleteNote, err := collection.DeleteOne(ctx, notes_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"}` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(deleteNote)
}
