import json
from datetime import datetime

def convert_transaction(transaction):
    return {
        "code": transaction["titulo"],
        "type": transaction["tipo"].upper(),
        "operation": transaction["cv"],
        "quantity": transaction["qtd"],
        "date": datetime.fromisoformat(transaction["data"].replace("Z", "")).strftime("%Y-%m-%d"),
        "unit_price": transaction["valorUnitario"],
        "currency": transaction["moeda"],
        "fees": {
            "settlement": abs(transaction.get("liquidacao", 0)),
            "emolument": abs(transaction.get("emolumentos", 0)),
            "brokerage": abs(transaction.get("corretagem", 0)),
            "iss": abs(transaction.get("iss", 0)),
        },
    }

with open("transactions.json", "r", encoding="utf-8") as file:
    transactions = json.load(file)

converted_transactions = [convert_transaction(tx) for tx in transactions]

with open("formatted_transactions.json", "w", encoding="utf-8") as outfile:
    json.dump(converted_transactions, outfile, ensure_ascii=False, indent=4)

print("Conversão concluída. Salvo em 'formatted_transactions.json'.")
