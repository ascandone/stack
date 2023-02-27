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


class LogsCursorResponse(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "cursor",
        }
        
        class properties:
            
            
            class cursor(
                schemas.DictSchema
            ):
            
            
                class MetaOapg:
                    required = {
                        "data",
                        "hasMore",
                        "pageSize",
                    }
                    
                    class properties:
                        
                        
                        class pageSize(
                            schemas.Int64Schema
                        ):
                        
                        
                            class MetaOapg:
                                format = 'int64'
                                inclusive_maximum = 1000
                                inclusive_minimum = 1
                        hasMore = schemas.BoolSchema
                        previous = schemas.StrSchema
                        next = schemas.StrSchema
                        
                        
                        class data(
                            schemas.ListSchema
                        ):
                        
                        
                            class MetaOapg:
                                
                                @staticmethod
                                def items() -> typing.Type['Log']:
                                    return Log
                        
                            def __new__(
                                cls,
                                _arg: typing.Union[typing.Tuple['Log'], typing.List['Log']],
                                _configuration: typing.Optional[schemas.Configuration] = None,
                            ) -> 'data':
                                return super().__new__(
                                    cls,
                                    _arg,
                                    _configuration=_configuration,
                                )
                        
                            def __getitem__(self, i: int) -> 'Log':
                                return super().__getitem__(i)
                        __annotations__ = {
                            "pageSize": pageSize,
                            "hasMore": hasMore,
                            "previous": previous,
                            "next": next,
                            "data": data,
                        }
                
                data: MetaOapg.properties.data
                hasMore: MetaOapg.properties.hasMore
                pageSize: MetaOapg.properties.pageSize
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["pageSize"]) -> MetaOapg.properties.pageSize: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["hasMore"]) -> MetaOapg.properties.hasMore: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["previous"]) -> MetaOapg.properties.previous: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["next"]) -> MetaOapg.properties.next: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["data"]) -> MetaOapg.properties.data: ...
                
                @typing.overload
                def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
                
                def __getitem__(self, name: typing.Union[typing_extensions.Literal["pageSize", "hasMore", "previous", "next", "data", ], str]):
                    # dict_instance[name] accessor
                    return super().__getitem__(name)
                
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["pageSize"]) -> MetaOapg.properties.pageSize: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["hasMore"]) -> MetaOapg.properties.hasMore: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["previous"]) -> typing.Union[MetaOapg.properties.previous, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["next"]) -> typing.Union[MetaOapg.properties.next, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["data"]) -> MetaOapg.properties.data: ...
                
                @typing.overload
                def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
                
                def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["pageSize", "hasMore", "previous", "next", "data", ], str]):
                    return super().get_item_oapg(name)
                
            
                def __new__(
                    cls,
                    *_args: typing.Union[dict, frozendict.frozendict, ],
                    data: typing.Union[MetaOapg.properties.data, list, tuple, ],
                    hasMore: typing.Union[MetaOapg.properties.hasMore, bool, ],
                    pageSize: typing.Union[MetaOapg.properties.pageSize, decimal.Decimal, int, ],
                    previous: typing.Union[MetaOapg.properties.previous, str, schemas.Unset] = schemas.unset,
                    next: typing.Union[MetaOapg.properties.next, str, schemas.Unset] = schemas.unset,
                    _configuration: typing.Optional[schemas.Configuration] = None,
                    **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                ) -> 'cursor':
                    return super().__new__(
                        cls,
                        *_args,
                        data=data,
                        hasMore=hasMore,
                        pageSize=pageSize,
                        previous=previous,
                        next=next,
                        _configuration=_configuration,
                        **kwargs,
                    )
            __annotations__ = {
                "cursor": cursor,
            }
    
    cursor: MetaOapg.properties.cursor
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["cursor"]) -> MetaOapg.properties.cursor: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["cursor", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["cursor"]) -> MetaOapg.properties.cursor: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["cursor", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        cursor: typing.Union[MetaOapg.properties.cursor, dict, frozendict.frozendict, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'LogsCursorResponse':
        return super().__new__(
            cls,
            *_args,
            cursor=cursor,
            _configuration=_configuration,
            **kwargs,
        )

from Formance.model.log import Log
