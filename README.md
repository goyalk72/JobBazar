# JobBazar

JobBazar is a web application for users to find and post blue-collar jobs.

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)

## General info
A simple web portal for listing, searching and applying to jobs.

## Technologies
* ReactJS version: 17.0.1
* Go version: 1.14.4
* MongoDB Atlas version: 4.4.3

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
 #### To do
* A recommendation feature to list jobs that are best suited for the user. (Using NLP or external APIs)
* An SMS service for people who do not have access to internet.
* Adding support for daily wage jobs.