package serv

import (
	"context"
	pr "diplom/api/proto"
	send "diplom/send"
	b64 "encoding/base64"
	"errors"
	"log"
	"os"
	"regexp"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// Server...
type ApiServ struct {
	pr.UnimplementedApiServer
}

func GetMessageId(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("could not grab metadata from context")
	}
	MessageId := meta.Get("MessageId")
	return MessageId[0], nil
}

// Вытягиваем из базы данные о параметрах оборудования
func (ApiServ) GetHardwareValue(ctx context.Context, req *pr.HardwareRequest) (*pr.HardwareResponse, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	messageId, err := GetMessageId(ctx)
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("could not grab metadata from context")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Unable to connect to database")
	}

	rows, err := dbPool.Query(context.Background(), "select param_name, current_value from public.unit u join public.params p on p.param_id = u.param_id where u.hardware_id = $1;", req.HarwareId)
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query execution error")
	}

	var ret []*pr.HardwareParams
	for rows.Next() {
		var r pr.HardwareParams
		err := rows.Scan(&r.ParamName, &r.ParamValue)
		if err != nil {
			errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("Error reading result of SQL query")
		}
		ret = append(ret, &r)

	}
	responce := &pr.HardwareResponse{MessageId: messageId, Params: ret}
	infoLog.Printf("GetHardwareValue: request successful. MessageId: %v\n", messageId)
	return responce, nil
}

// Обновление базы и пересылка на сериализатор
func (ApiServ) UpdateParamValue(ctx context.Context, req *pr.UpdateRequest) (*pr.UpdateResponse, error) {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	messageId, err := GetMessageId(ctx)
	if err != nil {
		errorLog.Printf("UpdateParamValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("could not grab metadata from context")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("UpdateParamValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Unable to connect to database")
	}

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorLog.Printf("UpdateParamValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Error reading result of SQL query")
	}

	rows, err := dbPool.Query(context.Background(), "select ip from public.hardware where hardware_id = $1 limit 1;", req.HardwareId)
	if err != nil {
		errorLog.Printf("UpdateParamValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query select execution error")
	}

	var host string
	for rows.Next() {
		err := rows.Scan(&host)
		if err != nil {
			errorLog.Printf("GetParamId: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("Error reading host from SQL query")
		}
	}

	client := send.NewUnaryClient(conn)
	var ret *send.MessageResponse
	for _, val := range req.Params {
		_, err := dbPool.Exec(context.Background(), "UPDATE public.params SET current_value=$1 from public.params p join public.unit u on  p.param_id = u.param_id  join public.hardware h on h.hardware_id = u.hardware_id  WHERE h.hardware_id = $2 and p.param_id = $3;", val.ParamValue, req.HardwareId, val.ParamId)
		if err != nil {
			errorLog.Printf("UpdateParamValue: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("SQL query execution error")
		}

		res, err := client.SendToClient(context.Background(), &send.Message{Host: host, HardId: req.HardwareId, ComandId: val.ParamId, Value: val.ParamValue, MessageId: messageId})
		if err != nil {
			errorLog.Printf("UpdateParamValue: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("Function SendToClient error")
		}
		ret = res
	}

	responce := &pr.UpdateResponse{MessageId: messageId, ErrorCode: ret.ErrorCode}
	infoLog.Printf("UpdateParamValue: request successful. MessageId: %v\n", messageId)
	return responce, nil
}

// Функция регистрирует пользователя в системе
func (ApiServ) Registration(ctx context.Context, req *pr.RegistrationRequest) (*pr.RegistrationResponse, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	messageId, err := GetMessageId(ctx)
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("could not grab metadata from context")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("Registration: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Unable to connect to database")
	}
	if req.Password == "" || len(req.Password) <= 6 {
		errorLog.Printf("Registration: %v MessageId : %v\n", errors.New("Password must be more than 6 characters"), messageId)
		return nil, errors.New("Password must be more than 6 characters")
	}

	_, err = dbPool.Exec(context.Background(), "INSERT INTO public.users(login, password, token) VALUES($1, $2, $3);", req.Login, req.Password, b64.StdEncoding.EncodeToString([]byte(req.Login+":"+req.Password)))
	if err != nil {
		errorLog.Printf("Registration: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query execution error")
	}
	responce := &pr.RegistrationResponse{MessageId: messageId, ErrorCode: "OK"}
	infoLog.Printf("Registration: request successful. MessageId: %v\n", messageId)
	return responce, nil
}

// Функция для регистрации оборудования
func (ApiServ) RegistrationHardware(ctx context.Context, req *pr.RegistrationHardwareRequest) (*pr.RegistrationResponse, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	messageId, err := GetMessageId(ctx)
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("could not grab metadata from context")
	}
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Unable to connect to database")
	}

	match, _ := regexp.MatchString(`([0-9]{1,3}[\.]){3}[0-9]{1,3}`, req.Ip)
	if !match {
		errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", errors.New("No valid Host"), messageId)
		return nil, errors.New("No valid Host")
	}

	rows, err := dbPool.Query(context.Background(), "select user_id from public.users where token = $1 limit 1;", req.Token)
	if err != nil {
		errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query select execution error")
	}

	var userId int

	rows.Next()
	err = rows.Scan(&userId)
	if err != nil {
		errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Error reading 2 result of SQL query")
	}

	var hardId int
	hardRows, err := dbPool.Query(context.Background(), "INSERT INTO public.hardware(hard_name, ip) VALUES($1, $2) returning hardware_id", req.HardName, req.Ip)
	if err != nil {
		errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query insert 1 execution error")
	}

	hardRows.Next()
	err = hardRows.Scan(&hardId)
	if err != nil {
		errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Error reading 1 result of SQL query")
	}

	for _, val := range req.Params {
		var paramId int
		ParamRows, err := dbPool.Query(context.Background(), "INSERT INTO public.params(param_name, current_value) VALUES($1, $2) returning param_id", val.ParamName, val.ParamValue)
		if err != nil {
			errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("SQL query insert 2 execution error")
		}
		ParamRows.Next()
		err = ParamRows.Scan(&paramId)
		if err != nil {
			errorLog.Printf("RegistrationHardware: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("Error reading result of SQL query")
		}
		infoLog.Println(val, paramId)
		_, err = dbPool.Exec(context.Background(), "INSERT INTO public.unit (hardware_id, user_id, param_id) VALUES($1, $2, $3);", hardId, userId, paramId)
		if err != nil {
			errorLog.Printf("RegistrationHardware: %vMessageId : %v\n", err, messageId)
			return nil, errors.New("SQL query insert 3 execution error")
		}
	}

	responce := &pr.RegistrationResponse{MessageId: messageId, ErrorCode: "OK"}
	infoLog.Printf("RegistrationHardware: request successful. MessageId: %v\n", messageId)
	return responce, nil
}

func (ApiServ) GetHardwareId(ctx context.Context, req *pr.HardwareIdRequest) (*pr.HardwereIdResponce, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)

	messageId, err := GetMessageId(ctx)
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("could not grab metadata from context")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("GetHardwareId: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Unable to connect to database")
	}

	rows, err := dbPool.Query(context.Background(), "select h.hard_name, h.hardware_id from public.hardware h join unit u on u.hardware_id = h.hardware_id join public.users us on us.user_id = u.user_id where us.token = $1;", req.Token)
	if err != nil {
		errorLog.Printf("GetHardwareId: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query select execution error")
	}

	var ret []*pr.HardwareIdAll
	for rows.Next() {
		var r pr.HardwareIdAll
		err := rows.Scan(&r.HardwareName, &r.HardwareId)
		if err != nil {
			errorLog.Printf("GetHardwareId: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("Error reading result of SQL query")
		}
		ret = append(ret, &r)

	}
	responce := &pr.HardwereIdResponce{MessageId: messageId, Rows: ret}
	infoLog.Printf("GetHardwareId: request successful. MessageId: %v\n", messageId)
	return responce, nil

}

func (ApiServ) GetParamId(ctx context.Context, req *pr.ParamIdRequest) (*pr.ParamIdResponce, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)

	messageId, err := GetMessageId(ctx)
	if err != nil {
		errorLog.Printf("GetHardwareValue: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("could not grab metadata from context")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("GetParamId: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("Unable to connect to database")
	}

	rows, err := dbPool.Query(context.Background(), "select p.param_name, p.param_id from public.params p join unit u on u.p.param_id = p.param_id join public.users us on us.user_id = u.user_id join public.hardware u on u.hardware_id = h.hardware_id where us.token = $1 and h.hardware_id = $2 limit 1;", req.Token, req.HardwareId)
	if err != nil {
		errorLog.Printf("GetParamId: %v MessageId : %v\n", err, messageId)
		return nil, errors.New("SQL query select execution error")
	}

	var ret []*pr.ParamIdAll
	for rows.Next() {
		var r pr.ParamIdAll
		err := rows.Scan(&r.ParamName, &r.ParamId)
		if err != nil {
			errorLog.Printf("GetParamId: %v MessageId : %v\n", err, messageId)
			return nil, errors.New("Error reading result of SQL query")
		}
		ret = append(ret, &r)

	}
	responce := &pr.ParamIdResponce{MessageId: messageId, Rows: ret}
	infoLog.Printf("GetParamId: request successful. MessageId: %v\n", messageId)
	return responce, nil
}
