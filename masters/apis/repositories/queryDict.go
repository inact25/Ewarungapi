package repositories

//Menu
const (
	GetAllMenusQuery = `select p.priceID,m.menuDesc, p.Price, m.menuStock,m.menuStatus, p.priceDate from price p 
						inner join menu m on m.menuID = p.priceID
						inner join (select priceID, max(priceDate) as maxDate from price 
						group by priceID) pj on p.priceID = pj.priceID and p.priceDate = pj.maxDate;`

	GetAllMenusByStatusQuery = `select p.priceID,m.menuDesc, p.Price, m.menuStock, m.menuStatus,p.priceDate from price p 
								inner join menu m on m.menuID = p.priceID
								inner join (select priceID, max(priceDate) as maxDate from price 
								group by priceID) pj on p.priceID = pj.priceID and p.priceDate = pj.maxDate
							 	where menuStatus = ?;`

	GetAllMenusPrice = `select m.menuID, m.menuDesc, p.price, p.priceDate from price p inner join menu m 
						order by p.priceDate desc;`

	AddNewMenuQuery       = `insert into menu values (?,?,?,?);`
	AddNewMenuPricesQuery = `insert into price values (?,?,?);`

	UpdateMenuQuery = `update menu set menuDesc = ?, menuStock = ?, menuStatus = ? where menuID = ?;`

	UpdateMenuPriceQuery = `insert into price values (?,?,?)`

	DeleteMenuQuery = `update menu set menuStatus = 'i' where menuID = ?`
)

//Services
const (
	GetAllServices = `select sp.priceID,s.servicesDesc, sp.Price,s.servicesStatus, sp.priceDate from servicesprice sp 
					  inner join services s on s.servicesID = sp.priceID
					  inner join (select priceID, max(priceDate) as maxDate from servicesprice 
					  group by priceID) pj on sp.priceID = pj.priceID and sp.priceDate = pj.maxDate;`

	GetAllServicesByStatus = `select sp.priceID,s.servicesDesc, sp.Price,s.servicesStatus, sp.priceDate from servicesprice sp 
							  inner join services s on s.servicesID = sp.priceID
							  inner join (select priceID, max(priceDate) as maxDate from servicesprice 
							  group by priceID) pj on sp.priceID = pj.priceID and sp.priceDate = pj.maxDate 
							  where servicesStatus = ?;`

	AddNewServicesQuery       = `insert into services values (?,?,?);`
	AddNewServicesPricesQuery = `insert into servicesprice values (?,?,?);`

	UpdateServicesQuery = `update services set servicesDesc = ?, servicesStatus = ? where servicesID = ?;`

	UpdateServicesPriceQuery = `insert into servicesprice values (?,?,?)`

	DeleteServicesQuery = `update services set servicesStatus = 'i' where servicesID = ?`
)

//Categories
const (
	GetAllCategories = `select cp.priceID,c.categoryDesc, cp.Price,c.categoryStatus, cp.priceDate from categoriesprice cp 
						inner join category c on c.categoryID = cp.priceID
						inner join (select priceID, max(priceDate) as maxDate from categoriesprice 
						group by priceID) pj on cp.priceID = pj.priceID and cp.priceDate = pj.maxDate;`

	GetAllCategoriesByStatus = `select cp.priceID,c.categoryDesc, cp.Price,c.categoryStatus, cp.priceDate from categoriesprice cp 
						inner join category c on c.categoryID = cp.priceID
						inner join (select priceID, max(priceDate) as maxDate from categoriesprice 
						group by priceID) pj on cp.priceID = pj.priceID and cp.priceDate = pj.maxDate 
						where categoryStatus = ?;`

	AddNewCategoriesQuery       = `insert into category values (?,?,?);`
	AddNewCategoriesPricesQuery = `insert into categoriesprice values (?,?,?);`

	UpdateCategoriesQuery = `update category set categoryDesc = ?, categoryStatus = ? where categoryID = ?;`

	UpdateCategoriesPriceQuery = `insert into categoriesprice values (?,?,?)`

	DeleteCategoriesQuery = `update category set categoryStatus = 'i' where categoryID = ?`
)

//transaction
const (
	GetLatestMenuPriceByIDQuery = `select p.Price from price p 
								   inner join menu m on m.menuID = p.priceID
								   inner join (select priceID, max(priceDate) as maxDate from price 
								   group by priceID) pj on p.priceID = pj.priceID and p.priceDate = pj.maxDate
								   where p.priceID = ?;`

	GetLatestServicePriceByIDQuery = `select sp.Price from servicesprice sp 
									  inner join services s on s.servicesID = sp.priceID
								   	  inner join (select priceID, max(priceDate) as maxDate from servicesprice 
								   	  group by priceID) pj on sp.priceID = pj.priceID and sp.priceDate = pj.maxDate
								   	  where sp.priceID = ?;`

	GetLatestCategoryPriceByIDQuery = `select cp.Price from categoriesprice cp 
									   inner join category c on c.categoryID = cp.priceID
									   inner join (select priceID, max(priceDate) as maxDate from categoriesprice 
									   group by priceID) pj on cp.priceID = pj.priceID and cp.priceDate = pj.maxDate
									   where cp.priceID = ?;`

	GetAllTransactionsQuery = `select t.transactionID, t.transactionDate,m.menuDesc,t.menuPrice,c.categoryDesc,t.categoryPrice,t.qty,s.servicesDesc, t.servicePrice, 
							   ((t.menuPrice+t.categoryPrice)*t.qty)+servicePrice as SubTotal
							   from transaction t 
							   inner join services s on t.servicesID = s.servicesID
							   inner join menu m on t.menuID = m.menuID
							   inner join category c on t.categoryID = c.categoryID;`

	AddNewTransactions = `insert into transaction values (uuid(), ?, ?, ?, ?, ?, ?, ?, ?);`

	UpdateTransactions = `update transaction set servicesID = ?, servicePrice = ?, menuID = ?,menuPrice = ?, 
						  categoryID = ?, categoryPrice = ?, Qty = ? where transactionID = ?;`
)
