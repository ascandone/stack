<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class OrchestrationCreditWalletRequest
{
	#[\JMS\Serializer\Annotation\SerializedName('amount')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\Monetary')]
    public Monetary $amount;
    
    /**
     * The balance to credit
     * 
     * @var ?string $balance
     */
	#[\JMS\Serializer\Annotation\SerializedName('balance')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $balance = null;
    
    /**
     * Metadata associated with the wallet.
     * 
     * @var array<string, mixed> $metadata
     */
	#[\JMS\Serializer\Annotation\SerializedName('metadata')]
    #[\JMS\Serializer\Annotation\Type('array<string, mixed>')]
    public array $metadata;
    
	#[\JMS\Serializer\Annotation\SerializedName('reference')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $reference = null;
    
    /**
     * $sources
     * 
     * @var array<mixed> $sources
     */
	#[\JMS\Serializer\Annotation\SerializedName('sources')]
    #[\JMS\Serializer\Annotation\Type('array<mixed>')]
    public array $sources;
    
	public function __construct()
	{
		$this->amount = new \formance\stack\Models\Shared\Monetary();
		$this->balance = null;
		$this->metadata = [];
		$this->reference = null;
		$this->sources = [];
	}
}
