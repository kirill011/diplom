from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class HardwareIdAll(_message.Message):
    __slots__ = ["HardwareId", "HardwareName"]
    HARDWAREID_FIELD_NUMBER: _ClassVar[int]
    HARDWARENAME_FIELD_NUMBER: _ClassVar[int]
    HardwareId: int
    HardwareName: str
    def __init__(self, HardwareName: _Optional[str] = ..., HardwareId: _Optional[int] = ...) -> None: ...

class HardwareIdRequest(_message.Message):
    __slots__ = ["Token"]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    Token: str
    def __init__(self, Token: _Optional[str] = ...) -> None: ...

class HardwareParams(_message.Message):
    __slots__ = ["ParamName", "ParamValue"]
    PARAMNAME_FIELD_NUMBER: _ClassVar[int]
    PARAMVALUE_FIELD_NUMBER: _ClassVar[int]
    ParamName: str
    ParamValue: float
    def __init__(self, ParamName: _Optional[str] = ..., ParamValue: _Optional[float] = ...) -> None: ...

class HardwareRequest(_message.Message):
    __slots__ = ["HarwareId", "Token"]
    HARWAREID_FIELD_NUMBER: _ClassVar[int]
    HarwareId: int
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    Token: str
    def __init__(self, HarwareId: _Optional[int] = ..., Token: _Optional[str] = ...) -> None: ...

class HardwareResponse(_message.Message):
    __slots__ = ["MessageId", "Params"]
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    MessageId: str
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    Params: _containers.RepeatedCompositeFieldContainer[HardwareParams]
    def __init__(self, MessageId: _Optional[str] = ..., Params: _Optional[_Iterable[_Union[HardwareParams, _Mapping]]] = ...) -> None: ...

class HardwereIdResponce(_message.Message):
    __slots__ = ["MessageId", "Rows"]
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    MessageId: str
    ROWS_FIELD_NUMBER: _ClassVar[int]
    Rows: _containers.RepeatedCompositeFieldContainer[HardwareIdAll]
    def __init__(self, MessageId: _Optional[str] = ..., Rows: _Optional[_Iterable[_Union[HardwareIdAll, _Mapping]]] = ...) -> None: ...

class ParamIdAll(_message.Message):
    __slots__ = ["ParamId", "ParamName"]
    PARAMID_FIELD_NUMBER: _ClassVar[int]
    PARAMNAME_FIELD_NUMBER: _ClassVar[int]
    ParamId: int
    ParamName: str
    def __init__(self, ParamName: _Optional[str] = ..., ParamId: _Optional[int] = ...) -> None: ...

class ParamIdRequest(_message.Message):
    __slots__ = ["HardwareId", "Token"]
    HARDWAREID_FIELD_NUMBER: _ClassVar[int]
    HardwareId: int
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    Token: str
    def __init__(self, Token: _Optional[str] = ..., HardwareId: _Optional[int] = ...) -> None: ...

class ParamIdResponce(_message.Message):
    __slots__ = ["MessageId", "Rows"]
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    MessageId: str
    ROWS_FIELD_NUMBER: _ClassVar[int]
    Rows: _containers.RepeatedCompositeFieldContainer[ParamIdAll]
    def __init__(self, MessageId: _Optional[str] = ..., Rows: _Optional[_Iterable[_Union[ParamIdAll, _Mapping]]] = ...) -> None: ...

class RegistrationHardwareRequest(_message.Message):
    __slots__ = ["HardName", "Ip", "Params", "Token"]
    HARDNAME_FIELD_NUMBER: _ClassVar[int]
    HardName: str
    IP_FIELD_NUMBER: _ClassVar[int]
    Ip: str
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    Params: _containers.RepeatedCompositeFieldContainer[HardwareParams]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    Token: str
    def __init__(self, HardName: _Optional[str] = ..., Ip: _Optional[str] = ..., Token: _Optional[str] = ..., Params: _Optional[_Iterable[_Union[HardwareParams, _Mapping]]] = ...) -> None: ...

class RegistrationRequest(_message.Message):
    __slots__ = ["Login", "Password"]
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    Login: str
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    Password: str
    def __init__(self, Login: _Optional[str] = ..., Password: _Optional[str] = ...) -> None: ...

class RegistrationResponse(_message.Message):
    __slots__ = ["ErrorCode", "MessageId"]
    ERRORCODE_FIELD_NUMBER: _ClassVar[int]
    ErrorCode: str
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    MessageId: str
    def __init__(self, MessageId: _Optional[str] = ..., ErrorCode: _Optional[str] = ...) -> None: ...

class UpdateParams(_message.Message):
    __slots__ = ["ParamId", "ParamValue"]
    PARAMID_FIELD_NUMBER: _ClassVar[int]
    PARAMVALUE_FIELD_NUMBER: _ClassVar[int]
    ParamId: int
    ParamValue: float
    def __init__(self, ParamId: _Optional[int] = ..., ParamValue: _Optional[float] = ...) -> None: ...

class UpdateRequest(_message.Message):
    __slots__ = ["HardwareId", "Params", "Token"]
    HARDWAREID_FIELD_NUMBER: _ClassVar[int]
    HardwareId: int
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    Params: _containers.RepeatedCompositeFieldContainer[UpdateParams]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    Token: str
    def __init__(self, HardwareId: _Optional[int] = ..., Token: _Optional[str] = ..., Params: _Optional[_Iterable[_Union[UpdateParams, _Mapping]]] = ...) -> None: ...

class UpdateResponse(_message.Message):
    __slots__ = ["ErrorCode", "MessageId"]
    ERRORCODE_FIELD_NUMBER: _ClassVar[int]
    ErrorCode: str
    MESSAGEID_FIELD_NUMBER: _ClassVar[int]
    MessageId: str
    def __init__(self, MessageId: _Optional[str] = ..., ErrorCode: _Optional[str] = ...) -> None: ...
