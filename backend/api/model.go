package api

import "go.mongodb.org/mongo-driver/bson/primitive"

// The User data structures as in users table in the database
type User struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname  string             `json:"firstname" bson:"firstname"`
	Lastname   string             `json:"lastname" bson:"lastname"`
	DOB        string             `json:"dob" bson:"dob"`
	Location   string             `json:"location" bson:"location"`
	Gender     string             `json:"gender" bson:"gender"`
	Contact    int                `json:"contact" bson:"contact"`
	Email      string             `json:"email" bson:"email"`
	Skills     []string           `json:"skills" bson:"skills"`
	Experience []string           `json:"experience" bson:"experience"`
}

// A data structure to tell the application if the user is authenticated or not
type Auth struct {
	IsNewUser bool `json:"isNewUser"`
	IsAuth    bool `json:"isAuth"`
}

// The jobs collection in the database
type Job struct {
	Id          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Companyname string               `json:"companyname" bson:"companyname"`
	Title       string               `json:"title" bson:"title"`
	Location    string               `json:"location" bson:"location"`
	Contact     int                  `json:"contact" bson:"contact"`
	Description string               `json:"description" bson:"description"`
	Postedby    primitive.ObjectID   `json:"postedby" bson:"postedby"`
	Type        string               `json:"type" bson:"type"`
	Salary      int                  `json:"salary" bson:"salary"`
	Tokens      []string             `json:"tokens" bson:"tokens"`
	Date        string               `json:"date" bson:"date"`
	Applicants  []primitive.ObjectID `json:"applicants" bson:"applicants"`
}

// for login and storing the jobs the user has applied and their status
type AppliedJobs struct {
	Jobid  primitive.ObjectID `json:"jobid" bson:"jobid"`
	Status string             `json:"status" bson:"status"`
}

// login collection in database, contains all the information
type Login struct {
	Id          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Email       string               `json:"email" bson:"email"`
	Password    string               `json:"password" bson:"password"`
	Contact     int                  `json:"contact" bson:"contact"`
	Profile     primitive.ObjectID   `json:"profile" bson:"profile"`
	AppliedJobs []*AppliedJobs       `json:"appliedjobs" bson:"appliedjobs"`
	PostedJobs  []primitive.ObjectID `json:"postedjobs" bson:"postedjobs"`
	SavedJobs   []primitive.ObjectID `json:"savedjobs" bson:"savedjobs"`
}

// Differs from Job data structure in the Applicants variable
// this DS is used to send the job information along with the details of the applicants
// when the user queries for the job that has been posted by him/her
type Postjob struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Companyname string             `json:"companyname" bson:"companyname"`
	Title       string             `json:"title" bson:"title"`
	Location    string             `json:"location" bson:"location"`
	Contact     int                `json:"contact" bson:"contact"`
	Description string             `json:"description" bson:"description"`
	Postedby    primitive.ObjectID `json:"postedby" bson:"postedby"`
	Type        string             `json:"type" bson:"type"`
	Salary      int                `json:"salary" bson:"salary"`
	Tokens      []string           `json:"tokens" bson:"tokens"`
	Date        string             `json:"date" bson:"date"`
	Applicants  []*User            `json:"applicants" bson:"applicants"`
}

// When the user searches for All Jobs or Saved Jobs this
// will tell the status of all the jobs wrt the user
type Userjob struct {
	Job       *Job `json:"job" bson:"job"`
	IsSaved   bool `json:"issaved" bson:"issaved"`
	IsApplied bool `json:"isapplied" bson:"isapplied"`
}
