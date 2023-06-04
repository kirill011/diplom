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
		return nil, err
	}

	rows, err := dbPool.Query(context.Background(), "select param_name, current_value from public.unit u join public.params p on p.param_id = u.param_id where u.hardware_id = $1;", req.HarwareId)
	if err != nil {
		log.Fatalf("Я снова еблан забыл переделать логи %v\n", err)
		return nil, err
	}

	var ret []*pr.HardwareParams
	for rows.Next() {
		var r pr.HardwareParams
		err := rows.Scan(&r.ParamName, &r.ParamValue)
		if err != nil {
			log.Fatal("Кто еблан забывший переделать логи?")
			return nil, err
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
		return nil, err
	}
	//Скорее всего можно сделать лучше, но пока и так сойдёт
	for _, val := range req.Params {
		_, err := dbPool.Exec(context.Background(), "UPDATE public.params SET p.current_value=$1 from public.params p join public.unit u on  p.param_id = u.param_id  join public.hardware h on h.hardware_id = u.hardware_id  WHERE h.hardware_id = $2 and p.param_id = $3;", val.ParamValue, req.HardwareId, val.ParamId)
		if err != nil {
			log.Fatal("Я снова еблан забыл переделать логи")
			return nil, err
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
