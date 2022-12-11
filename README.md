# Stock Trading Bot

A stock trading bot utilizing TD Ameritrade's API to scan for positions, create and monitor orders, and report results.

## Go Version
    Go v1.19

## Usage
Start-Up using Make:
```
make run -f .\MakeFile
```

## Environment Variables

| Environment Variable Name | Description                    | Required |
|---------------------------|--------------------------------|----------|
| ACCOUNT_ID                | TD Ameritrade Account Username | True     |
| ACCOUNT_KEY               | TD Ameritrade Account Password | True     |
