#!/bin/bash

module_name=$1

echo $module_name

mkdir -p "src/$module_name"

touch src/$module_name/model.go
touch src/$module_name/repository.go
touch src/$module_name/service.go
touch src/$module_name/routes.go

sudo chmod -R a+w src/$module_name

echo "package $module_name" > src/$module_name/model.go 

printf "package $module_name

import (
  \"database/sql\"
)

type Repository interface {}

type repo struct {
  db *sql.DB
}

func NewRepo(DB *sql.DB) Repository {
  return &repo{db: DB}
}" > src/$module_name/repository.go

printf "package $module_name

type Service interface {}

type service struct {
  repository Repository
}

func NewService(rep Repository) Service {
  return &service{repository: rep}
}" > src/$module_name/service.go

printf "package $module_name

import (
  \"context\"

  \"github.com/gorilla/mux\"
)

func SetRoutes(ctx context.Context, r *mux.Router, ${module_name}Service Service) {
}" > src/$module_name/routes.go
