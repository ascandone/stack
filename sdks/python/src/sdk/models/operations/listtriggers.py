"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import error as shared_error
from ..shared import listtriggersresponse as shared_listtriggersresponse
from typing import Optional


@dataclasses.dataclass
class ListTriggersResponse:
    
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    error: Optional[shared_error.Error] = dataclasses.field(default=None)
    r"""General error"""
    list_triggers_response: Optional[shared_listtriggersresponse.ListTriggersResponse] = dataclasses.field(default=None)
    r"""List of triggers"""
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    