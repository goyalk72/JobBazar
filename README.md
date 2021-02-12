# JobBazar

JobBazar is a web application for users to find and post blue-collar jobs.

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)

## General info
A simple web portal for listing, searching and applying to jobs. The app supports login using email or mobile number (as some workers might not have an email id). The app is kept the same for job seekers as well as recruiters. After logging in, a worker can see his profile including name, location, skills, experience and other relevant details. The search job feature is kept simple as well wherein, you can enter a keyword which can be the location, company name, job type or job description and you can view the jobs accordingly. The app works for recruiters as well wherein they can view the relevant details of the applicants for every job as well.

## Technologies
* ReactJS version: 17.0.1
* Go version: 1.14.4
* MongoDB Atlas version: 4.4.3

## Design
The design for backend and database is included in the file /backend/api/model.go

## Setup
Clone the repository and follow these steps.
```bash
cd /jobBazar
cd /frontend
npm install
```

## Run
* Running the server
```bash
cd /backend
go run main.go
```
* Running the web app
```bash
cd /frontend
npm start
```

## Features
* A simple UI with support for registering, job searching and applying jobs
* Logging in with either or both ,email and phone supported
* Search using any keyword i.e. location, job title, company name, description keywords is supported
* Keywords for each job are displayed by which the job can be searched
 #### To do
* A recommendation feature to list jobs that are best suited for the user. (Using NLP or external APIs)
* An SMS service for people who do not have access to internet.
* Adding support for daily wage jobs.
* Adding all the other basic features of a job portal which could not be added due to the time constraint
