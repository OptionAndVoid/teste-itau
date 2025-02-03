# teste-itau 🏦

Solução em Go do teste técnico de backend do Itaú

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

Uma API REST para rastrear transações e calcular estatísticas em tempo real
(sum, avg, min, max) dos últimos 60 segundos.\
Desenvolvida em Go em vez de Java 8!

---

## 📂 Estrutura do Projeto

.\
├── Makefile\
├── README.md\
├── cmd\
│   └── teste-itau # A main está aqui\
├── docker\
│   ├── api # Dockerfile da aplicação principal\
│   ├── grafana # Configurações do Grafana\
│   └── prometheus # Configurações do Prometheus\
├── docker-compose.yaml\
├── enunciado.md # Enúnciado original do teste técnico\
├── go.mod\
├── go.sum\
├── internal\
│   ├── api # Os handlers e rotas da API\
│   └── registry # Lógica da nossa estrutura de dados usada como db\
├── pkg\
│   ├── logging # Setup de Log estruturado\
│   └── server # Configuração do Servidor\
└── ssl_credentials # Certificados SSL auto-assinados\
    ├── server.crt\
    └── server.key

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
  - **Suporte para TLS**: Conexão encriptada.
  - **Swagger UI** (WIP).
  - **Observabilidade**: Grafana e Prometheus no compose.

---

## 🛠️ Setup

### Pré-Requisitos

1. Para Rodar Local

- Go 1.20+
- Make (opcional (mas recomendado xD))

2. Para Rodar no Compose

- docker compose

### Como Rodar?

#### Somente a API

**Sem TLS**:

   ```bash
   make run
   ```

**Com TLS**:

   ```bash
   make runtls
   ```

#### No Docker Compose (Com Grafana e Prometheus)

```bash
docker compose up
```

## Extras

### 🐳 Tudo em Containers

Além de nossa aplicação principal estar 'containerizada',
temos um compose com serviços extras de observabilidade
para constante monitoramento de nossa API.

### 📈 Métricas Para Prometheus

Configuramos um exporter de métricas para o Prometheus.
As métricas exportadas são sobre o runtime de Go e o quanto de recursos que a API
está utilizando.
As métricas estão disponíveis em: ```http(s)://localhost:8080/metrics```

### 📊 Dashboard No Grafana

Configuramos um container de Grafana para monitoramento visual das métricas
de uso de recursos da API em Go!
Basta acessar localhost:3000 e entrar com usuário e senha "admin"
(sim, uma senha super forte xD), ir em "dashboards" e você poderá monitorar o uso
de recursos da API em tempo real!

### 🚨 Alerta No Prometheus

Além de configurar métricas, nós fizemos a configuração de um alerta! Caso a API
esteja fique incomunicável por um minuto, um alerta é gerado pelo Prometheus,
podendo posteriormente ser automatizado para ser colocado no slack,
ser enviado por email, ou enviado por quaisquer canais que o AllertManager suporte!
Para ver os alertas, basta ir em ```localhost:9090``` e ir até "alerts".

### 🔒 TLS

Fizemos também configuração de TLS para rodar com https!
Geramos um par .key e .crt auto-assinados na pasta ./ssl_credentials
para serem usadas pela API.

### ✅ Rota de Healthcheck

Basta pingar na rota ```/healthcheck``` para saber se a API está saudável e de pé!

### 📝 Log Estruturado

Utilizamos o pacote slog para produzir Log Estruturado em Json, pronto para ser consumido
por serviços de consumo de log como ElasticSearch, LogStash e Kibana (ELK Stack).
