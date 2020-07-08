package repositories

const (
	//-> menu.MenuID, menu.MenuDesc, menu.MenuPrice,  menu.MenuStock
	GetAllMenusQuery = `select p.priceID,m.menuDesc, p.Price, m.menuStock, p.priceDate from price p 
						inner join menu m on m.menuID = p.priceID
						inner join (select priceID, max(priceDate) as maxDate from price 
						group by priceID) pj on p.priceID = pj.priceID and p.priceDate = pj.maxDate;`

	GetAllMenusByStatusQuery = `select p.priceID,m.menuDesc, p.Price, m.menuStock, p.priceDate from price p 
								inner join menu m on m.menuID = p.priceID
								inner join (select priceID, max(priceDate) as maxDate from price 
								group by priceID) pj on p.priceID = pj.priceID and p.priceDate = pj.maxDate where menuStatus = ?;`

	GetAllMenusPrice = `select m.menuID, m.menuDesc, p.price, p.priceDate from price p inner join menu m order by p.priceDate desc;`

	AddNewMenuQuery       = `insert into menu values (?,?,?);`
	AddNewMenuPricesQuery = `insert into price values (?,?,?);`

	UpdateMenuQuery = `update menu set menuDesc = ?, menuStock = ? where menuID = ?;`

	UpdateMenuPriceQuery = `insert into price values (?,?,?)`

	DeleteMenuQuery = `update menu set menuStatus = 'i' where menuID = ?`
)
