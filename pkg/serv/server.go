package serv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	pr "diplom/api/proto"
	send "diplom/send"

	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

// Server...
type ApiServ struct {
	pr.UnimplementedApiServer
}

// Вытягиваем из базы данные о параметрах оборудования
func (ApiServ) GetHardwareValue(cont context.Context, req *pr.HardwareRequest) (*pr.HardwareResponse, error) {

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Логи 1: %v\n", err)
		return nil, errors.New("Ошибка подключения к БД")
	}

	rows, err := dbPool.Query(context.Background(), "select param_name, current_value from public.unit u join public.params p on p.param_id = u.param_id where u.hardware_id = $1;", req.HarwareId)
	if err != nil {
		log.Fatalf("Логи 2 %v\n", err)
		return nil, errors.New("Ошибка выполнения SQL запроса")
	}

	var ret []*pr.HardwareParams
	for rows.Next() {
		var r pr.HardwareParams
		err := rows.Scan(&r.ParamName, &r.ParamValue)
		if err != nil {
			log.Fatal("Логи 3")
			return nil, errors.New("Ошибка чтения запроса")
		}
		ret = append(ret, &r)

	}

	return &pr.HardwareResponse{MessageId: uuid.NewV4().String(), Params: ret}, nil
}

// Обновление базы и пересылка на сериализатор
func (ApiServ) UpdateParamValue(cont context.Context, req *pr.UpdateRequest) (*pr.UpdateResponse, error) {

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Логи 4 %v\n", err)
		return nil, errors.New("Ошибка подключения к БД")
	}

	for _, val := range req.Params {
		_, err := dbPool.Exec(context.Background(), "UPDATE public.params SET p.current_value=$1 from public.params p join public.unit u on  p.param_id = u.param_id  join public.hardware h on h.hardware_id = u.hardware_id  WHERE h.hardware_id = $2 and p.param_id = $3;", val.ParamValue, req.HardwareId, val.ParamId)
		if err != nil {
			log.Fatal("Логи 5")
			return nil, errors.New("Ошибка выполнения SQL запроса")
		}
	}

	conn, err := grpc.Dial("77.222.42.182:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := send.NewUnaryClient(conn)
	res, err := client.SendToClient(context.Background(), &send.Message{})
	fmt.Println(res)
	return &pr.UpdateResponse{MessageId: uuid.NewV4().String(), ErrorCode: "OK"}, nil
}

// Функция регистрирует пользователя в системе
func (ApiServ) Registration(ctx context.Context, req *pr.RegistrationRequest) (*pr.RegistrationResponse, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Логи 4 %v\n", err)
		return nil, errors.New("Ошибка подключения к БД")
	}
	_, err = dbPool.Exec(context.Background(), "INSERT INTO public.users(login, password, actual) VALUES($1, $2, true);", req.Login, req.Password)
	if err != nil {
		log.Fatal("Логи 5")
		return nil, errors.New("Ошибка выполнения SQL запроса")
	}
	return &pr.RegistrationResponse{MessageId: uuid.NewV4().String(), ErrorCode: "OK"}, nil
}
