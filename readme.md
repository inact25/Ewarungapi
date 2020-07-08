-- Menu
#GET EXAMPLE
menus/	=> ALL
menus/a => Active
menus/i => Inactive 

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
    "menu_stock": "150"
}

==> Price
{
     "menu_id": "M001",
     "menu_price": "17500"
}

#DELETE :
menus/menu_id => SOFT DELETE


