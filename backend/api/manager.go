package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Helper Functions for API Handlers

// Check if the user exists in the system
// If the user exists check if he has entered the correct password or not and return isAuth true or false depending on taht
// If the user does not exist isNewUser is true and the user is redirected to registration page
func IsLoginValid(s *Server, context echo.Context, field string, password string, fieldval interface{}) error {

	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	errormsg := map[string]interface{}{"message": "error"}
	cursor, err := collection.Find(ctx, bson.M{field: fieldval})
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	if len(users) == 1 {
		for _, user := range users {
			if user["password"] == password {
				context.JSON(http.StatusOK, Auth{false, true})
			} else {
				context.JSON(http.StatusOK, Auth{false, false})
			}
		}
	} else {
		context.JSON(http.StatusOK, Auth{true, true})
	}
	return nil
}

// Send user information based on email or mobile number
func SendUserInfo(s *Server, context echo.Context, field string, fieldval interface{}) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	errormsg := map[string]interface{}{"message": "error"}
	cursor, err := collection.Find(ctx, bson.M{field: fieldval})
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	var logins []bson.M
	if err = cursor.All(ctx, &logins); err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	for _, login := range logins {
		collection := s.db.Database("jobBazar").Collection("users")
		var user User
		err = collection.FindOne(ctx, bson.M{"_id": login["profile"]}).Decode(&user)
		if err != nil {
			log.Fatal(err)
			context.JSON(http.StatusInternalServerError, errormsg)
			return err
		}
		context.JSON(http.StatusOK, user)
		return nil
	}
	return nil
}

// Send the user all the job details based on the search query along with the
// information whether he has already saved or applied for the job before hand
func SendJobs(s *Server, context echo.Context, query string, field string, fieldval interface{}) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("jobs")
	errormsg := map[string]interface{}{"message": "error"}
	if query == "" {
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
			context.JSON(http.StatusInternalServerError, errormsg)
			return err
		}
		var login Login
		collection = s.db.Database("jobBazar").Collection("login")
		err = collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
		var response []*Userjob
		for cursor.Next(ctx) {
			var j *Job
			err = cursor.Decode(&j)
			if err != nil {
				log.Fatal(err)
				context.JSON(http.StatusInternalServerError, errormsg)
				return err
			}
			uj := &Userjob{
				Job:       j,
				IsSaved:   false,
				IsApplied: false}
			for _, app := range login.AppliedJobs {
				if app.Jobid.Hex() == j.Id.Hex() {
					uj.IsApplied = true
				}
			}
			for _, app := range login.SavedJobs {
				if app.Hex() == j.Id.Hex() {
					uj.IsSaved = true
				}
			}
			response = append(response, uj)
		}
		context.JSON(http.StatusOK, response)
	} else {
		cursor, err := collection.Find(ctx, bson.D{{"tokens", bson.D{{"$all", bson.A{query}}}}})
		if err != nil {
			log.Fatal(err)
			context.JSON(http.StatusInternalServerError, errormsg)
			return err
		}
		var login Login
		collection = s.db.Database("jobBazar").Collection("login")
		err = collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
		var response []*Userjob
		for cursor.Next(ctx) {
			var j *Job
			err = cursor.Decode(&j)
			if err != nil {
				log.Fatal(err)
				context.JSON(http.StatusInternalServerError, errormsg)
				return err
			}
			uj := &Userjob{
				Job:       j,
				IsSaved:   false,
				IsApplied: false}
			for _, app := range login.AppliedJobs {
				if app.Jobid.Hex() == j.Id.Hex() {
					uj.IsApplied = true
				}
			}
			for _, app := range login.SavedJobs {
				if app.Hex() == j.Id.Hex() {
					uj.IsSaved = true
				}
			}
			response = append(response, uj)
		}
		context.JSON(http.StatusOK, response)
	}
	return nil
}

// Send the user the details of the job saved by him and the details whether he has applied
// for them or not
func SendSavedJobs(s *Server, context echo.Context, field string, fieldval interface{}) error {
	ctx := context.Request().Context()
	var login Login
	collection := s.db.Database("jobBazar").Collection("login")
	errormsg := map[string]interface{}{"message": "error"}
	err := collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}

	var response []*Userjob
	savedjobs := login.SavedJobs
	appliedjobs := login.AppliedJobs
	for _, jobid := range savedjobs {
		var job Job
		collection = s.db.Database("jobBazar").Collection("jobs")
		err = collection.FindOne(ctx, bson.M{"_id": jobid}).Decode(&job)
		if err != nil {
			log.Fatal(err)
			context.JSON(http.StatusInternalServerError, errormsg)
			return err
		}

		userjob := &Userjob{&job, true, false}
		for _, app := range appliedjobs {
			if app.Jobid.Hex() == jobid.Hex() {
				userjob.IsApplied = true
			}
		}

		response = append(response, userjob)
	}
	context.JSON(http.StatusOK, response)
	return nil
}

// When a user applies for a job the function takes the jobid and the user details
// add the job to the user's appliedjobs list
//and add user to the applicants list of the job - this part is done by the function AddUserToJob
func SendApplication(s *Server, context echo.Context, jobid, field string, fieldval interface{}) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	errormsg := map[string]interface{}{"message": "error"}
	ispresent := map[string]interface{}{"ispresent": true}
	isnotPresent := map[string]interface{}{"ispresent": false}
	var login Login
	err := collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	jid, _ := primitive.ObjectIDFromHex(jobid)
	for _, j := range login.AppliedJobs {
		if j.Jobid.Hex() == jid.Hex() {
			context.JSON(http.StatusOK, ispresent)
			return nil
		}
	}
	context.JSON(http.StatusOK, isnotPresent)
	newApplied := login.AppliedJobs
	newjob := &AppliedJobs{jid, "applied"}
	newApplied = append(newApplied, newjob)
	_, err = collection.UpdateOne(
		ctx,
		bson.M{field: fieldval},
		bson.D{
			{"$set", bson.D{{"appliedjobs", newApplied}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	AddUserToJob(s, context, jid, login.Id)
	return nil
}

// add user to the applicants list of the job
func AddUserToJob(s *Server, context echo.Context, jobid, loginid primitive.ObjectID) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("jobs")
	var job Job
	err := collection.FindOne(ctx, bson.M{"_id": jobid}).Decode(&job)
	if err != nil {
		log.Fatal(err)
		return err
	}

	applicants := append(job.Applicants, loginid)
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": jobid},
		bson.D{
			{"$set", bson.D{{"applicants", applicants}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

// Add a jobid to the savedjobs field of the user
func Savejob(s *Server, context echo.Context, jobid, field string, fieldval interface{}) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	errormsg := map[string]interface{}{"message": "error"}
	var login Login
	err := collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	success := map[string]interface{}{"message": "success"}
	jid, _ := primitive.ObjectIDFromHex(jobid)
	for _, j := range login.SavedJobs {
		if j.Hex() == jid.Hex() {
			context.JSON(http.StatusOK, success)
			return nil
		}
	}
	newsaved := login.SavedJobs
	newsaved = append(newsaved, jid)
	_, err = collection.UpdateOne(
		ctx,
		bson.M{field: fieldval},
		bson.D{
			{"$set", bson.D{{"savedjobs", newsaved}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}

	context.JSON(http.StatusOK, success)
	return nil
}

// Send the jobs posted by the user along with the user details of applicants
func SendPostedJobs(s *Server, context echo.Context, field string, fieldval interface{}) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	errormsg := map[string]interface{}{"message": "error"}
	var login Login
	err := collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}

	var jobs []*Postjob
	for _, jobid := range login.PostedJobs {
		jobs = append(jobs, GetJobDetails(s, context, jobid))
	}
	context.JSON(http.StatusOK, jobs)
	return nil
}

// Helper Function - SendPostedJobs
// Get the job details for a single job with the applicant details
func GetJobDetails(s *Server, context echo.Context, jobid primitive.ObjectID) *Postjob {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("jobs")
	// errormsg := map[string]interface{}{"message": "error"}
	var job Job
	err := collection.FindOne(ctx, bson.M{"_id": jobid}).Decode(&job)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	postjob := CreateNewPostjob(job)
	var users []*User
	for _, appid := range job.Applicants {
		users = append(users, GetUserDetails(s, context, appid))
	}
	postjob.Applicants = users
	return postjob
}

// Helper Function - GetJobDetails
// Create a new Postjob object with the passed job data structure
func CreateNewPostjob(job Job) *Postjob {
	var applicants []*User
	return &Postjob{
		Id:          job.Id,
		Companyname: job.Companyname,
		Title:       job.Title,
		Location:    job.Location,
		Contact:     job.Contact,
		Description: job.Description,
		Postedby:    job.Postedby,
		Type:        job.Type,
		Salary:      job.Salary,
		Tokens:      job.Tokens,
		Date:        job.Date,
		Applicants:  applicants,
	}
}

// Helper Function - GetJobDetails
// With the ObjectId of the login, return the user details
func GetUserDetails(s *Server, context echo.Context, loginid primitive.ObjectID) *User {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	// errormsg := map[string]interface{}{"message": "error"}
	var login *Login
	err := collection.FindOne(ctx, bson.M{"_id": loginid}).Decode(&login)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	collection = s.db.Database("jobBazar").Collection("users")
	var user *User
	err = collection.FindOne(ctx, bson.M{"_id": login.Profile}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return user
}

// Create a new user in the database
// A new document in the users collection is created
// A new document in the login collection is created
func RegisterUserDatabase(s *Server, context echo.Context, user *User, password string) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("users")
	errormsg := map[string]interface{}{"message": "error"}
	profile, err := collection.InsertOne(ctx, user)
	var userid primitive.ObjectID
	userid = profile.InsertedID.(primitive.ObjectID)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}

	collection = s.db.Database("jobBazar").Collection("login")
	login := &Login{
		Email:    user.Email,
		Password: password,
		Contact:  user.Contact,
		Profile:  userid,
	}
	_, err = collection.InsertOne(ctx, login)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	context.JSON(http.StatusOK, map[string]interface{}{"message": "Success"})
	return nil
}

// Helper Function - createjob function in user.go
// get the login object from email id or mobile number
func GetLoginID(s *Server, context echo.Context, field string, fieldval interface{}) *Login {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("login")
	var login Login
	err := collection.FindOne(ctx, bson.M{field: fieldval}).Decode(&login)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &login
}

// Add a new job to the database
func RegisterJobDatabase(s *Server, context echo.Context, job *Job, login *Login) error {
	ctx := context.Request().Context()
	collection := s.db.Database("jobBazar").Collection("jobs")
	insertedId, err := collection.InsertOne(ctx, job)
	errormsg := map[string]interface{}{"message": "error"}
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	jobid := insertedId.InsertedID.(primitive.ObjectID)
	collection = s.db.Database("jobBazar").Collection("login")
	postedjobs := append(login.PostedJobs, jobid)
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": login.Id},
		bson.D{
			{"$set", bson.D{{"postedjobs", postedjobs}}},
		},
	)

	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	context.JSON(http.StatusOK, map[string]interface{}{"message": "Success"})
	return nil
}
