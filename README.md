# avatar-product

## Prerequisite to run this program locally

on /main subdirectory, create an .env file containing:
```bash
DB_USER = your mysql database username
DB_PASS = your mysql database username
DB_NAME = your mysql database name
SECRET_JWT = some random string
TOKEN_LIFESPAN = 1
```

## Public routes
- /products
  
  Endpoint for getting all products. Return a JSON containing:
  
  > {"success": true, "data": []Product}

- /products/:productID

  Endpoint for fetching product by ID. Return product.
  

## Public auth routes
- /auth/register
  
  Endpoint for register new admin account. Required request body
  
  > username, password

- /auth/login
  
  Endpoint for log into existing admin account. Required request body:
  > username, password

- /auth/logout
  Endpoint for log out, removing existing cookie.

## Admin routes (must login)
- /admin/add-product
  
  Endpoint for adding new product into database. Required request body:
  
  > title, desc, image (link to image drive), playstore (play store link), appstore (app store link)

- /admin/edit-product

  Endpoint for update existing product in database. Required request body:

  > title, desc, image (link to image drive), playstore (play store link), appstore (app store link), prodID

- /admin/delete-product

  Endpoint for delete existing product from database. Required request body:

  > prodID
 
