# stress-test-goexpert
Sistema em Cobra CLI para teste de Stress em aplicações Web

# Desafio GOLang Stress Test - FullCycle 

Aplicação em Go sendo: 
  - Cobra CLI   

&nbsp;
- **Aplicação em Container com - Docker-compose e Dockerfile**

## Funcionalidades

- **Executar Teste Stress**
  - A aplicação permite a execução de chamadas para um endereço HTTP informado, parametrizado com quantidade de requests e quantidade de request simultaneas (Concorrência).

## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.

&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/tiago-g-sales/stress-test-goexpert.git
```

&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd stress-test-goexpert
```

&nbsp;
4. **Rode o docker-compose para buildar:**
&nbsp;

```bash 
 docker-compose up
```
&nbsp;

## Como testar localmente:

### Executar atraves do docker  

```bash 
docker run stress-test-goexpert_cli --url=http://google.com --requests=20 --concurrency=10
```

### Resultado Report contendo as informações sobre a execução conforme exemplo abaixo

```bash 
Tempo Total Execução Segundos-: 1 
Total Requests----------------: 20 
Total Requests HTTP 200-------: 20 
Total Requests HTTP outros----: 0 
```