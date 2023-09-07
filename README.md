# Code Duplicates Detection on Go Repositories 

Escuela Colombiana de Ingeniería Master Thesis on Computer Science - Bogotá, Colombia

## Available Go repositories on GitHub

```python
import requests

api_url = "https://api.github.com/search/repositories"
query_params = {
    "q": "language:go"
}

response = requests.get(api_url, params=query_params)
data = response.json()

total_count = data.get("total_count", 0)
print(f"Total Go repositories on GitHub: {total_count}")
```

Total Go repositories on GitHub: `1,332,711`

## Author

* Julián Benítez Gutiérrez - [julian.benitez@mail.escuelaing.edu.co](mailto:julian.benitez@mail.escuelaing.edu.co)

