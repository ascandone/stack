# coding: utf-8

"""
    Formance Stack API

    Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />   # noqa: E501

    The version of the OpenAPI document: 1.0.0-20230225
    Contact: support@formance.com
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from Formance import schemas  # noqa: F401


class StageStatus(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "instanceID",
            "stage",
            "startedAt",
        }
        
        class properties:
            stage = schemas.NumberSchema
            instanceID = schemas.StrSchema
            startedAt = schemas.DateTimeSchema
            terminatedAt = schemas.DateTimeSchema
            error = schemas.StrSchema
            __annotations__ = {
                "stage": stage,
                "instanceID": instanceID,
                "startedAt": startedAt,
                "terminatedAt": terminatedAt,
                "error": error,
            }
    
    instanceID: MetaOapg.properties.instanceID
    stage: MetaOapg.properties.stage
    startedAt: MetaOapg.properties.startedAt
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["stage"]) -> MetaOapg.properties.stage: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["instanceID"]) -> MetaOapg.properties.instanceID: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["startedAt"]) -> MetaOapg.properties.startedAt: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["terminatedAt"]) -> MetaOapg.properties.terminatedAt: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["error"]) -> MetaOapg.properties.error: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["stage", "instanceID", "startedAt", "terminatedAt", "error", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["stage"]) -> MetaOapg.properties.stage: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["instanceID"]) -> MetaOapg.properties.instanceID: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["startedAt"]) -> MetaOapg.properties.startedAt: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["terminatedAt"]) -> typing.Union[MetaOapg.properties.terminatedAt, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["error"]) -> typing.Union[MetaOapg.properties.error, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["stage", "instanceID", "startedAt", "terminatedAt", "error", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        instanceID: typing.Union[MetaOapg.properties.instanceID, str, ],
        stage: typing.Union[MetaOapg.properties.stage, decimal.Decimal, int, float, ],
        startedAt: typing.Union[MetaOapg.properties.startedAt, str, datetime, ],
        terminatedAt: typing.Union[MetaOapg.properties.terminatedAt, str, datetime, schemas.Unset] = schemas.unset,
        error: typing.Union[MetaOapg.properties.error, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'StageStatus':
        return super().__new__(
            cls,
            *_args,
            instanceID=instanceID,
            stage=stage,
            startedAt=startedAt,
            terminatedAt=terminatedAt,
            error=error,
            _configuration=_configuration,
            **kwargs,
        )
