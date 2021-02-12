package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	*Config
	router *echo.Echo
	db     *mongo.Client
}

func NewServer(cfg *Config, db *mongo.Client) *Server {
	server := &Server{
		Config: cfg,
		router: echo.New(),
		db:     db,
	}
	// for allowing cross origin requests
	server.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// 	Design document:
	// Problem statementHigh level solutionCode implementation high level overview(data design of DB)Good to havesExtensions(like SMS)
	server.router.GET(path.Join(cfg.APIPath, "/verifyLogin"), server.verifyLogin)
	server.router.GET(path.Join(cfg.APIPath, "/userinfo"), server.userInfo)
	server.router.GET(path.Join(cfg.APIPath, "/searchjobs"), server.searchjobs)
	server.router.GET(path.Join(cfg.APIPath, "/applyjob"), server.hasappliedtojob)
	server.router.GET(path.Join(cfg.APIPath, "/savejob"), server.savejob)
	//view posted jobs
	server.router.GET(path.Join(cfg.APIPath, "/postedjobs"), server.postedjobs)
	server.router.GET(path.Join(cfg.APIPath, "/viewsavedjobs"), server.viewSavedJobs)
	server.router.POST(path.Join(cfg.APIPath, "/register"), server.registerUser)
	//create a new job
	server.router.POST(path.Join(cfg.APIPath, "/postjob"), server.postjob)
	return server
}

func (s *Server) ServeHTTP(w *httptest.ResponseRecorder, request *http.Request) {
	s.router.ServeHTTP(w, request)
}

func (server *Server) Start() error {
	address := fmt.Sprintf("%s:%d", server.Host, server.Port)
	log.Infof("listening on %s", address)
	return server.router.Start(address)
}

func (server *Server) Stop(ctx context.Context) error {
	return server.router.Shutdown(ctx)
}

func (s *Server) verifyLogin(context echo.Context) error {

	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")
	pwd := context.QueryParam("password")
	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = IsLoginValid(s, context, field, pwd, contact)
		return err
	}
	err := IsLoginValid(s, context, field, pwd, email)
	return err
}

func (s *Server) userInfo(context echo.Context) error {
	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = SendUserInfo(s, context, field, contact)
		return err
	}
	err := SendUserInfo(s, context, field, email)
	return err

}

func (s *Server) searchjobs(context echo.Context) error {
	query := context.QueryParam("query")
	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = SendJobs(s, context, query, field, contact)
		return err
	}
	err := SendJobs(s, context, query, field, email)
	return err
}

func (s *Server) hasappliedtojob(context echo.Context) error {
	jobid := context.QueryParam("job")
	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = SendApplication(s, context, jobid, field, contact)
		return err
	}
	err := SendApplication(s, context, jobid, field, email)
	return err
}

func (s *Server) savejob(context echo.Context) error {
	jobid := context.QueryParam("job")
	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = Savejob(s, context, jobid, field, contact)
		return err
	}
	err := Savejob(s, context, jobid, field, email)
	return err
}

func (s *Server) postedjobs(context echo.Context) error {
	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = SendPostedJobs(s, context, field, contact)
		return err
	}
	err := SendPostedJobs(s, context, field, email)
	return err
}

func (s *Server) viewSavedJobs(context echo.Context) error {
	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = SendSavedJobs(s, context, field, contact)
		return err
	}
	err := SendSavedJobs(s, context, field, email)
	return err
}

func Decoder(s string) []string {
	var items []string
	dec := json.NewDecoder(strings.NewReader(s))
	err := dec.Decode(&items)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return items
}

func (s *Server) registerUser(context echo.Context) error {
	skills := Decoder(context.QueryParam("skills"))
	experience := Decoder(context.QueryParam("experience"))
	contact, err := strconv.Atoi(context.QueryParam("mobile"))
	if err != nil {
		log.Fatal(err)
		errormsg := map[string]interface{}{"message": err}
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	user := &User{
		Firstname:  context.QueryParam("firstname"),
		Lastname:   context.QueryParam("lastname"),
		DOB:        context.QueryParam("dob"),
		Location:   context.QueryParam("location"),
		Gender:     context.QueryParam("gender"),
		Contact:    contact,
		Email:      context.QueryParam("email"),
		Skills:     skills,
		Experience: experience,
	}
	password := context.QueryParam("password")
	return RegisterUserDatabase(s, context, user, password)
}

func (s *Server) postjob(context echo.Context) error {

	email := context.QueryParam("email")
	mobile := context.QueryParam("mobile")

	field := "email"
	if email == "" {
		contact, err := strconv.Atoi(mobile)
		if err != nil {
			log.Fatal(err)
			return err
		}
		field = "contact"
		err = s.createjob(context, field, contact)
		return err
	}
	err := s.createjob(context, field, email)
	return err

}

func (s *Server) createjob(context echo.Context, field string, fieldval interface{}) error {

	login := GetLoginID(s, context, field, fieldval)

	contact, err := strconv.Atoi(context.QueryParam("contact"))
	if err != nil {
		log.Fatal(err)
		errormsg := map[string]interface{}{"message": err}
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}
	salary, err := strconv.Atoi(context.QueryParam("salary"))
	if err != nil {
		log.Fatal(err)
		errormsg := map[string]interface{}{"message": err}
		context.JSON(http.StatusInternalServerError, errormsg)
		return err
	}

	job := &Job{
		Companyname: context.QueryParam("companyname"),
		Title:       context.QueryParam("title"),
		Description: context.QueryParam("description"),
		Location:    context.QueryParam("location"),
		Type:        context.QueryParam("type"),
		Contact:     contact,
		Postedby:    login.Id,
		Date:        context.QueryParam("date"),
		Salary:      salary,
	}

	tokens := GetTokens(job)
	job.Tokens = tokens
	return RegisterJobDatabase(s, context, job, login)
}
