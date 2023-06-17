# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import externalApi_pb2 as externalApi__pb2


class apiStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Get_hardware_value = channel.unary_unary(
                '/api.api/Get_hardware_value',
                request_serializer=externalApi__pb2.HardwareRequest.SerializeToString,
                response_deserializer=externalApi__pb2.HardwareResponse.FromString,
                )
        self.Update_param_value = channel.unary_unary(
                '/api.api/Update_param_value',
                request_serializer=externalApi__pb2.UpdateRequest.SerializeToString,
                response_deserializer=externalApi__pb2.UpdateResponse.FromString,
                )
        self.Registration = channel.unary_unary(
                '/api.api/Registration',
                request_serializer=externalApi__pb2.RegistrationRequest.SerializeToString,
                response_deserializer=externalApi__pb2.RegistrationResponse.FromString,
                )
        self.Registration_hardware = channel.unary_unary(
                '/api.api/Registration_hardware',
                request_serializer=externalApi__pb2.RegistrationHardwareRequest.SerializeToString,
                response_deserializer=externalApi__pb2.RegistrationResponse.FromString,
                )


class apiServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Get_hardware_value(self, request, context):
        """Метод позволяет получить список с параметрами 
        и их значением для необходимого оборудования.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update_param_value(self, request, context):
        """Метод позволяет менять параметры оборудования. 
        Неуказанные параметры остаются неизменными.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Registration(self, request, context):
        """Метод позволяет зарегистрировать нового пользователя.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Registration_hardware(self, request, context):
        """Метод позволяет зарегистрировать оборудование пользователя
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_apiServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Get_hardware_value': grpc.unary_unary_rpc_method_handler(
                    servicer.Get_hardware_value,
                    request_deserializer=externalApi__pb2.HardwareRequest.FromString,
                    response_serializer=externalApi__pb2.HardwareResponse.SerializeToString,
            ),
            'Update_param_value': grpc.unary_unary_rpc_method_handler(
                    servicer.Update_param_value,
                    request_deserializer=externalApi__pb2.UpdateRequest.FromString,
                    response_serializer=externalApi__pb2.UpdateResponse.SerializeToString,
            ),
            'Registration': grpc.unary_unary_rpc_method_handler(
                    servicer.Registration,
                    request_deserializer=externalApi__pb2.RegistrationRequest.FromString,
                    response_serializer=externalApi__pb2.RegistrationResponse.SerializeToString,
            ),
            'Registration_hardware': grpc.unary_unary_rpc_method_handler(
                    servicer.Registration_hardware,
                    request_deserializer=externalApi__pb2.RegistrationHardwareRequest.FromString,
                    response_serializer=externalApi__pb2.RegistrationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.api', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class api(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Get_hardware_value(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.api/Get_hardware_value',
            externalApi__pb2.HardwareRequest.SerializeToString,
            externalApi__pb2.HardwareResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update_param_value(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.api/Update_param_value',
            externalApi__pb2.UpdateRequest.SerializeToString,
            externalApi__pb2.UpdateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Registration(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.api/Registration',
            externalApi__pb2.RegistrationRequest.SerializeToString,
            externalApi__pb2.RegistrationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Registration_hardware(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.api/Registration_hardware',
            externalApi__pb2.RegistrationHardwareRequest.SerializeToString,
            externalApi__pb2.RegistrationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
