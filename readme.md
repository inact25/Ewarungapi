-- Menu
#GET EXAMPLE
menus/		=> ALL
menus/prices	=> Prices Data
menus/a 	=> Active
menus/i 	=> Inactive 

#POST EXAMPLE :
{
    "menu_id": "M003",
    "menu_desc": "Ayam Rendang",
    "menu_stock": "10",
    "menu_price": "17000"
}

#PUT EXAMPLE :

==> Menu
{
    "menu_id": "M001",
    "menu_desc": "Ayam Crispy",
    "menu_stock": "150",
    "menu_status": "A"
}

==> Price
{
     "menu_id": "M001",
     "menu_price": "17500"
}

#DELETE :
menus/menu_id => SOFT DELETE

-- Services
#GET EXAMPLE
services/	=> ALL
services/a	=> Active
services/i	=> Inactive 

#POST EXAMPLE:
{
     "services_id": "S002",
     "services_desc": "Wifi Premium",
     "price_date": "5000",
     "services_status": "A"        
}


