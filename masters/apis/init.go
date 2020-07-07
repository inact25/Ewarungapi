package apis

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/masters/apis/controllers"
	"github.com/inact25/E-WarungApi/masters/apis/middlewares"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"github.com/inact25/E-WarungApi/masters/apis/usecases"
)

func Init(r *mux.Router, db *sql.DB) {
	menuRepo := repositories.InitMenuRepoImpl(db)
	menuUseCase := usecases.InitMenuUseCase(menuRepo)
	controllers.MenuControll(r, menuUseCase)

	categoriesRepo := repositories.InitCategoryRepoImpl(db)
	categoriesUseCase := usecases.InitCategoryUseCase(categoriesRepo)
	controllers.CategoriesControll(r, categoriesUseCase)

	transactionRepo := repositories.InitTransactionRepoImpl(db)
	transactionUseCase := usecases.InitTransactionUseCase(transactionRepo)
	controllers.TransactionControll(r, transactionUseCase)

	//mdInit := repositories.InitLogsRepoImpl(db)
	//mdUsc := usecases.InitLogsUseCase(mdInit)

	r.Use(middlewares.ActivityLogMiddleware)
}
