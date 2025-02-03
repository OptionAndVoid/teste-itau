# teste-itau 🏦

Solução em Go do teste técnico de backend do Itaú

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

Uma API REST para rastrear transações e calcular estatísticas em tempo real
(sum, avg, min, max) dos últimos 60 segundos.\
Desenvolvida em Go em vez de Java 8!

---

## 📂 Estrutura do Projeto **

.\
├── Makefile\
├── cmd\
│ └── teste-itau\
│ └── main.go\
├── internal\
│ ├── api # Os handlers e rotas da API\
│ └── registry # Lógica da nossa estrutura de dados usada como db\
├── pkg\
│ ├── logging # Setup de log\
│ └── server # Cofiguração do Servidor\
└── ssl_credentials # Certificados SSL auto-assinados.

---

## 🚀 Features

- **Endpoints Obrigatórios**:
  - `POST /transacao`: Valida e guarda transações.
  - `DELETE /transacao`: Limpa todas as transações.
  - `GET /estatistica`: Pega estatísticas das transações
  que ocorreram nos últimos 60 segundos.

- **Extras**:  
  - `/metrics`: Métricas para o Prometheus.
  - `/healthcheck`: Monitoramento de saúde da API.
  - **Suporte para TLS**: Conexão encriptada
  - **Swagger UI** (WIP).

---

## 🛠️ Setup

### Pré-Requisitos

- Go 1.20+
- Make (opcional (mas recomendado xD))

### Como Rodar?

1. **Sem TLS**:

   ```bash
   make run
   ```

2. **Com TLS**

   ```bash
   make runtls
   ```

3. **Build**

```bash
make build  # Outputs to ./bin
```

## Extras

### 📈 Métricas Para Prometheus

Configuramos um exporter de métricas para o Prometheus.
As métricas exportadas são sobre o runtime de Go e o quanto de recursos que a API
está utilizando.
As métricas estão disponíveis em: ```http(s)://localhost:8080/metrics```

### 🔒 TLS

Fizemos também configuração de TLS para rodar com https!
Geramos um par .key e .crt autoassinados na pasta ./ssl_credentials
para serem usadas pela API.

### ✅ Rota de Healthcheck

Basta pingar na rota ```/healthcheck``` para saber se a API está saudável e de pé!


