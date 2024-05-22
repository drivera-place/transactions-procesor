# Requirements

```
Terraform v1.8.3
GO v1.22+
```

# Setup
Apply terraform script to allocate DynamoDB table Transactions.
```
terraform init

terraform apply
```

# Creating the txns.csv file
Run the project ```transactions``` first in order to generate the txns.csv file

```
go run ./cmd/main.go 

```

Testing:
```
go test ./tests/
```

# Populate the DB
Run the project ```producer```

