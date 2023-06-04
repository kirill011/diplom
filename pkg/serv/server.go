package serv

import (
	"context"
	"log"
	"os"

	pr "diplom/api/proto"

	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/satori/go.uuid"
)

// Server...
type ApiServ struct {
	pr.UnimplementedApiServer
}

// Вытягиваем из базы данные о параметрах оборудования
func (ApiServ) GetHardwareValue(cont context.Context, req *pr.HardwareRequest) (*pr.HardwareResponse, error) {

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Я еблан забыл переделать логи: %v\n", err)
	}

	rows, err := dbPool.Query(context.Background(), "select param_name, current_value from public.hardware h join public.params p on p.param_id = h.param_id where h.hardware_id = $1;", req.HarwareId)
	if err != nil {
		log.Fatalf("Я снова еблан забыл переделать логи %v\n", err)
	}

	var ret []*pr.HardwareParams
	for rows.Next() {
		var r pr.HardwareParams
		err := rows.Scan(&r.ParamName, &r.ParamValue)
		if err != nil {
			log.Fatal("Кто еблан забывший переделать логи?")
		}
		ret = append(ret, &r)

	}
	//Добавить генерацию Message id и узнать действительно ли она нужна
	return &pr.HardwareResponse{MessageId: uuid.NewV4().String(), Params: ret}, nil
}

// Обновление базы и пересылка на сериализатор
func (ApiServ) UpdateParamValue(cont context.Context, req *pr.UpdateRequest) (*pr.UpdateResponse, error) {

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Я еблан забыл переделать логи: %v\n", err)
	}
	//Скорее всего можно сделать лучше, но пока и так сойдёт
	for _, val := range req.Params {
		_, err := dbPool.Exec(context.Background(), "UPDATE public.hardware SET current_value=$1 WHERE hardware_id = $2 and param_id = $3;", val.ParamValue, req.HardwareId, val.ParamId)
		if err != nil {
			log.Fatal("Я снова еблан забыл переделать логи")
		}
	}

	//Далее идёт передача на сериализатор, которого пока нет

	//Добавить генерацию Message id и узнать действительно ли она нужна
	return &pr.UpdateResponse{MessageId: uuid.NewV4().String(), ErrorCode: "OK"}, nil
}

// Нихуя не работает
func (ApiServ) Registration(ctx context.Context, req *pr.RegistrationRequest) (*pr.RegistrationResponse, error) {
	return &pr.RegistrationResponse{}, nil
}
