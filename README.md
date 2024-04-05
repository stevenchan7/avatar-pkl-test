# avatar-product

## Prerequisite to run this program locally

on /main subdirectory, create an .env file containing:
```bash
TOKEN_LIFESPAN=1
SECRET_JWT=avatarsolution

DB_USER=root
DB_PASS="YOUR_DB_PASS"
DB_NAME=avatar_pkl_test
APP_KEY="YOUR_GOOGLE_ACC_APP_KEY"
```

## Public routes
- /todo-item
  
  Endpoint for getting all todo items. Return a JSON containing:
  ```
  {
    "data": [
        {
            "id": 1,
            "title": "tugas satu",
            "desc": "deskripsi tugas 1",
            "status": "undone",
            "due_date": "2024-04-05T18:00:00Z",
            "CreatedAt": "2024-04-05T20:30:12.84+08:00",
            "UpdatedAt": "2024-04-05T20:30:12.84+08:00",
            "DeletedAt": null
        },
        {
            "id": 2,
            "title": "tugas dua",
            "desc": "deskripsi tugas 2",
            "status": "undone",
            "due_date": "2024-04-05T18:00:00Z",
            "CreatedAt": "2024-04-05T20:30:12.84+08:00",
            "UpdatedAt": "2024-04-05T20:30:12.84+08:00",
            "DeletedAt": null
        }
    ],
    "success": true
  }
  ```

- /todo-item/:id

  Endpoint for fetching product by ID. Return product.
  

## Public auth routes
- /auth/register
  
  Endpoint for register new admin account. Required request body
  
  > username, password, email

- /auth/login
  
  Endpoint for log into existing admin account. Required request body:
  
  > username, password

- /auth/logout

  Endpoint for log out, removing existing cookie.

- /pass-recovery
  
  Endpoint to sent password recovery link to user's email

- /change-pass
  
  Endpoint to reset password. Required request body:
  
  > newPassword1, newPassword2, username, verHash(get from /pass-recovery)

## Admin routes (must login)
- /admin/add-todo-iteim
  
  Endpoint for adding new todo item into database. Required request body:
  
  > title, desc, status, due_date

- /admin/update-todo-item

  Endpoint for update existing todo item in database. Required request body:

  > title, desc, status, due_date, id

- /admin/delete-todo-item

  Endpoint for delete existing product from database. Required request body:

  > id
 
- /admin/update-todo-item-status

  Endpoint for update existing todo item status in database. Required request body:

  > status(undone, on progress, done), id
