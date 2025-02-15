package repository

const (
	UserTableName = "user_"

	UserIdColumn = "id"
	UserUsernameColumn = "username"
	UserPasswordColumn = "password_hash"
	UserCoinsColumn = "coins"

	TransactionTableName = "user_transaction_"

	TransactionIdColumn = "id"
	TransactionIdSenderColumn = "id_sender"
	TransactionIdReceiverColumn = "id_receiver"
	TransactionAmountColumn = "amount"

	InventoryTableName = "inventory_"

	InventoryIdColumn = "id"
	InventoryIdUserColumn = "id_user"
	InventoryIdMerchColumn = "id_merch"
	InventoryQuantityColumn = "quantity"

	MerchTableName = "merch_"

	MerchIdColumn = "id"
	MerchTitleColumn = "title"
	MerchPriceColumn = "price"
)