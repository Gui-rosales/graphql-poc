Queries desenvolvidas:


mutation createCategory {
  createCategory(input: {name: "Tecnologia", description: "Cursos de Tecnologia"}) {
    id
    name
    description
  }
}


query queryCategories {
   categories{
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    courses {
      id
      title
    }
  }
}


mutation createCourse {
  createCourse(input: {
    name: "JStack",
    description: "Um dos cursos já feitos",
    categoryId: "c87d0882-8fdc-4511-a4cc-ba7c7b14a813"
  }) {
    id
    title
  }
}

query queryCourses {
  courses{
    id
    title
    description
  }
}

query queryCoursesWithCaregory {
  courses {
    id
    title
    description
    category {
      id
      name
      description
    }
  }
}



Comandos rodados:

``` 
- go mod init project_name
- go mod tidy (depois de adicionar a dependência do gqlgen ao tools.go)
- go run github.com/99designs/gqlgen generate
- go run cmd/server/server.go 

```


link do repo do [gqlgen](https://gqlgen.com/)