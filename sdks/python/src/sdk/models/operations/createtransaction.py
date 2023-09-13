"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import errorresponse as shared_errorresponse
from ..shared import posttransaction as shared_posttransaction
from ..shared import transactionsresponse as shared_transactionsresponse
from typing import Optional


@dataclasses.dataclass
class CreateTransactionRequest:
    
    ledger: str = dataclasses.field(metadata={'path_param': { 'field_name': 'ledger', 'style': 'simple', 'explode': False }})
    r"""Name of the ledger."""
    post_transaction: shared_posttransaction.PostTransaction = dataclasses.field(metadata={'request': { 'media_type': 'application/json' }})
    r"""The request body must contain at least one of the following objects:
      - `postings`: suitable for simple transactions
      - `script`: enabling more complex transactions with Numscript
    """
    idempotency_key: Optional[str] = dataclasses.field(default=None, metadata={'header': { 'field_name': 'Idempotency-Key', 'style': 'simple', 'explode': False }})
    r"""Use an idempotency key"""
    preview: Optional[bool] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'preview', 'style': 'form', 'explode': True }})
    r"""Set the preview mode. Preview mode doesn't add the logs to the database or publish a message to the message broker."""
    

@dataclasses.dataclass
class CreateTransactionResponse:
    
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    error_response: Optional[shared_errorresponse.ErrorResponse] = dataclasses.field(default=None)
    r"""Error"""
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    transactions_response: Optional[shared_transactionsresponse.TransactionsResponse] = dataclasses.field(default=None)
    r"""OK"""
    