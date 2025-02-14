# Проект: Система для публикации постов и комментариев

Этот проект представляет собой систему для добавления и чтения постов и комментариев с использованием GraphQL.

## Запуск проекта

1. ```bash
    git clone https://github.com/f7rzen/post-system.git
    cd post-system
    ```
2. Создайте файл `.env` с переменными окружения для базы данных:
    ```env
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    DB_HOST=your_db_host
    DB_PORT=yuor_db_port
    PGADMIN_DEFAULT_EMAIL=your_pgadmin_email
    PGADMIN_DEFAULT_PASSWORD=your_pgadmin_password
    ```
3. Установите зависимости с помощью команды:
    ```bash
    go mod tidy
    ```
4. Запустите проект с помощью команды:
    ```bash
    docker-compose up --build -d
    ```
5. API будет доступно по адресу http://localhost:8080/query
## Описание запросов

##### Создание поста


```graphql
mutation {
  createPost(title: "Новый пост", content: "Это контент поста", authorID: "1") {
    id
    title
    content
  }
}

##### Создание родительского комментария


```mutation {
  createParentComment(postID: "5", authorID: "2", content: "Это родительский комментарий2!") {
    id
    content
  }
}


##### Создание дочернего комментария


```mutation {
  createChildComment(postID: "5", authorID: "2", content: "Это ответ на комментарий!", parentID: "2") {
    id
    content
    parentId
  }
}

##### Создание пользователя


```mutation {
  createUser(name: "John Doe") {
    id
    name
    createdAt
  }
}

##### Получение поста по ID с комментариями


```query {
  post(id: "5", limit: 10, offset: 0) {
    id
    title
    content
    allowComments
    createdAt
    comments {
      id
      postId
      authorId
      content
      createdAt
      parentId
    }
  }
}

##### Отключение комментария


```mutation {
  toggleAllowComments(postID: "5") {
    id
    title
    allowComments
  }
}

##### Получение всех постов с ограничением на количество


```query {
  posts(limit: 5, offset: 0) {
    id
    title
    content
    createdAt
  }
}