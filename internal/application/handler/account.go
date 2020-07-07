package handler

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/shiningrush/demo-service/internal/application/service"
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/domain/repo/account"
	"github.com/shiningrush/demo-service/internal/infrastructure/repo"
	"github.com/shiningrush/demo-service/internal/utils/dao"
	"net/http"
	"strconv"
)

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		ar: repo.NewAccountRepo(dao.DefaultDB),
	}
}

type AccountHandler struct {
	ar account.Repo
}

func (h *AccountHandler) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/demo-service/accounts")

	ws.POST("").To(h.Create)
	ws.GET("{id}").To(h.Get)
	ws.DELETE("").To(h.Delete)
	ws.POST("{id}:pay").To(h.Pay)
	ws.POST("{id}:transfer").To(h.Transfer)
	ws.POST("{id}:close").To(h.Close)

	return ws
}

func (h *AccountHandler) Create(req *restful.Request, resp *restful.Response) {
	phone, err := req.BodyParameter("phone")
	if err != nil {
		handleErr(err, resp)
		return
	}

	newAcc := entity.NewAccount(phone)
	savedAcc, err := h.ar.Create(newAcc)
	if err != nil {
		handleErr(err, resp)
		return
	}

	err = resp.WriteHeaderAndEntity(http.StatusOK, savedAcc)
	if err != nil {
		handleErr(err, resp)
	}
}

func (h *AccountHandler) Get(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handleErr(err, resp)
	}

	acc, err := h.ar.Get(intId)
	if err != nil {
		handleErr(err, resp)
	}

	err = resp.WriteHeaderAndEntity(http.StatusOK, acc)
	if err != nil {
		handleErr(err, resp)
	}
}

func (h *AccountHandler) ChangePhone(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		handleErr(err, resp)
	}
	newPhone, err := req.BodyParameter("newPhone")
	if err != nil {
		handleErr(err, resp)
	}

	svc := service.NewAccountAppSvc(repo.NewAccountRepo(dao.DefaultDB), repo.NewBillRepo(dao.DefaultDB))
	err = svc.ChangePhone(intId, newPhone)
	if err != nil {
		handleErr(err, resp)
	}
	resp.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) Pay(req *restful.Request, resp *restful.Response) {

}

func (h *AccountHandler) Delete(req *restful.Request, resp *restful.Response) {

}

type TransferInput struct {
	FromID int
	ToID   int
	Amount float64
	Desc   string
}

func (h *AccountHandler) Transfer(req *restful.Request, resp *restful.Response) {
	input := TransferInput{}
	if err := req.ReadEntity(&input); err != nil {
		handleErr(err, resp)
	}

	trans := dao.DefaultDB.Begin()
	ar, cr := repo.NewAccountRepo(trans), repo.NewBillRepo(trans)
	svc := service.NewAccountAppSvc(ar, cr)
	if err := svc.Transfer(input.FromID, input.ToID, input.Amount, input.Desc); err != nil {
		if rErr := trans.Rollback().Error; rErr != nil {
			handleErr(rErr, resp)
			return
		}
		handleErr(err, resp)
	}

	if rErr := trans.Commit().Error; rErr != nil {
		handleErr(rErr, resp)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) Close(req *restful.Request, resp *restful.Response) {

}

func handleErr(err error, resp *restful.Response) {
	wErr := resp.WriteError(http.StatusInternalServerError, err)
	if wErr != nil {
		panic(wErr)
	}
}
