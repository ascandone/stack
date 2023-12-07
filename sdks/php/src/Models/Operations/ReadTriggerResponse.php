<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;


class ReadTriggerResponse
{
	
    public string $contentType;
    
    /**
     * General error
     * 
     * @var ?\formance\stack\Models\Shared\Error $error
     */
	
    public ?\formance\stack\Models\Shared\Error $error = null;
    
    /**
     * A specific trigger
     * 
     * @var ?\formance\stack\Models\Shared\ReadTriggerResponse $readTriggerResponse
     */
	
    public ?\formance\stack\Models\Shared\ReadTriggerResponse $readTriggerResponse = null;
    
	
    public int $statusCode;
    
	
    public ?\Psr\Http\Message\ResponseInterface $rawResponse = null;
    
	public function __construct()
	{
		$this->contentType = "";
		$this->error = null;
		$this->readTriggerResponse = null;
		$this->statusCode = 0;
		$this->rawResponse = null;
	}
}
