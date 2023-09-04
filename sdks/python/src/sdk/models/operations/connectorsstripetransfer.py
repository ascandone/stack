"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import stripetransferresponse as shared_stripetransferresponse
from typing import Optional



@dataclasses.dataclass
class ConnectorsStripeTransferResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    stripe_transfer_response: Optional[shared_stripetransferresponse.StripeTransferResponse] = dataclasses.field(default=None)
    r"""OK"""
    
