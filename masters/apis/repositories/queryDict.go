package repositories

const (
	//-> menu.MenuID, menu.MenuDesc, menu.MenuPrice,  menu.MenuStock
	GetAllMenusQuery = `select p.priceID,m.menuDesc, p.Price, m.menuStock from price p 
inner join menu m on m.menuID = p.priceID
inner join (select priceID, max(priceDate) as maxDate from price 
group by priceID) pj on p.priceID = pj.priceID and p.priceDate = pj.maxDate;`

	// ->
	GetAllMenusPrice = `select priceID,priceDate,price from price order by priceID;`

	//-> menus.MenuID, menus.MenuDesc, menus.MenuStock
	AddNewMenuQuery       = `insert into menu values (?,?,?)`
	AddNewMenuPricesQuery = `insert into price values (?,?,?)`
)
