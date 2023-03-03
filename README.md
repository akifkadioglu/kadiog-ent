# Hello to everyone!
I made this setup to create the folder structure that I use while developing each backend application.
thank you

- https://echo.labstack.com/
- https://entgo.io/


# Commands

- a &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; make create schema=YourSchema 
- b &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; make generate
- c &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; make run
- d &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; make describe


# Localization

The Kadiog initially comes with localization package
If you want add your own messages, add your json to ./resources/Language and add keys to Keys.go as in the example


# Environments

For see environments go ./env/env.go


# Database

The Kadiog use https://entgo.io/ for database and it is ready to SQLite, MySQL and PostgreSQL.. 
- if you want to create new schema, run 'command a' sure change 'YourSchema' :D
- if you want to genererate schemas functions, run 'command b'
- as you see in Idatabase main function set your database


# Routes

The Kadiog use https://echo.labstack.com/ for rest api

- check ./routes/init.go for set middleware or something
- you can add your endpoints to ./routes/routes.go 