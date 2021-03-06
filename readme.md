# EWarungApi

## Feature
- view all menu
- view active or inactive menu
- addmenu (product ex: food, fashion and etc)
- menu price update without affecting an old data (menu price and transactions)
- update menu
- delete menu (safe delete)

- view all services
- view active or inactive services
- add service (services ex: live music, romance dinner and etc )
- service price update without affecting an old data (service price and transactions)
- update service
- delete service (safe delete)

- view all categories
- view active or inactive categories
- add category (category ex : food category = 'pedas','normal' and etc)
- category price update without affecting an old data (category price and transactions)
- update service
- delete service (safe delete)

- view all transactions
- view daily transactions
- add transaction (with latest price data)
- update transaction (with latest price data)

## database design
![database design](https://1.bp.blogspot.com/-yOYChEY4Irk/XwYF-JEC0_I/AAAAAAAAIr0/1jhx8lNRKQ4o1tj5UFbmA3KRRTMy4oN6ACK4BGAsYHg/s867/Annotation%2B2020-07-09%2B004401.png)
 
# feature on upcoming state
- api validation - Checking On Progress
- jwt authentication
- multi add transactionS

# Menu
#### GET EXAMPLE
````
menus/		=> ALL
menus/prices	=> Prices Data
menus/a 	=> Active
menus/i 	=> Inactive 
````

#### POST EXAMPLE :
````
{
    "menu_id": "M003",
    "menu_desc": "Ayam Rendang",
    "menu_stock": "10",
    "menu_price": "17000"
}
````

### PUT EXAMPLE :

##### MenuPath : /menu
````
{
    "menu_id": "M001",
    "menu_desc": "Ayam Crispy",
    "menu_stock": "150",
    "menu_status": "A"
}
````
#### PricePath : /menu/prices
````
{
     "menu_id": "M001",
     "menu_price": "17500"
}
````
#### DELETE :
````
menus/{menu_id} => SOFT DELETE
````

# Services

#### GET EXAMPLE
````
services/	=> ALL
services/a	=> Active
services/i	=> Inactive 
````

#### POST EXAMPLE:
````
{
     "services_id": "S002",
     "services_desc": "Wifi Premium",
     "price_date": "5000",
     "services_status": "A"        
}
````

#### PUT EXAMPLE :

##### ServicesPath : /services
````
{
    "services_id": "S001",
    "services_desc": "Wifi Pertalite",
    "services_status": "A"
}
````

##### ServicesPricePath : /services/prices
````
{
     "services_id": "S002",
     "services_price": "1500"
}
````

#### DELETE :
````
services/{services_id} => SOFT DELETE
````

# Categories

#### GET EXAMPLE
````
categories/	=> ALL
categories/a	=> Active
categories/i	=> Inactive 
````

#### POST EXAMPLE:
````
{
     "categories_id": "S002",
     "categories_desc": "Wifi Premium",
     "price_date": "5000",
     "categories_status": "A"        
}
````

#### PUT EXAMPLE :

##### Categories : /categories
````
{
    "categories_id": "C001",
    "categories_desc": "Pedes Betul",
    "categories_status": "A"
}
````

##### Categories Price : /categories/prices
````
{
     "categories_id": "C001",
     "categories_price": "2500"
}
````

#### DELETE :
````
categories/{categories_id} => SOFT DELETE
````

# Transactions

#### GET EXAMPLE
````
/transactions
````

#### POST EXAMPLE
````
{
     "services_desc": "S002",
     "menu_desc": "M001",
     "category_desc": "C003",
     "qty":"4"
}
````

#### PUT EXAMPLE
````
{ 
     "transaction_id": "84d60a11-bffa-11ea-b9f1-705a0f09d721",
     "services_desc": "S002",
     "menu_desc": "M001",
     "category_desc": "C003",
     "qty":"7"
}
````
